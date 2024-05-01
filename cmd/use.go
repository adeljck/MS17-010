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
		Name:     "Eternalblue",
		Help:     "use EternalBlue",
		LongHelp: "use EternalBlue",
		Usage:    "use EternalBlue",
		Run: func(c *grumble.Context) error {
			exploits.EternalBlue()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "Eternalchampion",
		Help:     "use EternalChampion",
		LongHelp: "use EternalChampion",
		Usage:    "use EternalChampion",
		Run: func(c *grumble.Context) error {
			exploits.EternalChampion()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "Smbtouch",
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
	useCommand.AddCommand(&grumble.Command{
		Name:     "Eternalromance",
		Help:     "use EternalRomance",
		LongHelp: "use EternalRomance",
		Usage:    "use EternalRomance",
		Run: func(c *grumble.Context) error {
			exploits.EternalRomance()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "Eternalsynergy",
		Help:     "use Eternalsynergy",
		LongHelp: "use Eternalsynergy",
		Usage:    "use Eternalsynergy",
		Run: func(c *grumble.Context) error {
			exploits.EternalSynergy()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "Explodingcan",
		Help:     "use Explodingcan",
		LongHelp: "use Explodingcan",
		Usage:    "use Explodingcan",
		Run: func(c *grumble.Context) error {
			exploits.ExplodingCan()
			return nil
		},
	})
	useCommand.AddCommand(&grumble.Command{
		Name:     "NamedPipeTouch",
		Help:     "use NamedPipeTouch",
		LongHelp: "use NamedPipeTouch",
		Usage:    "use NamedPipeTouch",
		Run: func(c *grumble.Context) error {
			touches.NamePipeTouch()
			return nil
		},
	})
	App.AddCommand(useCommand)
}
