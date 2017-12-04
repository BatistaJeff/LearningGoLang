package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	texto, versao := exibeIntroducao()
	fmt.Println(texto, "versao ->", versao)

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando não identificado")
			os.Exit(-1)
		}
	}
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func exibeIntroducao() (string, float64) {
	nome := "Boa noite Sr(a)."
	versao := 1.1
	return nome, versao
}

func exibeMenu() {
	fmt.Println("Escolha uma função")
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair")
}

func iniciaMonitoramento() {
	fmt.Println("Iniciando Monitoramento")

	sites := []string{"https://www.alura.com.br",
		"http://random-status-code.herokuapp.com",
		"https://www.google.com.br"}

	for i, site := range sites {

		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println(i, "--> O site: ", site, " --> Foi carregado <--")
		} else {
			fmt.Println(i, "--> O site: ", site, " --> Deve estar com problemas <--")
		}

	}

}
