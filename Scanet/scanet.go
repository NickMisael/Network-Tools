package main

import (
	"fmt"
	"os"
	"os/exec"
)

func MyIp(so string) {
	if so == "1"{
		eu := exec.Command("ipconfig /all")
	} else {
		eu := exec.Command("ip addr show")
	}
}

func main() {
	var esc string
	fmt.Printf("Digite seu Sistema: 1- Windows 2- linux")
	fmt.Scanf("%s", &esc)
	if es, err := strconv.Atoi(esc); err != nil or es < 1 or es > 2{
		fmt.Println("Número Inválido!!")
	}
	fmt.Println()
}
