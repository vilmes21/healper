package main

import (
	"encoding/json"
	// "log"
	"net/http"
	"strconv"
	// "../db1"
	// "../model"
	// "../config"
	"../viewmodel"
	"fmt"
	// "dbbase/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getCategoryByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["cattypeid"]) < 1 {
		return
	}
	
	var vm []viewmodel.Cat

	rows, err := db.Query("SELECT id, name, name_zh AS namezh, pinyin FROM category WHERE category_type_id=$1", vars["cattypeid"])

	defer rows.Close()

	if err != nil {
		fmt.Printf("getCategoryByType err: %v", err)
		jsonBytes, _ := json.Marshal(vm)
		w.Write(jsonBytes)
		return
	}

	cattype, err := strconv.Atoi(vars["cattypeid"])
	if err !=nil {
		fmt.Printf("getCategoryByType converting int err: %v", err)
	}

	for rows.Next() {
		var x viewmodel.Cat 
		x.CatTypeId = cattype
	
		err = rows.Scan(&x.Id, &x.Name, &x.Namezh, &x.Pinyin)
	
		if err != nil {
			fmt.Printf("getCategoryByType loop err: %v", err)
		}
		
		vm = append(vm, x)
	}

	jsonBytes, _ := json.Marshal(vm)
	w.Write(jsonBytes)
}