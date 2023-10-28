package cmd

import (
	"github.com/desertbit/grumble"
	"time"
)

func init() {
	timeCommand := &grumble.Command{
		Name:     "time",
		Help:     "Print Time",
		LongHelp: "Print Time",
	}
	timeCommand.AddCommand(&grumble.Command{
		Name:     "unix",
		Help:     "Print Unix Time",
		LongHelp: "Print Unix Time",
		Run: func(c *grumble.Context) error {
			c.App.Println(time.Now().Unix())
			return nil
		},
	})
	timeCommand.AddCommand(&grumble.Command{
		Name:     "loc",
		Help:     "Print Loc Time",
		LongHelp: "Print Loc Time",
		Run: func(c *grumble.Context) error {
			c.App.Println(time.Now().Format("2006-01-02 15:04:05"))
			return nil
		},
	})
	App.AddCommand(timeCommand)
}
