package presto

import (
	"common/conn"
	"database/sql"
	"fmt"
	. "fmt"

	_ "github.com/prestodb/presto-go-client/presto"
)

func Query(_sql string, dest string) {
	db := conn.PrestoConn(dest)
	_, err := db.Query(_sql)
	if err != nil {
		Println("err: ", err)
	}
	//Println("sql ===> ", _sql)
}

func QueryTable(table string, dest string) bool {
	db := conn.PrestoConn(dest)
	_sql := fmt.Sprintf("show tables like '%s'", table)
	rows, err := db.Query(_sql)
	if err != nil {
		Println("err: ", err)
	}

	for rows.Next() {
		var tb sql.NullString
		err = rows.Scan(&tb)
		if err != nil {
			Println(err)
			return false
		}
		if tb.Valid {
			return true
		}
		return false
	}
	return false
}

func QueryAll(_sql string, dest string) [][]string {
	db := conn.PrestoConn(dest)
	rows, err := db.Query(_sql)
	if err != nil {
		Println("err: ", err)
	}

	Println("sql ===> ", _sql)
	columns, _ := rows.Columns()
	var LEN = len(columns)
	rawResult := make([][]byte, LEN)
	result := make([]string, LEN)

	_dest := make([]interface{}, LEN)
	var M [][]string

	for i, _ := range rawResult {
		_dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(_dest...)
		if err != nil {
			Println("err: ", err)
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}
		M = append(M, result)
	}
	return M
}
