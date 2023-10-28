package cmd

import (
	"eternal/modules/exploits"
	"eternal/modules/touches"
	"github.com/desertbit/grumble"
)

func init() {
	useCommand := &grumble.Command{
		Name:     "use",
		Help:     "use exploits",
		LongHelp: "use exploit",
	}
	useCommand.AddCommand(&grumble.Command{
		Name:     "EternalBlue",
		Help:     "use EternalBlue",
		LongHelp: "use EternalBlue",
		Usage:    "use EternalBlue",
		Run: func(c *grumble.Context) error {
			exploits.EternalBlue()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "SmbTouch",
		Help:     "use SmbTouch",
		LongHelp: "use SmbTouch",
		Usage:    "use SmbTouch",
		Run: func(c *grumble.Context) error {
			touches.SmbTouch()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "Doublepulsar",
		Help:     "use Doublepulsar",
		LongHelp: "use Doublepulsar",
		Usage:    "use Doublepulsar",
		Run: func(c *grumble.Context) error {
			exploits.Doublepulsar()
			return nil
		},
	})
	App.AddCommand(useCommand)
}
