package main

import (
	"academia/src/config"
	"academia/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Println(config.Porta, config.StringConexaoBanco)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	fmt.Sprintln("Rodando na porta 5000")
}
