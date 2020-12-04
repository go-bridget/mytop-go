package db

import (
	"database/sql"
	"flag"
)

type Process struct {
	Id           int            `db:"ID"`
	Host         string         `db:"HOST"`
	User         string         `db:"USER"`
	Db           sql.NullString `db:"DB"`
	Command      string         `db:"COMMAND"`
	Time         int            `db:"TIME"`
	State        string         `db:"STATE"`
	Info         sql.NullString `db:"INFO"`
	RowsSent     int            `db:"ROWS_SENT"`
	RowsExamined int            `db:"ROWS_EXAMINED"`
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
	flag.StringVar(&o.Password, "p", "example", "Password")
	flag.StringVar(&o.Database, "d", "mysql", "Database")
	flag.StringVar(&o.Port, "P", "3306", "Port")

	flag.IntVar(&o.Delay, "s", 1, "Delay")
	flag.BoolVar(&o.SkipIdle, "i", false, "Hide Idle (sleeping) threads")
	return o
}
