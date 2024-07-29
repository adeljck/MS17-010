package cmd

import (
	"eternal/metatdatas"
	"eternal/modules/touches"
	"eternal/utils"
	"fmt"
	"github.com/desertbit/grumble"
)

func init() {
	setCommand := &grumble.Command{
		Name:     "set",
		Help:     "set data",
		LongHelp: "set data of param",
	}
	setCommand.AddCommand(&grumble.Command{
		Name:     "TargetIp",
		Help:     "set TargetIp",
		LongHelp: "set TargetIp",
		Usage:    "set TargetIp",
		Run: func(c *grumble.Context) error {
			metatdatas.MetaData.SetTargetIp()
			touches.SmbTouch()
			if utils.CheckDefaultRDPPort(metatdatas.MetaData.TargetIp, "3389") {
				fmt.Println(utils.ColorPrint(1, "RDP port In 3389 Inject RDPbl.dll To Enable Blank Password"))

			} else {
				fmt.Println(utils.ColorPrint(1, "Inject RDP.dll AND RDPbl.dll To Enable RDP ANd Blank Password"))

			}
			return nil
		},
	})
	setCommand.AddCommand(&grumble.Command{
		Name:     "NetworkTimeout",
		Help:     "set NetworkTimeout",
		LongHelp: "set NetworkTimeout",
		Usage:    "set NetworkTimeout",
		Run: func(c *grumble.Context) error {
			metatdatas.MetaData.SetNetworkTimeout()
			return nil
		},
	})
	setCommand.AddCommand(&grumble.Command{
		Name:     "TargetPort",
		Help:     "set TargetPort",
		LongHelp: "set TargetPort",
		Usage:    "set TargetPort",
		Run: func(c *grumble.Context) error {
			metatdatas.MetaData.SetTargetPort()
			return nil
		},
	})
	//setCommand.AddCommand(&grumble.Command{
	//	Name:     "Target",
	//	Help:     "set Target",
	//	LongHelp: "set Target",
	//	Usage:    "set Target",
	//	Run: func(c *grumble.Context) error {
	//		metatdatas.MetaData.SetTarget("*")
	//		return nil
	//	},
	//})
	//setCommand.AddCommand(&grumble.Command{
	//	Name:     "DLLPath",
	//	Help:     "set DLLPath",
	//	LongHelp: "set DLLPath",
	//	Usage:    "set DLLPath",
	//	Run: func(c *grumble.Context) error {
	//		metatdatas.MetaData.SetDLLPath()
	//		return nil
	//	},
	//})
	//setCommand.AddCommand(&grumble.Command{
	//	Name:     "Arch",
	//	Help:     "set Arch",
	//	LongHelp: "set Arch",
	//	Usage:    "set Arch",
	//	Run: func(c *grumble.Context) error {
	//		metatdatas.MetaData.SetArch()
	//		return nil
	//	},
	//})
	App.AddCommand(setCommand)
}
