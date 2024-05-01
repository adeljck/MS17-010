package utils

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

func Execute(cmd *exec.Cmd) {
	defer fmt.Println(color.HiGreenString("[+] Execute Finish."))
	confirm := ""
	fmt.Printf(color.HiBlueString("Confirm To Execute Command(y):"))
	fmt.Scanf("%s\n", &confirm)
	confirm = strings.ToLower(confirm)
	if confirm == "n" {
		return
	}
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
	stderr, err := cmd.StderrPipe()
	defer stderr.Close()
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
	outScanner := bufio.NewScanner(stdout)
	go func() {
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error:", err)
	}
	return
}
