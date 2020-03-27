package main

import (
	"encoding/json"
	// "log"
	"../fn"
	"net/http"
	"strconv"
	// "../db1"
	"../model"
	"../config"
	"../viewmodel"
	"fmt"
	"strings"
	// "dbbase/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func createCategory(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var din model.Category
	err := decoder.Decode(&din)

	res := viewmodel.JsonRes{}
	if err != nil {
		res.Msg = `Err parsing data`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	din.Namezh = strings.Trim(din.Namezh, config.SpaceStr)
	din.Name = strings.Trim(din.Name, config.SpaceStr)

	if din.CatTypeId == 0 || (len(din.Namezh) == 0 && len(din.Name) == 0) {
		res.Ok = false
		res.Msg = `All data provided are bad`
		jsonBytes, _ := json.Marshal(res)
		w.Write(jsonBytes)
		return
	}

	//check if duplicate in db
	var catTypeIdIn = din.CatTypeId
	var q1 string = fmt.Sprintf(
		`select %s, %s, %s from %s where %s = $1`,
		din.FieldId(),
		din.FieldName(),
		din.FieldNamezh(),
		din.Table(),
		din.FieldCatTypeId(),
	)

	rows, err := db.Query(q1, catTypeIdIn)
	defer rows.Close()

	if err != nil {
		fn.HandleErr(err, `q1`)
	}

	var existing []model.Category
	for rows.Next() {
		var x model.Category

		err = rows.Scan(&x.Id, &x.Name, &x.Namezh)

		if err != nil {
			fn.HandleErr(err, `q1 loop`)
		}

		existing = append(existing, x)
	}

	var existMaxId int = 100*din.CatTypeId + 1
	if len(existing) > 0 {
		for _, x := range existing {
			if x.Namezh == din.Namezh || x.Name == din.Name {
				res.Ok = false
				res.Msg = `Name or name-zh duplicate in database`
				jsonBytes, _ := json.Marshal(res)
				w.Write(jsonBytes)
				return
			}

			if x.Id > existMaxId {
				existMaxId = x.Id
			}
		}
	}

	var q2 string = fmt.Sprintf(
		`insert into %s (%s, %s, %s, %s, %s) values ($1,$2,$3,$4,$5) returning %s`,
		din.Table(),
		din.FieldId(),
		din.FieldCatTypeId(),
		din.FieldName(),
		din.FieldNamezh(),
		din.FieldPinyin(),
		din.FieldId(),
	)

	var newId int

	fmt.Println(fmt.Sprintf(`existMaxId+1: %d`, existMaxId+1))

	err = db.QueryRow(
		q2,
		existMaxId+1,
		din.CatTypeId,
		din.Name,
		din.Namezh,
		din.Pinyin,
	).Scan(&newId)

	if err != nil || newId == 0 {
		fn.HandleErr(err, `q2 queryRow`)
	}

	res.Ok = true
	res.Id = newId
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func getCategoryByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["cattypeid"]) < 1 {
		return
	}

	var vm []model.Category

	rows, err := db.Query("SELECT id, name, name_zh AS namezh, pinyin FROM category WHERE category_type_id=$1", vars["cattypeid"])

	defer rows.Close()

	if err != nil {
		fmt.Printf("getCategoryByType err: %v", err)
		jsonBytes, _ := json.Marshal(vm)
		w.Write(jsonBytes)
		return
	}

	cattype, err := strconv.Atoi(vars["cattypeid"])
	if err != nil {
		fmt.Printf("getCategoryByType converting int err: %v", err)
	}

	for rows.Next() {
		var x model.Category
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
