package touches

import (
	"eternal/conf"
	"eternal/metadatas"
	"eternal/utils"
	"os/exec"
	"path/filepath"
)

func SmbTouch() {
	cmd := exec.Command(filepath.Join(conf.Touches, "Smbtouch-1.1.1.exe"), "--InConfig", filepath.Join(conf.Touches, "Smbtouch-1.1.1.0.xml"), "--OutConfig", filepath.Join(metadatas.MetaData.LogDir, "logs.txt"), "--TargetIp", metadatas.MetaData.TargetIp, "--TargetPort", metadatas.MetaData.TargetPort, "--Protocol", metadatas.MetaData.Protocol, "--NetworkTimeout", metadatas.MetaData.NetworkTimeout)
	metadatas.MetaData.ShowMetaData()
	utils.Execute(cmd)
}
