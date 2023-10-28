package modules

import (
	"eternal/cmd"
	"eternal/metatdatas"
	"github.com/desertbit/grumble"
)

func Run() {
	metatdatas.SetDefaultMetaData()
	grumble.Main(cmd.App)
}
