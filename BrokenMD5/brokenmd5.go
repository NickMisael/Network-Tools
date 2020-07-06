package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var alfabeto = [52]string{"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "I", "i", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "O", "o", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "Z"}

func CalcTime(inicio, final int64) (segundo float32) {
	total := final - inicio
	segundo = float32(total) / 1000000000.0
	return
}

func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func Comparar(word, hash string) string {
	if chash := Gerar(word); chash == hash {
		return word
	}
	return ""
}

func Discovery(hash string, result chan string, velo chan int) {
	// Definindo variáveis
	var word string
	var ver bool
	tamalfa := len(alfabeto)
	tam := 1
	count := 0
	inicio := time.Now().UnixNano()

	// For repsonsável por juntar letras
	for {
		// Gera um número aleatório
		pos := Random(0, 51)

		// Concatena uma letra a palavra
		word += alfabeto[pos]

		// Armazena o tamanho da palavra
		tamword := len(word)

		// Verifica se a palavra passou do tamanho atual
		if tamword > tam {
			word = ""
			// Se for igual ao tamanho
		} else if tam == tamword {
			// Verifica se o hash da palavra gerada é o mesmo proposto à inicio
			if h := Comparar(word, hash); h != "" {
				//  em caso veridico encerra o for infinito
				result <- word
				ver = true
				break
			}
			// Se não adiciona um ao contador
			count += 1
		}
		// se o contador for maior que o dobro do tamanho do alfabeto vezes o tamanho atual da palavra
		if count > (tamalfa*tam)*2 {
			// Zera o contador e acrescenta 1 ao tamanho da palavra
			count = 0
			tam += 1
		}
		// Se o tamanho da palavra for maior que 5
		if tam > 5 {
			// Reinicia o tamanho
			tam = 1
		}
	}

	if ver == true {
		// Calcula o tempo final
		final := time.Now().UnixNano()

		// Calcula o tempo decorrido
		fmt.Printf("%f Segundos Decorridos\n", CalcTime(inicio, final))
	}
}

func Gerar(word string) (hash string) {
	h := md5.New()
	io.WriteString(h, word)
	hash = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func main() {
	var palavra string
	speed := make(chan int, 1000)
	result := make(chan string)
	fmt.Printf("Digite uma palavra: ")
	fmt.Scanf("%s", &palavra)
	cod := Gerar(palavra)
	fmt.Println("Codifica ->", cod)
	for i := 0; i < cap(speed); i++ {
		go Discovery(cod, result, speed)
	}
	r := <-result
	close(result)
	fmt.Println("Decodifica ->", r)
}
