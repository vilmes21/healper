package main

import (
	"encoding/json"
	// "log"
	"net/http"
	"strconv"

	"../fn"

	// "../db1"
	"fmt"
	"strings"

	"../config"
	"../model"
	"../viewmodel"

	// "dbbase/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func viewDisease(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["id"]) < 1 {
		return
	}

	var tbDise model.Disease
	var dout viewmodel.Dise
	var q1 string = fmt.Sprintf(
		`SELECT COUNT(*) FROM %s WHERE id=$1`,
		tbDise.Table(),
	)
	var count int
	err := db.QueryRow(q1, vars["id"]).Scan(&count)

	if err != nil {
		fn.HandleErr(err, `getDise q1`)
	}

	if count < 1 {
		dout.Msg = config.NotExist
		jsonBytes, _ := json.Marshal(dout)
		w.Write(jsonBytes)
		return
	}

	var tbSource model.Source
	var q2 string = fmt.Sprintf(
		`SELECT %s.name, %s.name_zh, %s.name as sourcename FROM %s 
		JOIN %s
		ON %s.id = %s.%s
		WHERE disease.id=$1`,
		tbDise.Table(),
		tbDise.Table(),
		tbSource.Table(),
		tbDise.Table(),
		tbSource.Table(),
		tbSource.Table(),
		tbDise.Table(),
		tbDise.FieldSourceId(),
	)

	var diseName, diseNamezh string

	err = db.QueryRow(q2, vars["id"]).Scan(&diseName, &diseNamezh, &dout.Source)

	if err != nil {
		fn.HandleErr(err, `getDise q2`)
	}

	dout.Id, _ = strconv.Atoi(vars["id"])
	dout.Name = fmt.Sprintf(`%s %s`, diseNamezh, diseName)
	if strings.Trim(dout.Name, ` `) == config.EmptyStr {
		dout.Name = fmt.Sprintf(`Untitled-%d`, dout.Id)
	}
	dout.Syms = getDiseOrgSyms(dout.Id)
	dout.Pathos = getDisePathos(dout.Id)

	jsonBytes, _ := json.Marshal(dout)
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
