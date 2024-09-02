package mainscreen

import (
	"image/color"
	"strconv"
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type RPMLabel struct {
	name string
	ctrl *controller.Controller
	view *canvas.Text
}

func NewRPMLabel(ctrl *controller.Controller) fyne.CanvasObject {
	l := &RPMLabel{
		name: "rpm_label",
		ctrl: ctrl,
	}
	return l.makeUi()
}

func (l *RPMLabel) makeUi() fyne.CanvasObject {
	l.ctrl.Subscrube(l)

	color := color.NRGBA{R: 68, G: 157, B: 72, A: 255}
	label := canvas.NewText(strconv.FormatInt(l.ctrl.GetEncoder(), 10), color)
	label.TextSize = 50

	l.view = label
	return label
}

func (l *RPMLabel) Update() {
	l.view.Text = (strconv.FormatInt(l.ctrl.GetEncoder(), 10))
	l.view.Refresh()

}

func (l *RPMLabel) GetName() string {
	return l.name
}
