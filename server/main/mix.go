package main

import (
	"encoding/json"
	// "log"
	"net/http"
	// "strconv"
	// "../db1"
	"errors"
	"fmt"
	"strings"

	"../config"
	"../fn"
	"../model"
	"../viewmodel"

	// "dbbase/sql"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//return viewmodel.Selectables
func getSelectables(w http.ResponseWriter, r *http.Request) {
	var dout viewmodel.Selectables

	var tbCat model.Category
	var q string = fmt.Sprintf(
		`select %s,%s,%s,%s,%s from %s where %s in ($1,$2,$3,$4)`,
		tbCat.FieldId(),
		tbCat.FieldCatTypeId(),
		tbCat.FieldName(),
		tbCat.FieldNamezh(),
		tbCat.FieldPinyin(),
		tbCat.Table(),
		tbCat.FieldCatTypeId(),
	)

	rows, err := db.Query(q,
		config.OrganCatType,
		config.PathoAdjCatType,
		config.TreatmentVerbCatType,
		config.SymptomAdjCatType,
	)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getSelectables`)
	}

	var organMap = make(map[int]viewmodel.Cat)
	var pathoAdjMap = make(map[int]viewmodel.Cat)
	var treatVerbMap = make(map[int]viewmodel.Cat)
	var sympAdjMap = make(map[int]viewmodel.Cat)

	for rows.Next() {
		var x viewmodel.Cat
		err = rows.Scan(&x.Id, &x.CatTypeId, &x.Name, &x.Namezh, &x.Pinyin)

		if err != nil {
			fn.HandleErr(err, `mix.go getSelectable rows.Next`)
		}

		x.Txt = fmt.Sprintf(
			`%s %s`,
			x.Namezh,
			x.Name,
		)

		x.Searchable = fmt.Sprintf(
			`%s%s%d`,
			strings.ToLower(x.Txt),
			strings.ToLower(x.Pinyin),
			x.Id,
		)

		switch x.CatTypeId {
		case config.OrganCatType:
			x.FrontendType = config.Organ

			dout.Organs = append(dout.Organs, x)
			organMap[x.Id] = x
		case config.PathoAdjCatType:
			x.FrontendType = config.PathoAdj

			dout.PathoAdjs = append(dout.PathoAdjs, x)
			pathoAdjMap[x.Id] = x
		case config.TreatmentVerbCatType:
			x.FrontendType = config.TreatVerb

			dout.TreatVerbs = append(dout.TreatVerbs, x)
			treatVerbMap[x.Id] = x
		case config.SymptomAdjCatType:
			x.FrontendType = config.SymptomAdj
			dout.SymptomAdjs = append(dout.SymptomAdjs, x)
			sympAdjMap[x.Id] = x
		default:
			err = errors.New(`x.CatTypeId should have matched before this`)
			fn.HandleErr(err, `mix.go getSelectable switch`)
		}
	}

	//step 2, fill dout.Pathos
	var tbPatho model.Pathology
	var q2 string = fmt.Sprintf(
		`select %s,%s,%s,%s from %s`,
		tbPatho.FieldId(),
		tbPatho.FieldOrganId(),
		tbPatho.FieldTcmOrganId(),
		tbPatho.FieldDescriptionId(),
		tbPatho.Table(),
	)
	rows, err = db.Query(q2)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getPatho`)
	}

	for rows.Next() {
		var x viewmodel.Patho
		err = rows.Scan(&x.Id, &x.OrganId, &x.TcmOrganId, &x.DescriptionId)

		fmt.Printf(`after scan,getPatho x:%+v`, x)

		if err != nil {
			fn.HandleErr(err, `mix.go getPatho rows.Next`)
		}

		var organObj = organMap[x.OrganId]
		var organObj2 = organMap[x.TcmOrganId]
		var adjObj = pathoAdjMap[x.DescriptionId]
		x.Txt = fmt.Sprintf(
			`%s%s%s %s-%s-%s`,
			organObj.Namezh,
			organObj2.Namezh,
			adjObj.Namezh,
			organObj.Name,
			organObj2.Name,
			adjObj.Name,
		)
		x.FrontendType = config.Patho
		x.Searchable = fmt.Sprintf(
			`%s%d %d %d`,
			strings.ToLower(x.Txt),
			x.OrganId,
			x.TcmOrganId,
			x.DescriptionId,
		)
		dout.Pathos = append(dout.Pathos, x)
	}

	//step 3 fill dout.OrganSymptoms
	var tbOS model.OrganSymptom
	var q3 string = fmt.Sprintf(
		`select %s,%s,%s from %s`,
		tbOS.FieldId(),
		tbOS.FieldOrganId(),
		tbOS.FieldSymptomId(),
		tbOS.Table(),
	)
	rows, err = db.Query(q3)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getOrgSymp`)
	}

	for rows.Next() {
		var x viewmodel.OrganSymp
		err = rows.Scan(&x.Id, &x.OrganId, &x.SymptomId)

		fmt.Printf(`after scan,getOrgSymp x:%+v`, x)

		if err != nil {
			fn.HandleErr(err, `mix.go getOrgSymp rows.Next`)
		}

		var organObj = organMap[x.OrganId]
		var adjObj = sympAdjMap[x.SymptomId]
		x.Txt = fmt.Sprintf(
			`%s%s %s-%s`,
			organObj.Namezh,
			adjObj.Namezh,
			organObj.Name,
			adjObj.Name,
		)
		x.FrontendType = config.OrganSymptom
		x.Searchable = fmt.Sprintf(
			`%s%s%s%d %d`,
			strings.ToLower(x.Txt),
			strings.ToLower(organObj.Pinyin),
			strings.ToLower(adjObj.Pinyin),
			x.OrganId,
			x.SymptomId,
		)
		dout.OrganSymptoms = append(dout.OrganSymptoms, x)
	}

	//step 4, fill dout.Policies
	var tbPolicy model.Policy
	var q4 string = fmt.Sprintf(
		`select %s,%s,%s,%s from %s`,
		tbPolicy.FieldId(),
		tbPolicy.FieldTreatVerbId(),
		tbPolicy.FieldOrganId(),
		tbPolicy.FieldTcmOrganId(),
		tbPolicy.Table(),
	)
	rows, err = db.Query(q4)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getPolicy`)
	}

	for rows.Next() {
		var x viewmodel.Policy
		err = rows.Scan(&x.Id, &x.TreatVerbId, &x.OrganId, &x.TcmOrganId)

		fmt.Printf(`after scan,getPolicy x:%+v`, x)

		if err != nil {
			fn.HandleErr(err, `mix.go getPolicy rows.Next`)
		}

		var verbObj = treatVerbMap[x.TreatVerbId]
		var organObj = organMap[x.OrganId]
		var organObj2 = organMap[x.TcmOrganId]
		x.Txt = fmt.Sprintf(
			`%s%s%s %s %s-%s`,
			verbObj.Namezh,
			organObj.Namezh,
			organObj2.Namezh,
			verbObj.Name,
			organObj.Name,
			organObj2.Name,
		)
		x.FrontendType = config.Policy
		x.Searchable = fmt.Sprintf(
			`%s%s%s%s%d %d %d`,
			strings.ToLower(x.Txt),
			strings.ToLower(verbObj.Pinyin),
			strings.ToLower(organObj.Pinyin),
			strings.ToLower(organObj2.Pinyin),
			x.OrganId,
			x.TreatVerbId,
			x.TcmOrganId,
		)
		dout.Policies = append(dout.Policies, x)
	}

	//step 5 fill dout.Recipes
	var tbRecipe model.Recipe
	var q5 string = fmt.Sprintf(
		`select %s,%s,%s,%s from %s`,
		tbRecipe.FieldId(),
		tbRecipe.FieldName(),
		tbRecipe.FieldNamezh(),
		tbRecipe.FieldPinyin(),
		tbRecipe.Table(),
	)
	rows, err = db.Query(q5)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getRecipes`)
	}

	for rows.Next() {
		var x viewmodel.Recipe
		err = rows.Scan(&x.Id, &x.Name, &x.Namezh, &x.Pinyin)

		fmt.Printf(`after scan,getRecipes x:%+v`, x)

		if err != nil {
			fn.HandleErr(err, `mix.go getRecipes rows.Next`)
		}

		x.FrontendType = config.Recipe
		x.Txt = fmt.Sprintf(
			`%s %s`,
			x.Namezh,
			x.Name,
		)
		x.Searchable = fmt.Sprintf(
			`%s%s%d`,
			strings.ToLower(x.Txt),
			strings.ToLower(x.Pinyin),
			x.Id,
		)

		dout.Recipes = append(dout.Recipes, x)
	}

	//step 6 fill dout.Herbs
	var tbHerb model.Herb
	var q6 string = fmt.Sprintf(
		`select %s,%s,%s,%s from %s`,
		tbHerb.FieldId(),
		tbHerb.FieldName(),
		tbHerb.FieldNamezh(),
		tbHerb.FieldPinyin(),
		tbHerb.Table(),
	)
	rows, err = db.Query(q6)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `mix.go getHerbs`)
	}

	for rows.Next() {
		var x viewmodel.Herb
		err = rows.Scan(&x.Id, &x.Name, &x.Namezh, &x.Pinyin)

		fmt.Printf(`after scan,getHerbs x:%+v`, x)

		if err != nil {
			fn.HandleErr(err, `mix.go getHerbs rows.Next`)
		}

		x.FrontendType = config.Herb
		x.Txt = fmt.Sprintf(
			`%s %s`,
			x.Namezh,
			x.Name,
		)
		x.Searchable = fmt.Sprintf(
			`%s%s%d`,
			strings.ToLower(x.Txt),
			strings.ToLower(x.Pinyin),
			x.Id,
		)

		dout.Herbs = append(dout.Herbs, x)
	}

	jsonBytes, _ := json.Marshal(dout)
	w.Write(jsonBytes)
}
