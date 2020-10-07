package main

import (
    "flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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
	var pass []string
	flagNumPass := flag.Int("n", 0, "Numero de senhas a ser gerado. (entre 1 - 100000)")
	flagTamPass := flag.Int("t", 7, "Tamanho da senha a ser gerado. (entre 8 - 25)")
	flagHelp := flag.Bool("h", false, "Guia de Ajuda.")
	flag.Parse()

	if *flagNumPass >= 1 && *flagTamPass >= 8 && *flagNumPass <= 100000 && *flagTamPass <= 25 {
		for i := 0; i < *flagNumPass; i++{
			pw := GerarPass(*flagTamPass)
			pass = append(pass,pw)
		}
		for pos,psw := range pass{
			fmt.Printf("%dº Senha -> %s\n", pos+1, psw)
		}
	}

	if *flagHelp {
		fmt.Println("-n  -> numero de senhas a ser gerado. (entre 1 - 100000).")
		fmt.Println("-t -> Tamanho da senha a ser gerado. (entre 8 - 25).")
	}
}
