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

func getDiseOrgSyms(diseId int) []viewmodel.DiseSym {
	var dout []viewmodel.DiseSym

	var q string = `
	select organ_symptom.id as osId, 
	organTb.id as organId, 
	organTb.name_zh as organNamezh, 
	organTb.name as organName, 
	symTb.id as symId, 
	symTb.name_zh as symNamezh, 
	symTb.name as symName
	from disease_organsymptom 
	join organ_symptom on disease_organsymptom.organ_symptom_id = organ_symptom.id
	join category as organTb on organ_symptom.organ_id = organTb.id 
	join category as symTb on organ_symptom.symptom_id = symTb.id 
	where disease_organsymptom.disease_id=$1`

	rows, err := db.Query(q, diseId)
	if err != nil {
		fn.HandleErr(err, `q in getDiseOrgSyms`)
	}

	//osid | organid | organnamezh |    organname    | symid | symnamezh | symname

	/* to map:
	 DiseId
	OrgsympId
	OrganId
	SymptomId
	Txt
	*/
	for rows.Next() {
		var x viewmodel.DiseSym
		x.DiseId = diseId
		var organNamezh string
		var organName string
		var symNamezh string
		var symName string
		err = rows.Scan(&x.OrgsympId, &x.OrganId, &organNamezh, &organName, &x.SymptomId, &symNamezh, &symName)

		x.Txt = fmt.Sprintf(
			`%s%s  %s %s`,
			organNamezh,
			symNamezh,
			organName,
			symName,
		)

		if err != nil {
			fn.HandleErr(err, `rows loop q in getDiseOrgSyms`)
		}

		dout = append(dout, x)
	}

	return dout
}

func insertDiseSymp(diseId int, okSympIds []int) []int {
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

func createOrganSymptom(w http.ResponseWriter, r *http.Request) {
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
