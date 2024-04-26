package metatdatas

import (
	"eternal/utils"
	"fmt"
	"log"
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
var archs = []string{"x86", "x64", "unknown"}
var eb_targets = []string{"XP", "WIN72K8R2"}
var functions = []string{"OutputInstall", "Ping", "RunDLL", "RunShellcode", "Uninstall"}

func SetDefaultMetaData() {
	for {
		fmt.Print("Set Default Target:")
		TargetIp := ""
		fmt.Scanf("%s\n", &TargetIp)
		if !utils.IPChecker(TargetIp) {
			fmt.Println("Wrong IP Address....")
		} else {
			MetaData.TargetIp = TargetIp
			MetaData.IIS.HostString = TargetIp
			log.SetPrefix("[*] ")
			log.Println("TargetIp ==> ", MetaData.TargetIp)
			log.Println("HostString ==> ", MetaData.TargetIp)
			MetaData.ProjectName = fmt.Sprintf("z%s", MetaData.TargetIp)
			break
		}
	}
	LogDir := ""
	for {
		fmt.Print("Set Base Log Dir(Default ./logs/):")
		fmt.Scanf("%s\n", &LogDir)
		if utils.CheckPathLegal(LogDir) {
			break
		}
	}
	if LogDir != "" {
		MetaData.LogDir = LogDir
	}
	log.SetPrefix("[*] ")
	log.Println("LogDir ==> ", MetaData.LogDir)
	utils.CreateLogDir(MetaData.LogDir)
	Projects := utils.ListDirectory(MetaData.LogDir)
	ProjectIndex := 0
	for {

		for index, Project := range Projects {
			fmt.Printf("%d.%s\n", index, Project)
		}
		fmt.Print("Set Project Name(default 0):")
		fmt.Scanf("%d\n", &ProjectIndex)
		if ProjectIndex < 0 || ProjectIndex > len(Projects) {
			fmt.Println("Wrong Project Index.")
		} else {
			break
		}
	}
	ProjectName := ""
	if ProjectIndex == 0 {
		fmt.Printf("Set Project Name(default %s):", fmt.Sprintf("z%s", MetaData.TargetIp))
		fmt.Scanf("%s\n", &ProjectName)
		if ProjectName != "" {
			MetaData.ProjectName = ProjectName
		}
	} else {
		MetaData.ProjectName = Projects[ProjectIndex]
	}
	log.SetPrefix("[*] ")
	log.Println("Project ==> ", fmt.Sprintf("%s", MetaData.ProjectName))
	utils.CreateProject(MetaData.LogDir, MetaData.ProjectName)
	log.SetPrefix("[*] ")
	log.Println("Set Target Directory ==> ", fmt.Sprintf("z%s", MetaData.TargetIp))
	utils.CreateProject(filepath.Join(MetaData.LogDir, MetaData.ProjectName), fmt.Sprintf("z%s", MetaData.TargetIp))
	utils.CreateLog(MetaData.TargetIp, MetaData.LogDir, MetaData.ProjectName)
}
func (M *Meta) SetShellcodeFile() {
	ShellcodeFile := ""
	for {
		fmt.Printf("Set DOUP ShellcodeFile:")
		fmt.Scanf("%s\n", &ShellcodeFile)
		if ShellcodeFile == "" || !utils.CheckPathLegal(ShellcodeFile) {
			fmt.Println("Wrong Shellcode File Path ")
			continue
		} else {
			_, err := os.Stat(ShellcodeFile)
			if os.IsNotExist(err) {
				fmt.Println("Shellcode File Not Exists.")
				continue
			}
			M.ShellcodeFile = ShellcodeFile
			M.ShellcodeBuffer = utils.BinToHex(ShellcodeFile)
			break
		}
	}
	M.ShellcodeBuffer = utils.BinToHex(M.ShellcodeFile)
	fmt.Println("ShellcodeFile ==> " + M.ShellcodeFile)

}
func (M *Meta) SetTargetPort() {
	fmt.Printf("Set TargetPort:")
	fmt.Scanf("%s\n", &M.TargetPort)
	fmt.Println("TargetPort ==> " + M.TargetPort)
}
func (M *Meta) SetTargetIp() {
	TargetIp := ""
	for {
		fmt.Printf("Set TargetIp:")
		fmt.Scanf("%s\n", &TargetIp)
		if !utils.IPChecker(TargetIp) {
			fmt.Println("Wrong IP Address.")
			continue
		} else {
			M.TargetIp = TargetIp
			break
		}
	}
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
	fmt.Printf("0.x86\n1.x64\nSeletc Arch(%s):", M.Arch)
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
	}
	fmt.Println("Target ==> " + M.Target)
}
func (M *Meta) SetFunction() {
	function := 1
	fmt.Printf("0.OutputInstall\n1.Ping\n2.RunDLL\n3.RunShellcode\n4.Uninstall\nSeletc Function(%s):", functions[function])
	fmt.Scanf("%d\n", &function)
	M.Function = functions[function]
	fmt.Println("Function ==> " + M.Function)
}
func (M *Meta) SetOutputInstall() {
	OutputInstall := ""
	for {
		fmt.Printf("Input OutputInstall Path(%s):", M.OutputInstall)
		fmt.Scanf("%s\n", &OutputInstall)
		if OutputInstall == "" || !utils.CheckPathLegal(OutputInstall) {
			fmt.Println("Wrong OutputInstall Path.")
			continue
		}
		break
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
func (M *Meta) SetEnableSSL() {
	EnableSSL := ""
	for {
		fmt.Print("Set EnableSSL(default is false):")
		fmt.Scanf("%s\n", &EnableSSL)
		EnableSSL = strings.ToLower(EnableSSL)
		if EnableSSL == "" {
			EnableSSL = "false"
			break
		} else if EnableSSL == "true" {
			EnableSSL = "true"
			break
		} else {
			fmt.Println("Wrong Input,Input True or False or Stay Default.")
			continue
		}
	}
	M.IIS.EnableSSL = EnableSSL
	fmt.Println("EnableSSL ==> " + M.IIS.EnableSSL)
}
func (M *Meta) SetIisPort() {
	IisPort := ""
	for {
		fmt.Print("Set IisPort(default is 80):")
		fmt.Scanf("%s\n", &IisPort)
		if IisPort == "" {
			IisPort = "80"
			break
		} else {
			port, err := strconv.Atoi(IisPort)
			if err != nil {
				fmt.Println("Wrong Input,Input Iis Port Number or Stay Default.")
				continue
			}
			if port < 0 || port > 65535 {
				fmt.Println("Wrong Input,Input Iis Port Number Out Of Range.")
				continue
			}
			break
		}
	}
	M.IIS.Port = IisPort
	fmt.Println("IisPort ==> " + M.IIS.Port)
}
func (M *Meta) SetHostString() {
	HostString := ""
	for {
		fmt.Printf("Set HostString(default is TargetIP %s):", M.TargetIp)
		fmt.Scanf("%s\n", &HostString)
		if HostString == "" {
			HostString = M.TargetIp
			break
		}
	}
	M.IIS.HostString = HostString
	fmt.Println("HostString ==> " + M.IIS.HostString)
}
