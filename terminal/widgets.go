package terminal

import "github.com/rivo/tview"

func newInputField(label string, acceptanceFunc func(textToCheck string, lastChar rune) bool) *tview.InputField {
	inputField := tview.NewInputField()
	inputField.SetLabel(label)
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(acceptanceFunc)
	inputField.SetBorder(true)
	return inputField
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