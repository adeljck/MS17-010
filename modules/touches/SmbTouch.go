package touches

import (
	"eternal/conf"
	"eternal/metatdatas"
	"eternal/utils"
	"os/exec"
	"path/filepath"
)

func SmbTouch() {
	//--TargetIp {} --TargetPort {} --Protocol {} --NetworkTimeout {}
	cmd := exec.Command(conf.Touches+"Smbtouch-1.1.1.exe", "--InConfig", conf.Touches+"Smbtouch-1.1.1.0.xml", "--OutConfig", filepath.Join(metatdatas.MetaData.LogDir, "logs.txt"), "--TargetIp", metatdatas.MetaData.TargetIp, "--TargetPort", metatdatas.MetaData.TargetPort, "--Protocol", metatdatas.MetaData.Protocol, "--NetworkTimeout", metatdatas.MetaData.NetworkTimeout)
	utils.Execute(cmd)
}
