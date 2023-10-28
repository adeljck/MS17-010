package cmd

import (
	"eternal/metatdatas"
	"github.com/desertbit/grumble"
)

func init() {
	showCommand := &grumble.Command{
		Name:     "show",
		Help:     "show details",
		LongHelp: "Show details of param",
	}
	showCommand.AddCommand(&grumble.Command{
		Name:     "options",
		Help:     "show options",
		LongHelp: "show options",
		Usage:    "show options",
		Run: func(c *grumble.Context) error {
			metatdatas.MetaData.ShowMetaData()
			return nil
		},
	})
	App.AddCommand(showCommand)
}
