package conn

import (
	"database/sql"
	"fmt"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "172.26.192.107"
	PORT     = 3306
	DATABASE = "go_interface"
)

var conn string = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

func MysqlConn() *sql.DB {
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return nil
	}

	return DB
}

func PrestoConn(dest string) *sql.DB {
	var presto_dsn string = "http://hadoop@10.0.224.131:8988?"
	if dest == "hive" {
		presto_dsn += "catalog=hive&schema=tap_dm"
	} else if dest == "hbase" {
		presto_dsn += "catalog=phoenix&schema=default"
	}
	db, err := sql.Open("presto", presto_dsn)
	println("DEBUG: ", presto_dsn, err)
	return db
}
