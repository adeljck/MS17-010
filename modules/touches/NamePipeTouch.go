package touches

import (
	"eternal/conf"
	"eternal/metadatas"
	"eternal/utils"
	"os/exec"
	"path/filepath"
)

func NamePipeTouch() {
	cmd := exec.Command(filepath.Join(conf.Touches, "Namedpipetouch-2.0.0.exe"), "--InConfig", filepath.Join(conf.Touches, "Namedpipetouch-2.0.0.0.xml"), "--OutConfig", filepath.Join(metadatas.MetaData.LogDir, "logs.txt"), "--TargetIp", metadatas.MetaData.TargetIp, "--TargetPort", metadatas.MetaData.TargetPort, "--Protocol", metadatas.MetaData.Protocol, "--NetworkTimeout", metadatas.MetaData.NetworkTimeout)
	metadatas.MetaData.ShowMetaData()
	utils.Execute(cmd)
}
