package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var version = "0.0.1"
var appName = "MS17-010_NSA"

var App = grumble.New(&grumble.Config{
	Name:                  appName,
	Description:           "NSA Original Tools CLI",
	HistoryFile:           ".ms17",
	Prompt:                "MS » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		banner := ` ______________ 
|  __________  |
| | MS17-010 | |
| |__________| |
|______________|
		version:` + version
		a.Println(banner)
	})
}
