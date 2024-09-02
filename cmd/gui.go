package main

import (
	"tong_simulator/internal/controller"
	"tong_simulator/internal/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func runGui(ctrl *controller.Controller) {
	a := app.New()
	w := a.NewWindow("TONG_EMULATOR")
	w.Resize(fyne.NewSize(500, 100))
	w.CenterOnScreen()
	w.SetContent(view.MainScreenMakeUI(ctrl))
	w.ShowAndRun()
}
