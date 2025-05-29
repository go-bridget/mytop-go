package terminal

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/go-bridget/mytop-go/db"
)

func NewApp(monitor *db.MySQLMonitor) *App {
	return &App{
		tview.NewApplication(),
		tview.NewTable(),
		tview.NewPages(),
		nil,
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
	app.table.SetBorderPadding(0, 1, 0, 0)
	app.frame = tview.NewFrame(app.pages)
	h := fmt.Sprintf("s: change sort column    " +
		"f: filter by query    " +
		"u: filter by user    " +
		"t: filter by time    " +
		"k: kill process by PID    " +
		"K: kill all displayed processes")
	app.frame.AddText(h, false, tview.AlignLeft, tcell.ColorWhite)
}

func (app *App) Run() error {
	if err := app.application.SetRoot(app.frame, true).SetFocus(app.frame).Run(); err != nil {
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
		if j == app.Monitor.SortColumn()-1 {
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
