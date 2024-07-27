package metatdatas

import (
	"eternal/utils"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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
	LogDir          string
	ProjectName     string
	IIS             struct {
		Port           string
		EnableSSL      string
		HostString     string
		MaxSizeToCheck string
		Delay          string
	}
}

var MetaData = Meta{
	TargetPort:     "445",
	Protocol:       "SMB",
	NetworkTimeout: "60",
	DLLPath:        "./dlls/adduser.dll",
	Arch:           "x64",
	Target:         "SERVER_2008R2_SP1",
	Function:       "Ping",
	OutputInstall:  "./target.bin",
	ExploitMethod:  "Default",
	LogDir:         "./logs/",
	ProjectName:    "",
	IIS: struct {
		Port           string
		EnableSSL      string
		HostString     string
		MaxSizeToCheck string
		Delay          string
	}{Port: "80",
		EnableSSL:      "false",
		HostString:     "",
		MaxSizeToCheck: "70",
		Delay:          "0"},
}
var eb_targets = []string{"XP", "WIN72K8R2"}
var archs = []string{"x86", "x64"}
var functions = []string{"OutputInstall", "Ping", "RunDLL", "RunShellcode", "Uninstall"}

func SetDefaultMetaData() {
	for {
		fmt.Print(utils.ColorPrint(1, "Set Default Target:"))
		TargetIp := ""
		fmt.Scanf("%s\n", &TargetIp)
		if !utils.IPChecker(TargetIp) {
			fmt.Println(utils.ColorPrint(-1, "Wrong IP Address...."))
		} else {
			MetaData.TargetIp = TargetIp
			MetaData.IIS.HostString = TargetIp
			fmt.Println(utils.ColorPrint(0, "TargetIp ==> %s", MetaData.TargetIp))
			fmt.Println(utils.ColorPrint(0, "HostString ==> %s", MetaData.TargetIp))
			MetaData.ProjectName = fmt.Sprintf("z%s", MetaData.TargetIp)
			break
		}
	}
	LogDir := ""
	for {
		fmt.Print(utils.ColorPrint(1, "Set Base Log Dir(Default ./logs/):"))
		fmt.Scanf("%s\n", &LogDir)
		if utils.CheckPathLegal(LogDir) {
			break
		} else {
			fmt.Println(utils.ColorPrint(-1, "Wrong Log Directory..."))
		}
	}
	if LogDir != "" {
		MetaData.LogDir = LogDir
	}
	fmt.Println(utils.ColorPrint(0, "LogDir ==> %s", MetaData.LogDir))
	utils.CreateLogDir(MetaData.LogDir)
	Projects := utils.ListDirectory(MetaData.LogDir)
	ProjectIndex := 0
	for {

		for index, Project := range Projects {
			fmt.Println(utils.ColorPrint(1, "%d.%s", index, Project))
		}
		fmt.Print(utils.ColorPrint(1, "Choose Project(default 0):"))
		fmt.Scanf("%d\n", &ProjectIndex)
		if ProjectIndex < 0 || ProjectIndex > len(Projects) {
			fmt.Println(utils.ColorPrint(-1, "Wrong Project Index."))
		} else {
			break
		}
	}
	ProjectName := ""
	if ProjectIndex == len(Projects)-1 {
		fmt.Printf(utils.ColorPrint(1, "Set Project Name(default %s):", fmt.Sprintf("z%s", MetaData.TargetIp)))
		fmt.Scanf("%s\n", &ProjectName)
		if ProjectName != "" {
			MetaData.ProjectName = ProjectName
		}
	} else {
		MetaData.ProjectName = Projects[ProjectIndex]
	}
	fmt.Println(utils.ColorPrint(0, "Project ==> %s", fmt.Sprintf("%s", MetaData.ProjectName)))
	utils.CreateProject(MetaData.LogDir, MetaData.ProjectName)
	fmt.Println(utils.ColorPrint(0, "Set Target Directory ==> %s", fmt.Sprintf("z%s", MetaData.TargetIp)))
	utils.CreateProject(filepath.Join(MetaData.LogDir, MetaData.ProjectName), fmt.Sprintf("z%s", MetaData.TargetIp))
	utils.CreateLog(MetaData.TargetIp, MetaData.LogDir, MetaData.ProjectName)
}
func (M *Meta) SetShellcodeFile() {
	ShellcodeFile := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Set DOUP ShellcodeFile:"))
		fmt.Scanf("%s\n", &ShellcodeFile)
		if ShellcodeFile == "" || !utils.CheckPathLegal(ShellcodeFile) {
			fmt.Println(utils.ColorPrint(-1, "Wrong Shellcode File Path "))
			continue
		} else {
			_, err := os.Stat(ShellcodeFile)
			if os.IsNotExist(err) {
				fmt.Println(utils.ColorPrint(-1, "Shellcode File Not Exists."))
				continue
			}
			M.ShellcodeFile = ShellcodeFile
			M.ShellcodeBuffer = utils.BinToHex(ShellcodeFile)
			break
		}
	}
	fmt.Println(utils.ColorPrint(0, "ShellcodeFile ==> "+M.ShellcodeFile))
	fmt.Println(utils.ColorPrint(0, "ShellcodeBuffer ==> "+M.ShellcodeFile))

}
func (M *Meta) SetTargetPort() {
	TargetPort := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Set TargetPort:"))
		fmt.Scanf("%s\n", &TargetPort)
		results, err := strconv.Atoi(TargetPort)
		if err != nil {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input"))
			continue
		}
		if results < 0 || results > 65535 {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input"))
			continue
		}
		M.TargetPort = TargetPort
		break
	}
	fmt.Println(utils.ColorPrint(0, "TargetPort ==> "+M.TargetPort))
}
func (M *Meta) SetTargetIp() {
	TargetIp := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Set TargetIp:"))
		fmt.Scanf("%s\n", &TargetIp)
		if !utils.IPChecker(TargetIp) {
			fmt.Println(utils.ColorPrint(-1, "Wrong IP Address."))
			continue
		} else {
			M.TargetIp = TargetIp
			break
		}
	}
	fmt.Println(utils.ColorPrint(0, "TargetIp ==> "+M.TargetIp))
}
func (M *Meta) SetNetworkTimeout() {
	NetworkTimeout := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Set NetworkTimeout:"))
		fmt.Scanf("%s\n", &NetworkTimeout)
		results, err := strconv.Atoi(NetworkTimeout)
		if err != nil {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input"))
			continue
		}
		if results < 1 {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input"))
			continue
		}
		M.NetworkTimeout = NetworkTimeout
		break
	}
	fmt.Println(utils.ColorPrint(0, "NetworkTimeout ==> "+M.NetworkTimeout))
}
func (M *Meta) SetDLLPath() {
	DLLPath := ""
	fmt.Printf(utils.ColorPrint(1, "Set DLLPath(%s):", M.DLLPath))
	fmt.Scanf("%s\n", &DLLPath)
	if DLLPath == "" {
		fmt.Println(utils.ColorPrint(0, "DLLPath ==> "+M.DLLPath))
		return
	} else {
		M.DLLPath = DLLPath
	}
	fmt.Println(utils.ColorPrint(0, "DLLPath ==> "+M.DLLPath))
}
func (M *Meta) SetArch() {
	arch := -1
	fmt.Printf(utils.ColorPrint(1, "0.x86\n1.x64\nSeletc Arch(Default is %s):", M.Arch))
	fmt.Scanf("%d\n", &arch)
	if arch == -1 {
		fmt.Println(utils.ColorPrint(0, "Arch ==> "+M.Arch))
		return
	}
	if arch < 0 || arch > len(archs) {
		fmt.Println(utils.ColorPrint(0, "Arch ==> "+M.Arch))
		return
	}
	M.Arch = archs[arch]
	fmt.Println(utils.ColorPrint(0, "Arch ==> "+M.Arch))
}
func (M *Meta) SetTarget(exp string) {
	if exp == "eb" {
		target := 1
		fmt.Printf(utils.ColorPrint(1, fmt.Sprintf("0) XP\n1) WIN72K8R2\nSelect Target(%s):", eb_targets[target])))
		fmt.Scanf("%d\n", &target)
		M.Target = eb_targets[target]
	}
	fmt.Println(utils.ColorPrint(0, "Target ==> "+M.Target))
}
func (M *Meta) SetFunction() {
	function := 1
	for {
		fmt.Printf(utils.ColorPrint(1, fmt.Sprintf("0.OutputInstall\n1.Ping\n2.RunDLL\n3.RunShellcode\n4.Uninstall\nSeletc Function(%s):", functions[function])))
		fmt.Scanf("%d\n", &function)
		if function < 0 || function > len(functions) {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input"))
			continue
		}
		M.Function = functions[function]
		break
	}
	fmt.Println(utils.ColorPrint(0, "Function ==> "+M.Function))
}
func (M *Meta) SetOutputInstall() {
	OutputInstall := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Input OutputInstall Path(%s):", M.OutputInstall))
		fmt.Scanf("%s\n", &OutputInstall)
		if (M.OutputInstall == "" && OutputInstall == "") || !utils.CheckPathLegal(OutputInstall) {
			fmt.Println(utils.ColorPrint(-1, "[-] Wrong OutputInstall Path."))
			continue
		}
		if !(OutputInstall == "" && M.OutputInstall != "") {
			M.OutputInstall = OutputInstall
		}
		_, err := os.Stat(M.OutputInstall)
		if os.IsNotExist(err) {
			fmt.Println(utils.ColorPrint(-1, "Shellcode File Not Exists."))
			continue
		}
		M.ShellcodeFile = OutputInstall
		M.ShellcodeBuffer = utils.BinToHex(OutputInstall)
		break
	}
	fmt.Println(utils.ColorPrint(0, "OutputInstall ==> "+M.OutputInstall))
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
		fmt.Println(utils.ColorPrint(0, fmt.Sprintf("%v ==> %v", key, val)))
	}
}
func (M *Meta) SetEnableSSL() {
	EnableSSL := ""
	for {
		fmt.Print(utils.ColorPrint(1, "Set EnableSSL(default is false):"))
		fmt.Scanf("%s\n", &EnableSSL)
		EnableSSL = strings.ToLower(EnableSSL)
		if EnableSSL == "" {
			EnableSSL = "false"
			break
		} else if EnableSSL == "true" {
			EnableSSL = "true"
			break
		} else {
			fmt.Println(utils.ColorPrint(-1, "Wrong Input,Input True or False or Stay Default."))
			continue
		}
	}
	M.IIS.EnableSSL = EnableSSL
	fmt.Println(utils.ColorPrint(0, "EnableSSL ==> "+M.IIS.EnableSSL))
}
func (M *Meta) SetIisPort() {
	IisPort := ""
	for {
		fmt.Print(utils.ColorPrint(1, "Set IisPort(default is 80):"))
		fmt.Scanf("%s\n", &IisPort)
		if IisPort == "" {
			IisPort = "80"
			break
		} else {
			port, err := strconv.Atoi(IisPort)
			if err != nil {
				fmt.Println(utils.ColorPrint(-1, "[-] Wrong Input,Input Iis Port Number or Stay Default."))
				continue
			}
			if port < 0 || port > 65535 {
				fmt.Println(utils.ColorPrint(-1, "[-] Wrong Input,Input Iis Port Number Out Of Range."))
				continue
			}
			break
		}
	}
	M.IIS.Port = IisPort
	fmt.Println(utils.ColorPrint(0, "IisPort ==> "+M.IIS.Port))
}
func (M *Meta) SetHostString() {
	HostString := ""
	for {
		fmt.Printf(utils.ColorPrint(1, "Set HostString(default is TargetIP %s):", M.TargetIp))
		fmt.Scanf("%s\n", &HostString)
		if HostString == "" {
			HostString = M.TargetIp
			break
		}
	}
	M.IIS.HostString = HostString
	fmt.Println(utils.ColorPrint(0, "HostString ==> "+M.IIS.HostString))
}
