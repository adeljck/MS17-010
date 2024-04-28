package cmd

import (
	"eternal/utils"
	"fmt"
	"github.com/desertbit/grumble"
	"os/exec"
)

func init() {
	showCommand := &grumble.Command{
		Name:     "shell",
		Help:     "run shell command",
		LongHelp: "run shell command",
		Usage:    "shell your command",
		Run: func(c *grumble.Context) error {
			Command := ""
			fmt.Printf(utils.ColorPrint(1, "Input Your Command:"))
			fmt.Scanf("%s", &Command)
			if Command == "" {
				return nil
			}
			cmd := exec.Command("cmd.exe", "/c", Command)
			utils.Execute(cmd)
			return nil
		},
	}
	App.AddCommand(showCommand)
}
