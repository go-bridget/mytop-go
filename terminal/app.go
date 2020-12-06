package terminal

import (
	"fmt"
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

	// header
	for j, label := range labels {
		c := tview.NewTableCell(label).SetTextColor(tcell.ColorYellow).SetExpansion(1)
		if j == app.Monitor.SortColumn() - 1 {
			c.SetTextColor(tcell.ColorLime)
		}
		app.table.SetCell(0, j, c)
	}

	// rows
	for i, p := range pl {
		for j, label := range labels {
			v := p.GetValueByLabel(label)
			if len(v) > 70 {
				v = fmt.Sprintf("%.70q...", v)
			}
			c := tview.NewTableCell(v).SetTextColor(tcell.ColorWhite).SetExpansion(1)
			app.table.SetCell(i+1, j, c)
		}
	}
}
