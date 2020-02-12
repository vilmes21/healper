package main

import (
	"../db1"
	"log"
	"net/http"
	// "dbbase/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db = db1.Init()

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", sendHomeFile).Methods("GET")

	diseaseRouter := r.PathPrefix("/disease").Subrouter()
	diseaseRouter.HandleFunc("/create", createDisease).Methods("POST")
	// diseaseRouter.HandleFunc("/edit/{id}", GetBook)

	r.HandleFunc("/author/create", createAuthor).Methods("POST")

	r.HandleFunc("/recipe/create", createRecipe).Methods("POST")

	r.HandleFunc("/herb/create", createHerb).Methods("POST")
	r.HandleFunc("/recipe-herb/create", createRecipeHerb).Methods("POST")

	r.HandleFunc("/source/create", createSource).Methods("POST")
	r.HandleFunc("/policy/create", createPolicy).Methods("POST")
	r.HandleFunc("/disease-policy/create", createDiseasePolicy).Methods("POST")
	r.HandleFunc("/organ-symptom/create", createOrganSymptom).Methods("POST")

	r.HandleFunc("/disease-symptom/create", createDiseaseOrgansymptom).Methods("POST")

	port := ":3000"
	log.Print(`Running at http://localhost` + port)
	log.Fatal(http.ListenAndServe(port, r))
}
