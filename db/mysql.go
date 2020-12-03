package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLMonitor struct {
	db *sqlx.DB
}

func (m *MySQLMonitor) ShowProcessList() ([]ProcessList, error) {
	dest := []ProcessList{}
	err := m.db.Select(&dest, "SHOW FULL PROCESSLIST;")
	if err != nil {
		return dest, err
	}
	return dest, nil
}

func (m *MySQLMonitor) ShowGlobalStatus() {
	_, err := m.db.Query("SHOW GLOBAL STATUS;")
	if err != nil {
		panic(err)
	}
}

func GetMySQLMonitor(user, password, hostname string) (Monitor, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%v:%v@(%v:3306)/mysql", user, password, hostname))
	if err != nil {
		return &MySQLMonitor{}, err
	}
	s := MySQLMonitor{db: db}
	return &s, nil
}