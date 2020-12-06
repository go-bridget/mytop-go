package terminal

import (
	"github.com/carmo-evan/mytop-go/db"
	"github.com/rivo/tview"
)

type App struct {
	application *tview.Application
	table       *tview.Table
	pages       *tview.Pages
	frame 		*tview.Frame
	Monitor     *db.MySQLMonitor
	Refresh     chan struct{}
}
