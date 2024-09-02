package mainscreen

import (
	"regexp"
	"strconv"
	"strings"
	"tong_simulator/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DumpLevelWigget struct {
	ctrl *controller.Controller
	view *fyne.Container
}

func NewDumpLevelWidget(ctrl *controller.Controller) fyne.CanvasObject {
	c := &DumpLevelWigget{
		ctrl: ctrl,
	}
	return c.makeUI()
}

func (e *DumpLevelWigget) makeUI() fyne.CanvasObject {
	// label
	label := widget.NewLabel("MAX TORQUE:")
	label.Alignment = fyne.TextAlignTrailing

	// entry
	entry := widget.NewEntry()
	entry.PlaceHolder = strconv.FormatInt(e.ctrl.GetDumpLevel(), 10)
	entry.OnCursorChanged = func() {
		if entry.Text != "" {
			currentText := entry.Text
			pattern := regexp.MustCompile("[0-9]+")
			entries := pattern.FindAllString(currentText, -1)
			newText := strings.Join(entries, "")
			if newText != currentText {
				entry.SetText(newText)
			}
		}
	}
	entry.OnSubmitted = func(s string) {
		if s != "" {
			val, _ := strconv.ParseInt(s, 10, 64)
			e.ctrl.SetDumpLevel(val)
		}
		entry.FocusLost()
	}

	// btn
	btn := widget.NewButton("SET MAX TORQUE", func() {
		if entry.Text != "" {
			val, _ := strconv.ParseInt(entry.Text, 10, 64)
			e.ctrl.SetDumpLevel(val)
		}
		entry.FocusLost()
	})

	content := container.NewGridWithColumns(
		3,
		label,
		entry,
		btn,
	)

	e.view = content
	return content
}
