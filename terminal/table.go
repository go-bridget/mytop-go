package terminal

import (
	"github.com/carmo-evan/mytop-go/db"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApp() *tview.Application {
	return tview.NewApplication()
}

func NewTable() *tview.Table {
	return tview.NewTable()
}

func SetTableData(t *tview.Table, pl db.ProcessList) *tview.Table {
	t.Clear().SetBorders(true)
	labels := pl[0].GetLabels()

	for j, label := range labels {
		t.SetCell(0, j, tview.NewTableCell(label).SetTextColor(tcell.ColorYellow))
	}

	for i, p := range pl {
		for j, label := range labels {
			t.SetCell(i + 1, j, tview.NewTableCell(p.GetValueByLabel(label)).SetTextColor(tcell.ColorWhite))
		}
	}
	return t
}
