package service

import (
	"database/sql"
	"fmt"
	. "function"
	"presto"
	"strings"

	"common/conn"

	_ "github.com/go-sql-driver/mysql"
)

const MAX_INSERT_COUNT = 50000

func IsExistHiveTable(table string) bool {
	return presto.QueryTable(table, "hive")
}

func BindAkAcessTables(ak string, tables []string) {
	_tables := strings.Join(RemoveStringSliceRepeatedElement(tables), ",")
	DB := conn.MysqlConn()
	defer DB.Close()
	row, err := DB.Query("update access set tables = ? where ak = ?", _tables, ak)
	if err != nil {
		fmt.Println(err)
	}
	row.Close()
	DB.Close()
}

func VerifyAk(ak string) (bool, []string) {
	DB := conn.MysqlConn()
	row := DB.QueryRow("select ak, tables from access where ak = ?", ak)
	defer DB.Close()
	var _ak sql.NullString
	var _tables sql.NullString
	if err := row.Scan(&_ak, &_tables); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		DB.Close()
		return false, []string{}
	}
	DB.Close()
	return true, strings.Split(_tables.String, ",")
}

func BuildInsertSql(resp map[string]interface{}) string {
	dest, table, option, total := resp["dest"], resp["table"], resp["option"], resp["total"]
	fmt.Println("json info: ", dest, table, option, total)
	fields_arr := InterfaceArrayToStringArray(resp["fields"].([]interface{}))
	var fields_str string = strings.Join(fields_arr, ",")
	data := resp["data"].([]interface{})
	var sql string = ""
	var sql_head string = fmt.Sprintf("insert into %s(%s) values", table, fields_str)
	for _, ele := range data {
		_ele := InterfaceArrayToStringArray(ele.([]interface{}))
		tmp := fmt.Sprintf("(%s),", "'"+strings.Join(_ele, "','")+"'")
		sql += tmp
	}
	sql = sql_head + sql[:len(sql)-1]
	return sql
}

func BuildCreateSql(resp map[string]interface{}) string {
	dest, table, option, ak := resp["dest"], resp["table"], resp["option"], resp["ak"]
	primary := resp["primary"].(string)
	fmt.Println(dest, table, option, ak)
	fields_arr := InterfaceArrayToStringArray(resp["fields"].([]interface{}))
	for i, field := range fields_arr {
		if field == primary {
			fields_arr[i], fields_arr[len(fields_arr)-1] = fields_arr[len(fields_arr)-1], primary
		}
	}
	fields := strings.Join(fields_arr, " varchar, ")
	sql := fmt.Sprintf("create table %s(%s) WITH (format = 'ORC', partitioned_by = ARRAY['%s'])", table, fields+" varchar", primary)
	return sql
}

func VerifyTotal(resp map[string]interface{}) bool {
	data := resp["data"].([]interface{})
	var LEN int = len(data)
	if LEN > MAX_INSERT_COUNT {
		return false
	}
	return true
}
