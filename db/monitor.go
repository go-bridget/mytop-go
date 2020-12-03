package db

type Monitor interface {
	ShowProcessList() ([]ProcessList, error)
	ShowGlobalStatus()
}


type ProcessList struct {
	Id int `db:"Id"`
	Host string `db:"Host"`
	User string `db:"User"`
	Db *string `db:"db"`
	Command string `db:"Command"`
	Time int `db:"Time"`
	State string `db:"State"`
	Info *string `db:"Info"`
	RowsSent int `db:"Rows_sent"`
	RowsExamined int `db:"Rows_examined"`
}
