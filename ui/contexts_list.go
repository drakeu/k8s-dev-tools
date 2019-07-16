package ui

import (
	"github.com/drakeu/k8s-dev-tools/config"

	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ContextsList struct {
	config *config.Config
}

func NewContextsList(config *config.Config) *ContextsList {
	return &ContextsList{
		config: config,
	}
}

func (c *ContextsList) Render() {
	l := widgets.NewList()
	l.Title = "List"
	l.Rows = c.config.GetAvailableContexts()

	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 70, 8)
	termui.Render(l)

	previousKey := ""
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		termui.Render(l)
	}
}
