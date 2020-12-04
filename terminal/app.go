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

func newKillInputField() *tview.InputField {
	inputField := tview.NewInputField()
	inputField.SetLabel("Enter PID to kill: ")
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(tview.InputFieldInteger)
	inputField.SetBorder(true)
	return inputField
}

func (app *App) Draw() {
	app.application.Draw()
}

func (app *App) Stop() {
	app.application.Stop()
}

func (app *App) Init() {
	app.pages.AddPage("Table", app.table,true, true)
	app.table.SetInputCapture(app.getTableInputHandlerFunc())
}

func (app *App) Run() error {
	if err := app.application.SetRoot(app.pages, true).SetFocus(app.pages).Run(); err != nil {
		return err
	}
	return nil
}

func newModal (p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(p, height, 1, false).
		AddItem(nil, 0, 1, false), width, 1, false).
		AddItem(nil, 0, 1, false)
}

func (app *App) SetTableData(pl db.ProcessList) {
	app.table.Clear().SetBorders(true)
	labels := pl[0].GetLabels()

	for j, label := range labels {
		app.table.SetCell(0, j, tview.NewTableCell(label).SetTextColor(tcell.ColorYellow))
	}

	for i, p := range pl {
		for j, label := range labels {
			app.table.SetCell(i + 1, j, tview.NewTableCell(p.GetValueByLabel(label)).SetTextColor(tcell.ColorWhite))
		}
	}
}
