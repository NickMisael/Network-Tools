package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var alfabeto = [52]string{"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "I", "i", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "O", "o", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "Z"}
var num = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var special = [37]string{`"`, `\`, "?", "|", "!", "@", "'", "#", "%", "¨", "*", "&", "*", "(", " ", ")", "-", "_", "+", "=", "§", ",", ".", "<", ">", ":", ";", "^", "~", "`", "/", "{", "}", "[", "]", "ª", "º"}

func limpaTela() {
	so := runtime.GOOS
	if so == "windows" {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	} else if so == "linux" {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	}
}

func CalcTempo(i, f int64) {
	total := float32((f - i) / 1000000000.0)
	fmt.Printf("%.3f Segundos Decorridos\n", total)
}

func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func GerarPass(tamf int) (string) {
	var pass string
	tami := 0
	for {
		if tami == tamf {
			break
		}
		var pos []int
		pos = append(pos, Random(0, 51))
		pos = append(pos, Random(0, 9))
		pos = append(pos, Random(0, 36))

		posP := Random(0, len(pos))
		if (posP+posP)%2 == 1 {
			pass += special[pos[2]]
		}
		if posP == 0 {
			pass += alfabeto[pos[posP]]
		} else if posP == 1 {
			pass += num[pos[posP]]
		} else {
			pass += special[pos[posP]]
		}
		tami += 1
	}
	return pass
}

func main() {
	var tamanho,num int
	var err error
	//senha := make(chan string, 1000)
	limpaTela()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Digite o número de senhas: ")
	for scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Digite um número!!")
			time.Sleep(time.Second + 3)
			limpaTela()
			fmt.Printf("Digite o número de senhas: ")
		} else {
			break
		}
	}
	if num <= 1 {
		num = 3
	}
	fmt.Printf("Digite o tamanho da senha: ")
	for scanner.Scan() {
		tamanho, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Digite um número!!")
			time.Sleep(time.Second + 3)
			limpaTela()
			fmt.Printf("Digite o tamanho da senha: ")
		} else {
			break
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	var pass []string
	if tamanho >= 8 {
		for i := 0; i < num; i++{
			pw := GerarPass(tamanho)
			pass = append(pass,pw)
		}
		for pos,psw := range pass{
			fmt.Printf("%dº Senha -> %s\n", pos+1, psw)
		}
	}
}
