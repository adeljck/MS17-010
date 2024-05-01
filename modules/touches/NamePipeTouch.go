package touches

import (
	"eternal/conf"
	"eternal/metatdatas"
	"eternal/utils"
	"os/exec"
	"path/filepath"
)

func NamePipeTouch() {
	cmd := exec.Command(filepath.Join(conf.Touches, "Namedpipetouch-2.0.0.exe"), "--InConfig", filepath.Join(conf.Touches, "Namedpipetouch-2.0.0.0.xml"), "--OutConfig", filepath.Join(metatdatas.MetaData.LogDir, "logs.txt"), "--TargetIp", metatdatas.MetaData.TargetIp, "--TargetPort", metatdatas.MetaData.TargetPort, "--Protocol", metatdatas.MetaData.Protocol, "--NetworkTimeout", metatdatas.MetaData.NetworkTimeout)
	metatdatas.MetaData.ShowMetaData()
	utils.Execute(cmd)
}
