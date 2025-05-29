package terminal

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *App) getTableInputHandlerFunc() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		}
		switch event.Rune() {
		case 'q':
			app.Stop()
		case 's':
			app.Monitor.ToggleSortColumn()
		case 'k':
			app.showKillById()
		case 'K':
			app.showKillAll()
		case 'f':
			app.showFilterByQuery()
		case 'u':
			app.showFilterByUser()
		case 't':
			app.showFilterByTime()
		}
		app.Refresh <- struct{}{}
		return event
	}
}

func (app *App) showFilterByQuery() {
	pageName := "Filter By Query Input Field"
	inputField := newInputField("Filter By Query: ", tview.InputFieldMaxLength(250))
	inputField.SetText(app.Monitor.QueryFilter)
	inputField.SetDoneFunc(func(key tcell.Key) {
		input := inputField.GetText()
		if key == tcell.KeyESC {
			app.pages.RemovePage(pageName)
			return
		}
		app.Monitor.QueryFilter = input
		app.pages.RemovePage(pageName)
		app.Refresh <- struct{}{}
	})
	app.pages.AddPage(pageName, newModal(inputField, 31, 3), true, true)
	app.application.SetFocus(inputField)
}

func (app *App) showFilterByTime() {
	pageName := "Filter By Time Input Field"
	inputField := newInputField("Filter By Time: ", tview.InputFieldMaxLength(250))
	inputField.SetText(app.Monitor.TimeFilter)
	inputField.SetAcceptanceFunc(tview.InputFieldInteger)
	inputField.SetDoneFunc(func(key tcell.Key) {
		input := inputField.GetText()
		if key == tcell.KeyESC {
			app.pages.RemovePage(pageName)
			return
		}
		app.Monitor.TimeFilter = input
		app.pages.RemovePage(pageName)
		app.Refresh <- struct{}{}
	})
	app.pages.AddPage(pageName, newModal(inputField, 31, 3), true, true)
	app.application.SetFocus(inputField)
}

func (app *App) showFilterByUser() {
	pageName := "Filter By User Input Field"
	inputField := newInputField("Filter By User: ", tview.InputFieldMaxLength(250))
	inputField.SetText(app.Monitor.UserFilter)
	inputField.SetDoneFunc(func(key tcell.Key) {
		input := inputField.GetText()
		if key == tcell.KeyESC {
			app.pages.RemovePage(pageName)
			return
		}
		app.Monitor.UserFilter = input
		app.pages.RemovePage(pageName)
		app.Refresh <- struct{}{}
	})
	app.pages.AddPage(pageName, newModal(inputField, 31, 3), true, true)
	app.application.SetFocus(inputField)
}

func (app *App) showKillById() {
	// kill by id
	pageName := "Kill Input Field"
	inputField := newInputField("Enter PID to kill: ", tview.InputFieldInteger)
	inputField.SetDoneFunc(app.getKillDoneFunc(inputField))
	app.pages.AddPage(pageName, newModal(inputField, 31, 3), true, true)
	app.application.SetFocus(inputField)
}

func (app *App) showKillAll() {
	// kill by id
	pageName := "Kill All Input Field"
	button := tview.NewButton("Press Enter to Continue, ESC to cancel")
	button.SetTitle("Kill All Running Processes?")
	button.SetLabelColorActivated(tcell.ColorRed)
	button.SetBackgroundColorActivated(tcell.ColorBlack)
	button.SetBorder(true)
	button.SetSelectedFunc(func() {
		// TODO: Implement Kill All
		for _, p := range app.Monitor.ProcessList {
			app.Monitor.Kill(int(p.Id))
		}
		app.Stop()
	})
	button.SetExitFunc(func(key tcell.Key) {
		if key == tcell.KeyESC {
			app.pages.RemovePage(pageName)
		}
	})
	app.pages.AddPage(pageName, newModal(button, 50, 3), true, true)
	app.application.SetFocus(button)
}

func (app *App) getKillDoneFunc(inputField *tview.InputField) func(key tcell.Key) {
	return func(key tcell.Key) {
		pageName := "Kill Input Field"
		input := inputField.GetText()
		if input == "" || key == tcell.KeyESC {
			app.pages.RemovePage(pageName)
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
		app.pages.RemovePage(pageName)
		app.Refresh <- struct{}{}
	}

}
