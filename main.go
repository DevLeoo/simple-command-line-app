package main

import (
	"log"
	"os"
	"simple-command-line-app/app"
)

func main() {

	aplicacao := app.Generate()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
