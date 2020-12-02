package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

type MySQLMonitor struct {
	db *sqlx.DB
}

func (m *MySQLMonitor) ShowProcessList() []ProcessList {
	dest := []ProcessList{}
	err := m.db.Select(&dest, "SHOW FULL PROCESSLIST;")
	if err != nil {
		panic(err)
	}
	return dest
}

func (m *MySQLMonitor) ShowGlobalStatus() {
	_, err := m.db.Query("SHOW GLOBAL STATUS;")
	if err != nil {
		panic(err)
	}
}

func GetMySQLMonitor(user, password, hostname string) Monitor {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%v:%v@(%v:3306)/mysql", user, password, hostname))
	if err != nil {
		log.Fatalln(err)
	}
	s := MySQLMonitor{db: db}
	return &s
}