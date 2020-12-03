package db

import (
	"flag"
	"os"

	"database/sql"
)

type Monitor interface {
	ShowProcessList() (ProcessList, error)
	ShowGlobalStatus()
}

type Process struct {
	Id           int            `db:"Id"`
	Host         string         `db:"Host"`
	User         string         `db:"User"`
	Db           sql.NullString `db:"db"`
	Command      string         `db:"Command"`
	Time         int            `db:"Time"`
	State        string         `db:"State"`
	Info         sql.NullString `db:"Info"`
	RowsSent     int            `db:"Rows_sent"`
	RowsExamined int            `db:"Rows_examined"`
}

type ProcessList []Process

type Options struct {
	// mysql, but we could support postgres... sometime
	Driver string

	// connection credentials
	Hostname string
	Username string
	Password string
	Database string
	Port     string

	// polling interval
	Delay int

	// skip idle connections
	SkipIdle bool
}

func NewOptions() *Options {
	return new(Options).Bind()
}

func (o *Options) Bind() *Options {
	flag.StringVar(&o.Driver, "D", "mysql", "SQL Driver name")
	flag.StringVar(&o.Hostname, "h", "127.0.0.1", "Hostname")
	flag.StringVar(&o.Username, "u", "root", "Username")
	flag.StringVar(&o.Password, "p", os.Getenv("MYSQL_ROOT_PASSWORD"), "Password")
	flag.StringVar(&o.Database, "d", "mysql", "Database")
	flag.StringVar(&o.Port, "P", "3306", "Port")

	flag.IntVar(&o.Delay, "s", 5, "Delay")
	flag.BoolVar(&o.SkipIdle, "i", false, "Hide Idle (sleeping) threads")
	return o
}
