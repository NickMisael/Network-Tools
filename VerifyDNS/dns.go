// dns.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func Menu() {
	fmt.Println("\t| 1 - Verificar")
	fmt.Print("\t| -> ")
}

type DNS struct {
	Nome  string
	Addr  string
	Addr2 string
	Tempo float32
}

func Verificar() {
	const (
		CONN_TYPE = "tcp"
		CONN_PORT = ":53"
	)
	d := []DNS{
		DNS{
			Nome:  "OpenDns",
			Addr:  "208.67.222.222",
			Addr2: "208.67.220.220",
		},
		DNS{
			Nome:  "Cloudflare",
			Addr:  "1.1.1.1",
			Addr2: "1.0.0.1",
		},
		DNS{
			Nome:  "Google",
			Addr:  "8.8.8.8",
			Addr2: "8.8.4.4",
		},
		DNS{
			Nome:  "Norton",
			Addr:  "199.85.126.10",
			Addr2: "199.85.127.10",
		},
		DNS{
			Nome:  "Verisign",
			Addr:  "64.6.64.6",
			Addr2: "64.6.65.6",
		},
		DNS{
			Nome:  "NuSEC",
			Addr:  "8.26.56.26",
			Addr2: "8.20.247.20",
		},
		DNS{
			Nome:  "Quad9",
			Addr:  "9.9.9.9",
			Addr2: "9.9.9.9",
		},
		DNS{
			Nome:  "Neustar",
			Addr:  "156.154.70.5",
			Addr2: "156.154.71.5",
		},
		DNS{
			Nome:  "SafeDNS",
			Addr:  "195.46.39.39",
			Addr2: "195.46.39.40",
		},
		DNS{
			Nome:  "Yandex",
			Addr:  "77.88.8.8",
			Addr2: "77.88.8.1",
		},

		DNS{
			Nome:  "Level3",
			Addr:  "209.244.0.3",
			Addr2: "",
		},

		/*DNS{
			Nome:  "SmartViper",
			Addr:  "208.76.50.50",
			Addr2: "",
		},*/
		DNS{
			Nome:  "Dyn",
			Addr:  "216.146.35.35",
			Addr2: "",
		},
		DNS{
			Nome:  "DNS Watch",
			Addr:  "84.200.69.80",
			Addr2: "84.200.70.40",
		},
	}
	for j := 0; j < 10; j++ {
		for i := 0; i < len(d); i++ {
			in := time.Now().UnixNano()
			Address := d[i].Addr + CONN_PORT
			conn, err := net.Dial(CONN_TYPE, Address)
			if err != nil {
				fmt.Println("\tErro:", i)
				continue
			}
			conn.Close()
			fi := time.Now().UnixNano()
			tF := fi - in
			d[i].Tempo += float32(tF) / 1000000
		}
	}

	for y := 0; y < len(d); y++ {
		d[y].Tempo /= 10
	}
	for i := 0; i < len(d)-1; i++ {
		for y := 0; y < len(d)-1; y++ {
			if d[y].Tempo >= d[y+1].Tempo {
				d[y].Nome, d[y+1].Nome = d[y+1].Nome, d[y].Nome
				d[y].Addr, d[y+1].Addr = d[y+1].Addr, d[y].Addr
				d[y].Addr2, d[y+1].Addr2 = d[y+1].Addr2, d[y].Addr2
				d[y].Tempo, d[y+1].Tempo = d[y+1].Tempo, d[y].Tempo
			}
		}
	}
	fmt.Printf("\a")
	fmt.Printf("\t _________________________________________\n")
	fmt.Printf("\t|                                         |\n")
	fmt.Printf("\t|       Host | DNS Server | Time ms       |\n")
	fmt.Printf("\t|_________________________________________|\n")
	fmt.Printf("\t|                                         |\n")
	for i := 0; i < len(d); i++ {
		fmt.Printf("\t|-----------------------------------------|\n")
		fmt.Printf("\t| %s \n", d[i].Nome)
		fmt.Printf("\t| P: %s \n", d[i].Addr)
		fmt.Printf("\t| S: %s \n", d[i].Addr2)
		fmt.Printf("\t| %.2f milisegundos \n", d[i].Tempo)
		if i == len(d)-1 {
			fmt.Printf("\t|_________________________________________|\n")
		}
	}

}

func main() {
	for {
		var esc string
		Menu()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			esc = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			fmt.Println("Algo deu errado!!")
		}
		if esc == ":q!" {
			break
		}
		if es, err := strconv.Atoi(esc); err != nil {
			fmt.Println("Erro: Número inválido!!")

		} else {
			switch es {
			case 1:
				Verificar()

			}
			break
		}
		time.Sleep(time.Second + 2)
	}
	fmt.Println("")
}