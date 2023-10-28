package touches

import (
	"eternal/metatdatas"
	"fmt"
	"log"
	"os/exec"
)

func SmbTouch() {
	//--TargetIp {} --TargetPort {} --Protocol {} --NetworkTimeout {}
	cmd := exec.Command(".\\exp_lib\\Smbtouch-1.1.1.exe", "--InConfig", ".\\exp_lib\\Smbtouch-1.1.1.xml", "--OutConfig", "log.txt", "--TargetIp", metatdatas.MetaData.TargetIp, "--TargetPort", metatdatas.MetaData.TargetPort, "--Protocol", metatdatas.MetaData.Protocol, "--NetworkTimeout", metatdatas.MetaData.NetworkTimeout)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
