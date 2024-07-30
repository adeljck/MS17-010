package cmd

import (
	"eternal/metadatas"
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
			metadatas.MetaData.ShowMetaData()
			return nil
		},
	})
	App.AddCommand(showCommand)
}
