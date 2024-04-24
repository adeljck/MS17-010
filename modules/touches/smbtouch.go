package touches

import (
	"eternal/metatdatas"
	"eternal/utils"
	"os/exec"
)

func SmbTouch() {
	//--TargetIp {} --TargetPort {} --Protocol {} --NetworkTimeout {}
	cmd := exec.Command(".\\exp_lib\\Smbtouch-1.1.1.exe", "--InConfig", ".\\exp_lib\\Smbtouch-1.1.1.xml", "--OutConfig", "log.txt", "--TargetIp", metatdatas.MetaData.TargetIp, "--TargetPort", metatdatas.MetaData.TargetPort, "--Protocol", metatdatas.MetaData.Protocol, "--NetworkTimeout", metatdatas.MetaData.NetworkTimeout)
	utils.Execute(cmd)
}
