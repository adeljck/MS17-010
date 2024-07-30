package modules

import (
	"eternal/cmd"
	"eternal/conf"
	"eternal/metadatas"
	"eternal/utils"
	"fmt"
	"github.com/desertbit/grumble"
	"os"
	"runtime"
)

func Run() {
	osName := runtime.GOOS
	if osName != "windows" {
		fmt.Println("This program only runs on Windows.")
		os.Exit(1)
	}
	err := utils.UpdatePathEnv(conf.LibX86)
	if err != nil {
		fmt.Println(err)
	}
	metadatas.SetDefaultMetaData()
	grumble.Main(cmd.App)
}
