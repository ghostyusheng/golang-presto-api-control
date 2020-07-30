package main

import (
	//. "conn"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"presto"
	"reflect"
	"service"
	"strings"

	"_http"
	"log"
	"net/http"
)

func insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body_byte, _ := ioutil.ReadAll(r.Body)
	var resp map[string]interface{}
	if err := json.Unmarshal(body_byte, &resp); err != nil {
		panic(err)
	}
	ak, table := resp["ak"].(string), resp["table"].(string)
	is, access_tables := service.VerifyAk(ak)
	if !is {
		_http.Out(w, "fail ak", 9999, [][]string{{}})
		return
	}
	ACCESS := false
	for _, _access_table := range access_tables {
		if _access_table == table {
			ACCESS = true
			break
		}
	}
	if !ACCESS {
		_http.Out(w, "no access", 9999, [][]string{{}})
		return
	}
	fmt.Println("access: ", ACCESS)
	var sql string = service.BuildInsertSql(resp)
	if !service.VerifyTotal(resp) {
		_http.Out(w, "FAIL: len error or overflow", 9999, [][]string{{}})
		return
	}
	presto.Query(sql, "hive")
	_http.Out(w, "success", 200, [][]string{{}})
}

func create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body_byte, _ := ioutil.ReadAll(r.Body)
	var resp map[string]interface{}
	if err := json.Unmarshal(body_byte, &resp); err != nil {
		panic(err)
	}
	ak, table := resp["ak"].(string), resp["table"].(string)
	fmt.Println(ak, table, reflect.TypeOf(table))

	is, access_tables := service.VerifyAk(ak)
	if !is {
		_http.Out(w, "fail ak", 9999, [][]string{{}})
		return
	}

	access_tables = append(access_tables, table)
	if service.IsExistHiveTable(table) {
		fmt.Println("true")
		_http.Out(w, "operate exist table", 9999, [][]string{{}})
		return
	}
	service.BindAkAcessTables(ak, access_tables)
	sql := service.BuildCreateSql(resp)
	presto.Query(sql, "hive")
	fmt.Println(sql)
	_http.Out(w, "success", 200, [][]string{{}})
}

func truncate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body_byte, _ := ioutil.ReadAll(r.Body)
	var resp map[string]interface{}
	if err := json.Unmarshal(body_byte, &resp); err != nil {
		panic(err)
	}
	ak, table, primary, primary_value := resp["ak"].(string), resp["table"].(string), resp["primary"].(string), resp["primary_value"].(string)
	fmt.Println(ak, table, reflect.TypeOf(table))

	is, access_tables := service.VerifyAk(ak)
	if !is {
		_http.Out(w, "fail ak", 9999, [][]string{{}})
		return
	}

	ACCESS := false
	for _, _access_table := range access_tables {
		if _access_table == table {
			ACCESS = true
			break
		}
	}
	if !ACCESS {
		_http.Out(w, "no access", 9999, [][]string{{}})
		return
	}
	sql := fmt.Sprintf("delete from tap_dm.%s", table)
	if primary != "" {
		sql += fmt.Sprintf(" where %s = '%s'", primary, primary_value)
	}
	fmt.Println(sql)
	presto.Query(sql, "hive")
	_http.Out(w, "success", 200, [][]string{{}})
}

func query(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body_byte, _ := ioutil.ReadAll(r.Body)
	var resp map[string]interface{}
	if err := json.Unmarshal(body_byte, &resp); err != nil {
		panic(err)
	}
	ak := resp["ak"].(string)
	is, _ := service.VerifyAk(ak)
	if !is {
		_http.Out(w, "fail ak", 9999, [][]string{{}})
		return
	}
	sql := resp["sql"].(string)

	judge_sql := strings.ToLower(sql)
	var is_dangerous_sql bool = strings.Contains(judge_sql, "delete")
	if is_dangerous_sql {
		_http.Out(w, "sql no access", 9999, [][]string{{}})
		return
	}
	is_dangerous_sql = strings.Contains(judge_sql, "drop")
	if is_dangerous_sql {
		_http.Out(w, "sql no access", 9999, [][]string{{}})
		return
	}
	var M [][]string = presto.QueryAll(sql, "hive")
	_http.Out(w, "success", 200, M)
}

func main() {
	fmt.Println("start listen 8088")
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/create", create)
	http.HandleFunc("/truncate", truncate)
	http.HandleFunc("/query", query)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
