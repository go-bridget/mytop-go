package terminal

import (
	"github.com/rivo/tview"

	"github.com/go-bridget/mytop-go/db"
)

type App struct {
	application *tview.Application
	table       *tview.Table
	pages       *tview.Pages
	frame       *tview.Frame
	Monitor     *db.MySQLMonitor
	Refresh     chan struct{}
}
