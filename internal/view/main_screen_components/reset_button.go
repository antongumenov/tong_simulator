package mainscreen

import (
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ResetButton struct {
	ctrl *controller.Controller
	view *widget.Button
}

func NewResetButton(ctrl *controller.Controller) fyne.CanvasObject {
	c := &ResetButton{
		ctrl: ctrl,
	}
	return c.makeUI()
}

func (btn *ResetButton) makeUI() fyne.CanvasObject {
	btn.view = widget.NewButton("RESET", func() {
		btn.ctrl.Reset()
	})
	return btn.view
}
