package main

import (
	//"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func MyIp() {
	so := runtime.GOOS
	if so == "windows" {
		eu := exec.Command("cmd", "/c", "ipconfig", "/all")
		eu.Stdout = os.Stdout
		eu.Run()
	} else if so == "linux" {
		eu := exec.Command("ip", "addr", "show")
		eu.Stdout = os.Stdout
		str := eu.Run()

	}
}

func main() {
	MyIp()
	fmt.Println("Ok")
}
