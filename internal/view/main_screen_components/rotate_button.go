package mainscreen

import (
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type RotateButton struct {
	ctrl *controller.Controller
	view *widget.Button
}

func NewRotateButton(ctrl *controller.Controller) fyne.CanvasObject {
	c := &RotateButton{
		ctrl: ctrl,
	}
	return c.makeUI()
}

func (btn *RotateButton) makeUI() fyne.CanvasObject {
	btn.view = widget.NewButton("ROTATE", func() {
		btn.ctrl.Rotate()
	})
	return btn.view
}
