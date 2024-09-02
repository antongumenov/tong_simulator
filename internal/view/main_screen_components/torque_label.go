package mainscreen

import (
	"image/color"
	"strconv"
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type TorqueLabel struct {
	name string
	ctrl *controller.Controller
	view *canvas.Text
}

func NewTorqueLabel(ctrl *controller.Controller) fyne.CanvasObject {
	l := &TorqueLabel{
		name: "torque",
		ctrl: ctrl,
	}
	return l.makeUi()
}

func (l *TorqueLabel) makeUi() fyne.CanvasObject {
	l.ctrl.Subscrube(l)

	color := color.NRGBA{R: 204, G: 69, B: 27, A: 255}
	label := canvas.NewText(strconv.FormatInt(l.ctrl.GetLoadCell(), 10), color)
	label.TextSize = 50

	l.view = label
	return label
}

func (l *TorqueLabel) Update() {
	l.view.Text = strconv.FormatInt(l.ctrl.GetLoadCell(), 10)
	l.view.Refresh()
}

func (l *TorqueLabel) GetName() string {
	return l.name
}
