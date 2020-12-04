package terminal

import (
	"github.com/carmo-evan/mytop-go/db"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	*tview.Table
	*tview.Pages
	Monitor *db.MySQLMonitor
	Refresh chan struct{}
}