package view

import (
	"tong_simulator/internal/controller"
	mainscreen "tong_simulator/internal/view/main_screen_components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainScreenMakeUI(ctrl *controller.Controller) fyne.CanvasObject {

	sensors := container.NewGridWithColumns(
		2,
		container.NewCenter(mainscreen.NewRPMLabel(ctrl)),
		container.NewCenter(mainscreen.NewTorqueLabel(ctrl)),
	)

	lables := container.NewGridWithColumns(
		2,
		container.NewCenter(widget.NewLabel("RPM")),
		container.NewCenter(widget.NewLabel("TORQUE")),
	)

	buttons := container.NewGridWithColumns(
		4,
		mainscreen.NewRotateButton(ctrl),
		mainscreen.NewTorqueButton(ctrl),
		mainscreen.NewDumpButton(ctrl),
		mainscreen.NewResetButton(ctrl),
	)

	dumpLevelWitget := mainscreen.NewDumpLevelWidget(ctrl)

	wrapper := container.NewVBox(sensors, lables, buttons, dumpLevelWitget)

	return wrapper
}
