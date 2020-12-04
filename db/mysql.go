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
	sortColumn int
	columnCount int
}

func (m *MySQLMonitor) ToggleSortColumn() {
	m.sortColumn = (m.sortColumn % m.columnCount) + 1
}

func (m *MySQLMonitor) GetProcessColumnCount() (int, error) {
	count := 0
	err := m.db.Get(&count, "SELECT count(*) FROM information_schema.columns WHERE table_name = 'PROCESSLIST';")
	return count, err
}

func (m *MySQLMonitor) ShowProcessList(ctx context.Context) (ProcessList, error) {
	dest := ProcessList{}
	if m.options.SkipIdle {
		query := fmt.Sprintf("SELECT * FROM information_schema.processList WHERE `COMMAND` != 'SLEEP' ORDER BY %v ASC;", m.sortColumn)
		err := m.db.SelectContext(ctx, &dest, query)
		return dest, err
	}
	query := fmt.Sprintf("SELECT * FROM information_schema.processList ORDER BY %v ASC;", m.sortColumn)
	err := m.db.SelectContext(ctx, &dest, query)
	return dest, err
}

func (m *MySQLMonitor) ShowGlobalStatus(ctx context.Context) error {
	_, err := m.db.QueryContext(ctx, "SHOW GLOBAL STATUS;")
	return err
}

func NewMySQLMonitor(o *Options) *MySQLMonitor {
	return &MySQLMonitor{
		options: o,
		sortColumn: 6,
	}
}

func (m *MySQLMonitor) Connect(ctx context.Context) (err error) {
	m.db, err = sqlx.ConnectContext(ctx, m.options.Driver, fmt.Sprintf("%s:%s@(%s)/%s", m.options.Username, m.options.Password, net.JoinHostPort(m.options.Hostname, m.options.Port), m.options.Database))
	if err != nil {
		return err
	}
	m.columnCount, err = m.GetProcessColumnCount()
	return err
}
