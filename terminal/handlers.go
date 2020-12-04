package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"strconv"
)

func (app *App) GetTableInputHandlerFunc() func (event *tcell.EventKey) *tcell.EventKey {
	return func (event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		}
		switch event.Rune() {
		case 's':
			app.Monitor.ToggleSortColumn()
		case 'k':
			app.killById()
		case 'K':
			// kill all with confirmation
		case 'f':
			// filter by query
		case 'u':
			// filter by user
		}
		app.Refresh <- struct{}{}
		return event
	}
}

func (app *App) killById() {
	// kill by id
	pageName := "Kill Input Field"
	inputField := newKillInputField()
	inputField.SetDoneFunc(app.getKillDoneFunc(inputField))
	app.Pages.AddPage(pageName, newModal(inputField,31, 3), true, true)
	app.SetFocus(inputField)
}

func (app *App) getKillDoneFunc(inputField *tview.InputField) func(key tcell.Key) {
	return func(key tcell.Key) {
		pageName := "Kill Input Field"
		input := inputField.GetText()
		if input == "" || key == tcell.KeyESC {
			app.Pages.RemovePage(pageName)
			return
		}
		pid, err := strconv.Atoi(input)
		if err != nil {
			app.Stop()
			log.Fatalf("Bad input: %v", err)
		}
		if err := app.Monitor.Kill(pid); err != nil {
			app.Stop()
			log.Fatalf("Error while killing process: %v", err)
		}
		app.Pages.RemovePage(pageName)
	}

}