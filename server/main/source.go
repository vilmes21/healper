package main

import (
	"encoding/json"
	// "log"
	"net/http"
	// "../db1"
	// "../model"
	// "../config"
	"../viewmodel"
	"fmt"
	// "dbbase/sql"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getSources(w http.ResponseWriter, r *http.Request) {
	var vm []viewmodel.Src

	rows, err := db.Query("SELECT source.name AS name, source.id AS id, author.name AS author FROM source JOIN author ON source.author_id = author.id")

	defer rows.Close()

	if err != nil {
		fmt.Printf("getSources err: %v", err)
		jsonBytes, _ := json.Marshal(vm)
		w.Write(jsonBytes)
		return
	}

	for rows.Next() {
		var x viewmodel.Src
	
		err = rows.Scan(&x.Name, &x.Id, &x.Author)
	
		if err != nil {
			fmt.Printf("getSources loop err: %v", err)
		}
		
		vm = append(vm, x)
	}

	jsonBytes, _ := json.Marshal(vm)
	w.Write(jsonBytes)
}