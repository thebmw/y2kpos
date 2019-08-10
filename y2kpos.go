package main
//import "github.com/gdamore/tcell"
import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	form := tview.NewForm().
		AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 20, '*', nil).
		AddButton("Login", func() {
			app.SetRoot(modal, false).SetFocus(modal).Run();
		}).
		AddButton("Exit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)




	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
