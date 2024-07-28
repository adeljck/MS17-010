package cmd

import (
	"eternal/metatdatas"
	"eternal/modules/touches"
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
