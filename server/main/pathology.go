package main

import (
	"encoding/json"
	// "log"
	"net/http"
	// "strconv"
	// "../db1"
	// "../config"
	// "../diseOrgsympField"
	"fmt"
	"strings"

	"../fn"
	"../model"
	"../viewmodel"

	// "dbbase/sql"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getDisePathos(diseId int) []viewmodel.DisePatho {
	var dout []viewmodel.DisePatho

	var q string = `
	select pathology.id as patho_id, 
	organ_tb.id as organ_id, 
	organ_tb.name_zh as organ_namezh, 
	organ_tb.name as organ_name, 
	tcm_organ_tb.id as tcm_organ_id, 
	tcm_organ_tb.name_zh as tcm_organ_namezh, 
	tcm_organ_tb.name as tcm_organ_name,
    adj_tb.id as adj_id,
    adj_tb.name as adj_name,
    adj_tb.name_zh as adj_namezh
	from disease_pathology 
	join pathology on disease_pathology.pathology_id = pathology.id
	join category as organ_tb on pathology.organ_id = organ_tb.id 
	join category as tcm_organ_tb on pathology.tcm_organ_id = tcm_organ_tb.id 
    join category as adj_tb on pathology.description_id = adj_tb.id 
	where disease_pathology.disease_id=$1;
`

	rows, err := db.Query(q, diseId)
	if err != nil {
		fn.HandleErr(err, `q in getDisePathos`)
	}

	//patho_id | organ_id | organ_namezh |   organ_name    | tcm_organ_id | tcm_organ_namezh | tcm_organ_name | adj_id | adj_name  | adj_namezh

	/* to map:

	DiseId     int    `json:"diseId"`
	PathoId    int    `json:"pathoId"`
	OrganId    int    `json:"organId"`
	TcmOrganId int    `json:"tcmOrganId"`
	AdjId      int    `json:"adjId"`
	Txt        string `json:"txt"`

	*/
	for rows.Next() {
		var x viewmodel.DisePatho
		x.DiseId = diseId
		var organNamezh string
		var organName string
		var tcmOrganNamezh string
		var tcmOrganName string
		var adjNamezh string
		var adjName string
		err = rows.Scan(
			&x.PathoId,
			&x.OrganId,
			&organNamezh,
			&organName,
			&x.TcmOrganId,
			&tcmOrganNamezh,
			&tcmOrganName,
			&x.AdjId,
			&adjName,
			&adjNamezh,
		)

		x.Txt = fmt.Sprintf(
			`%s%s%s  %s %s %s`,
			organNamezh,
			tcmOrganNamezh,
			adjNamezh,
			organName,
			tcmOrganName,
			adjName,
		)

		if err != nil {
			fn.HandleErr(err, `rows loop q in getDisePathos`)
		}

		dout = append(dout, x)
	}

	return dout
}

//previous insertDiseSymp
func insertFoobar(diseId int, okSympIds []int) []int {
	var tbDOS model.DiseaseOrgansymptom

	var q6 string = fmt.Sprintf(
		`INSERT INTO %s (%s, %s) VALUES ($1,$2) RETURNING %s`, tbDOS.Table(),
		tbDOS.FieldDiseId(),
		tbDOS.FieldOrgSympId(),
		tbDOS.FieldOrgSympId(),
	)

	var savedOsIds []int
	for _, okId := range okSympIds {
		// fmt.Println(`\n q6 is `, q6)
		var savedOsId int
		err := db.QueryRow(q6, diseId, okId).Scan(&savedOsId)
		if err != nil {
			fn.HandleErr(err, `q6 Qry`)
		}

		if savedOsId > 0 {
			savedOsIds = append(savedOsIds, savedOsId)
		}
	}

	return savedOsIds
}

//prev createOrganSymptom
func foobar(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var din viewmodel.DiseSymptom
	var res viewmodel.SaveOrgSymptomsRes
	err := decoder.Decode(&din)

	fmt.Printf("din is: %+v\n", din)

	if err != nil {
		fn.HandleErr(err, `parsing din`)
		res.Msg = `Prase received data err: ` + err.Error()
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	var osIdsLen int = len(din.OrgsympIds)
	var sIdsLen int = len(din.NewSymps)
	if (din.DiseId == 0) || osIdsLen == 0 && sIdsLen == 0 {
		res.Msg = `Key info missing`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	//create err string builder
	builder := strings.Builder{}
	var tbDOS model.DiseaseOrgansymptom
	var tbOS model.OrganSymptom

	//create new records for tb organ_symptom
	//first, see if pair already exist in db
	var sIdsSaved bool = true
	var newOsIds []int
	if sIdsLen == 0 {
		sIdsSaved = true
	} else {
		var insertable []model.OrganSymptom

		for _, osPair := range din.NewSymps {
			var toSelect = []string{`COUNT(*)`}

			var q3 string = sqlFn.Select(toSelect).From(tbOS.Table()).Where(fmt.Sprintf(`%v = $1 AND %v = $2`, tbOS.FieldOrganId(), tbOS.FieldSymptomId())).End()

			// fmt.Println(`q3 statmt:` + q3)

			var existingCnt int
			err = db.QueryRow(q3, osPair.OrganId, osPair.SymptomId).Scan(&existingCnt)

			if err != nil {
				fn.HandleErr(err, `q3`)
				builder.WriteString(`q3` + err.Error())
			}

			if existingCnt == 0 {
				insertable = append(insertable, osPair)
			}
		}

		if len(insertable) == 0 {
			sIdsSaved = true
		} else {
			var q4 string = fmt.Sprintf(`insert into %s (%s, %s) values ($1,$2) RETURNING id`, tbOS.Table(), tbOS.FieldOrganId(), tbOS.FieldSymptomId())

			fmt.Println(`q4 statmt:` + q4)

			for _, pair := range insertable {
				var newOsId int
				err := db.QueryRow(q4, pair.OrganId, pair.SymptomId).Scan(&newOsId)

				if err != nil {
					fn.HandleErr(err, `q4 stmt`)
					builder.WriteString(`q4 ` + err.Error())
					sIdsSaved = false
				}

				if newOsId > 0 {
					newOsIds = append(newOsIds, newOsId)
				} else {
					sIdsSaved = false
				}
			}
		}
	}

	//BEGIN handle osIds

	var osIdsSaved bool
	var savedOsIds []int

	if osIdsLen == 0 {
		if len(newOsIds) == 0 {
			osIdsSaved = true
		} else {
			savedOsIds = insertDiseSymp(din.DiseId, newOsIds)

			if len(newOsIds) == len(savedOsIds) {
				osIdsSaved = true
			} else {
				osIdsSaved = false
			}
		}
	} else { //client side asked to insert some os ids
		//BEGIN filter non-duplicate orgsympIds to insert
		var toSelect = []string{
			tbDOS.FieldOrgSympId(),
		}
		var q string = sqlFn.Select(toSelect).From(tbDOS.Table()).Where(tbDOS.FieldDiseId() + `=$1`).End()

		// fmt.Println(`q: is:` , q)

		rows, err := db.Query(q, din.DiseId)

		if err != nil {
			fn.HandleErr(err, `q1`)
			builder.WriteString(`q1, ` + err.Error())
		}

		var existingSympIdArr []int
		for rows.Next() {
			var existingSympId int
			err = rows.Scan(&existingSympId)

			if err != nil {
				fn.HandleErr(err, `existingSympIdArr dnc`)
				builder.WriteString(err.Error())
			}

			existingSympIdArr = append(existingSympIdArr, existingSympId)
		}

		var okSympIds []int
		for _, id := range din.OrgsympIds {
			if !fn.Contains(existingSympIdArr, id) {
				okSympIds = append(okSympIds, id)
			}
		}
		//END filter non-duplicate orgsympIds to insert

		if len(newOsIds) > 0 {
			okSympIds = append(okSympIds, newOsIds...)
		}

		if len(okSympIds) == 0 {
			osIdsSaved = true
		} else {
			savedOsIds = insertDiseSymp(din.DiseId, okSympIds)

			if len(okSympIds) == len(savedOsIds) {
				osIdsSaved = true
			} else {
				osIdsSaved = false
			}
		}
	}
	//END handle osIds

	res.Msg = builder.String()
	if osIdsSaved && sIdsSaved {
		res.Ok = true
	}

	res.DiseSyms = getDiseOrgSyms(din.DiseId)
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}
