package db

import (
	"context"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
)

type MySQLMonitor struct {
	db      *sqlx.DB
	options *Options
}

func (m *MySQLMonitor) ShowProcessList(ctx context.Context) (ProcessList, error) {
	dest := ProcessList{}
	if m.options.SkipIdle {
		err := m.db.SelectContext(ctx, &dest, "SELECT * FROM information_schema.processList WHERE `COMMAND` != 'SLEEP';")
		return dest, err
	}
	err := m.db.SelectContext(ctx, &dest, "SELECT * FROM information_schema.processList;")
	return dest, err
}

func (m *MySQLMonitor) ShowGlobalStatus(ctx context.Context) error {
	_, err := m.db.QueryContext(ctx, "SHOW GLOBAL STATUS;")
	return err
}

func NewMySQLMonitor(o *Options) *MySQLMonitor {
	return &MySQLMonitor{
		options: o,
	}
}

func (m *MySQLMonitor) Connect(ctx context.Context) (err error) {
	m.db, err = sqlx.ConnectContext(ctx, m.options.Driver, fmt.Sprintf("%s:%s@(%s)/%s", m.options.Username, m.options.Password, net.JoinHostPort(m.options.Hostname, m.options.Port), m.options.Database))
	return
}
