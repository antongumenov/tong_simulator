package mainscreen

import (
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type DumpButton struct {
	ctrl *controller.Controller
	view *widget.Button
}

func NewDumpButton(ctrl *controller.Controller) fyne.CanvasObject {
	c := &DumpButton{
		ctrl: ctrl,
	}
	return c.makeUI()
}

func (btn *DumpButton) makeUI() fyne.CanvasObject {
	btn.view = widget.NewButton("DUMP", func() {
		btn.ctrl.Dump()
	})
	return btn.view
}
