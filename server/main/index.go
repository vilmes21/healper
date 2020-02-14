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

	r.HandleFunc("/", sendHomeFile).Methods(http.MethodGet)

	diseaseRouter := r.PathPrefix("/disease").Subrouter()
	diseaseRouter.HandleFunc("/create", midw(createDisease)).Methods(http.MethodPost)
	// diseaseRouter.HandleFunc("/edit/{id}", GetBook)

	r.HandleFunc("/author/create", midw(createAuthor)).Methods(http.MethodPost)

	r.HandleFunc("/category/{cattypeid}", midw(getCategoryByType)).Methods(http.MethodGet)

	

	r.HandleFunc("/source/index", midw(getSources)).Methods(http.MethodGet)

	r.HandleFunc("/recipe/create", midw(createRecipe)).Methods(http.MethodPost)

	r.HandleFunc("/herb/create", midw(createHerb)).Methods(http.MethodPost)
	r.HandleFunc("/recipe-herb/create", midw(createRecipeHerb)).Methods(http.MethodPost)

	r.HandleFunc("/source/create", midw(createSource)).Methods(http.MethodPost)
	r.HandleFunc("/policy/create", midw(createPolicy)).Methods(http.MethodPost)
	r.HandleFunc("/disease-policy/create", midw(createDiseasePolicy)).Methods(http.MethodPost)
	r.HandleFunc("/organ-symptom/create", midw(createOrganSymptom)).Methods(http.MethodPost)

	r.HandleFunc("/disease-symptom/create", midw(createDiseaseOrgansymptom)).Methods(http.MethodPost)

	  r.Use(mux.CORSMethodMiddleware(r))

	port := ":3000"
	log.Print(`Running at http://localhost` + port)
	log.Fatal(http.ListenAndServe(port, r))
}
