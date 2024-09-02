package mainscreen

import (
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TorqueButton struct {
	ctrl *controller.Controller
	view *widget.Button
}

func NewTorqueButton(ctrl *controller.Controller) fyne.CanvasObject {
	c := &TorqueButton{
		ctrl: ctrl,
	}
	return c.makeUI()
}

func (l *TorqueButton) makeUI() fyne.CanvasObject {
	l.view = widget.NewButton("TORQUE", func() {
		l.ctrl.TorqueOn()
	})
	return l.view
}
