package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "../db1"
	"../model"
	"../config"
	"../viewmodel"
	"fmt"
	// "dbbase/sql"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func createAuthor(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var author model.Author
	err := decoder.Decode(&author)
	if err != nil {
		panic(err)
	}

	var authorName string = author.Name
	res := viewmodel.JsonRes{}

	

	if authorName == `` {
		log.Print(`authorName is empty string! Fatal.`)
		res.Msg = `authorName cannot be empty`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	err = db.QueryRow("INSERT INTO author (name) values ($1) RETURNING id", authorName).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert author err: %v", err)
		res.Msg = `insert author err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createSource(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var source model.Source
	err := decoder.Decode(&source)

	// fmt.Printf(`After decoding source: %v`, source)
	// fmt.Printf("Source is: %+v\n", source)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if source.AuthorId == 0 || source.LanguageId == 0 {
		res.Msg = `Author & Language is needed.`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	err = db.QueryRow("INSERT INTO source (author_id, language_id, name) values ($1, $2, $3) RETURNING id", source.AuthorId, source.LanguageId, source.Name).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert source err: %v", err)
		res.Msg = `insert source err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createDisease(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var disease model.Disease
	err := decoder.Decode(&disease)

	fmt.Printf("disease is: %+v\n", disease)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if disease.SourceId == 0 {
		res.Msg = `Book reference is needed.`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	err = db.QueryRow("INSERT INTO disease (name, name_zh, source_id, pinyin) values ($1, $2, $3,$4) RETURNING id", disease.Name, disease.Namezh, disease.SourceId, disease.Pinyin).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert disease err: %v", err)
		res.Msg = `insert disease err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func sendHomeFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, `./static/home.html`)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var recipe model.Recipe
	err := decoder.Decode(&recipe)

	// fmt.Printf("recipe is: %+v\n", recipe)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if recipe.SourceId == 0 {
		recipe.SourceId=config.NoSourceId
	}

	err = db.QueryRow("INSERT INTO recipe (name, name_zh, source_id, pinyin) values ($1, $2, $3,$4) RETURNING id", recipe.Name, recipe.Namezh, recipe.SourceId, recipe.Pinyin).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert recipe err: %v", err)
		res.Msg = `insert recipe err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createHerb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var herb model.Herb
	err := decoder.Decode(&herb)

	// fmt.Printf("herb is: %+v\n", herb)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if herb.Name == `` && herb.Namezh==`` {
		res.Msg=`Herb: a name in either language needed.`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	err = db.QueryRow("INSERT INTO herb (name, name_zh, pinyin) values ($1, $2, $3) RETURNING id", herb.Name, herb.Namezh, herb.Pinyin).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert herb err: %v", err)
		res.Msg = `insert herb err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createRecipeHerb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rh model.RecipeHerb
	err := decoder.Decode(&rh)

	// fmt.Printf("rh is: %+v\n", rh)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if rh.RecipeId == 0 || rh.HerbId == 0 {
		res.Msg=`Recipe AND Herb must already exist in database.`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	_, err = db.Query("INSERT INTO recipe_herb (recipe_id,herb_id,gram,note) values ($1, $2, $3,$4)", rh.RecipeId, rh.HerbId, rh.Gram, rh.Note)

	if err != nil {
		fmt.Printf("insert rh err: %v", err)
		res.Msg = `insert rh err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createPolicy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var policy model.Policy
	err := decoder.Decode(&policy)

	// fmt.Printf("policy is: %+v\n", policy)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if policy.TreatmentVerbId == 0 {
		res.Msg=`Treatment verb must already exist`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	if policy.OrganId==0{
		policy.OrganId=config.BlankOrganId
	}

	if policy.TcmOrganId==0{
		policy.TcmOrganId=config.BlankOrganId
	}

	err = db.QueryRow("INSERT INTO policy (treatment_verb_id,organ_id, tcm_organ_id) values ($1, $2, $3) RETURNING id", policy.TreatmentVerbId, policy.OrganId, policy.TcmOrganId).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert policy err: %v", err)
		res.Msg = `insert policy err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createDiseasePolicy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dp model.DiseasePolicy
	err := decoder.Decode(&dp)

	// fmt.Printf("dp is: %+v\n", dp)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if dp.DiseaseId == 0 || dp.PolicyId==0 {
		res.Msg=`Disease AND Policy must both already exist`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	_, err = db.Query("INSERT INTO disease_policy (disease_id,policy_id, priority) values ($1, $2, $3)", dp.DiseaseId, dp.PolicyId, dp.Priority)

	if err != nil {
		fmt.Printf("insert dp err: %v", err)
		res.Msg = `insert dp err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createOrganSymptom(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var os model.OrganSymptom
	err := decoder.Decode(&os)

	// fmt.Printf("os is: %+v\n", os)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if os.SymptomId==0 {
		res.Msg=`Symptom id must already exist`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	if os.OrganId==0 {
		os.OrganId=config.BlankOrganId
	}

	err = db.QueryRow("INSERT INTO organ_symptom (organ_id,symptom_id) values ($1, $2) RETURNING id", os.OrganId, os.SymptomId).Scan(&res.Id)

	if err != nil {
		fmt.Printf("insert os err: %v", err)
		res.Msg = `insert os err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func createDiseaseOrgansymptom(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ds model.DiseaseOrgansymptom
	err := decoder.Decode(&ds)

	// fmt.Printf("ds is: %+v\n", ds)

	if err != nil {
		panic(err)
	}

	res := viewmodel.JsonRes{}
	

	if ds.OrgansymptomId==0 || ds.DiseaseId==0 {
		res.Msg=`Disease id & Organ-symptom id must both already exist`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	_,err = db.Query("INSERT INTO disease_organsymptom (disease_id,organ_symptom_id) values ($1, $2)", ds.DiseaseId, ds.OrgansymptomId)

	if err != nil {
		fmt.Printf("insert ds err: %v", err)
		res.Msg = `insert ds err. ` + err.Error()

		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	res.Ok = true
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}