package terminal

import (
	"github.com/carmo-evan/mytop-go/db"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApp(monitor *db.MySQLMonitor) *App {
	return &App{
		tview.NewApplication(),
		tview.NewTable(),
		tview.NewPages(),
		monitor,
		make(chan struct{}),
	}
}

func (app *App) Draw() {
	app.application.Draw()
}

func (app *App) Stop() {
	app.application.Stop()
}

func (app *App) Init() {
	app.pages.AddPage("Table", app.table, true, true)
	app.table.SetInputCapture(app.getTableInputHandlerFunc())
}

func (app *App) Run() error {
	if err := app.application.SetRoot(app.pages, true).SetFocus(app.pages).Run(); err != nil {
		return err
	}
	return nil
}

func (app *App) SetTableData(pl db.ProcessList) {
	app.table.Clear().SetBorders(true)
	labels := db.GetProcessListLabels()

	for j, label := range labels {
		app.table.SetCell(0, j, tview.NewTableCell(label).SetTextColor(tcell.ColorYellow))
	}

	for i, p := range pl {
		for j, label := range labels {
			app.table.SetCell(i+1, j, tview.NewTableCell(p.GetValueByLabel(label)).SetTextColor(tcell.ColorWhite))
		}
	}
}
