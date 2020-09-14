// dns.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var so = runtime.GOOS

type DNS struct {
	Nome  string
	Addr  string
	Addr2 string
	Tempo float32
}
var d = []DNS{
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

func Verificar() {
	for j := 0; j < 10; j++ {
		for i := 0; i < len(d); i++ {
			in := time.Now().UnixNano()
			Address := d[i].Addr + ":53"
			conn, err := net.Dial("tcp", Address)
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

func trocar() {
	if so == "windows" {
		fmt.Println("Windows")
	} else if so == "linux" {
		sudo := exec.Command("sudo echo")
		sudo.Stdout = os.Stdout
		sudo.Run()
		_, err := os.Create("/etc/resolv.conf")
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("# From DHCP\nnameserver %s\nnameserver %s\n", d[0].Addr, d[0].Addr2)
		err = ioutil.WriteFile("resolv.conf", []byte(str), 0644)
		if err != nil {
			panic(err)
		}
		cp := exec.Command("sudo cp resolv.conf /etc/resolv.conf")
		cp.Stdout = os.Stdout
		cp.Run()
		rm := exec.Command("rm -rf resolv.conf")
		rm.Stdout = os.Stdout
		rm.Run()
	}
}

func main() {
	
	flagStr := flag.Bool("V", false, "Verificar a melhor opção DNS")
	flagtro := flag.Bool("T", false, "Verificar a melhor opção DNS")
	flag.Parse()
	if *flagStr {
		Verificar()
	}
	if *flagtro {
		trocar()
	}

}
