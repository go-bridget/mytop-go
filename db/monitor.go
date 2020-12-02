package db

type Monitor interface {
	ShowProcessList() []ProcessList
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
}
