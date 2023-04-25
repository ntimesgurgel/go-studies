package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	var comando int = leComando()

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo")
		os.Exit(0)
	} else {
		fmt.Println("Comando nao reconhecido")
		os.Exit(-1)
	}

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo")
	default:
		fmt.Println("Comando nao reconhecido")
	}

	time.Sleep(5 * time.Second)
}

func exibeIntroducao() {
	const nome string = "Nathan"
	const versao float32 = 1.1
	fmt.Println("Olá, sr. ", nome)
	fmt.Println("Este programa esta na versao", versao)
	fmt.Println("O tipo da variavel eh", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1 - Exibir Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando")
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i, site := range sites {
		fmt.Println("testando site ", i, " :", site)
		testaSite(site)
	}
}

func exibeNomes() {
	nomes := []string{"Nathan", "Rayanne"}
	nomes = append(nomes, "Jorge")
	fmt.Println(nomes)
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas.\n Status code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Ocorreu um erro")
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if error != nil {
			fmt.Println("Ocorreu um erro:", error)
		}

		fmt.Println(linha)

		if error == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(site + "- online:" + strconv.FormatBool(status) + "hora: " + time.Now().Format("02/01/2006 15:04:05") + "\n")

	fmt.Println(arquivo)
}
