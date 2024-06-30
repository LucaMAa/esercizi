package main

import (
	"esercizi/service"
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "1":
		service.FirstFilter()
	case "2":
		service.SecondFilter()
	case "3":
		service.VotoService()
	case "4":
		service.Matrice()
	default:
		fmt.Println("Comando non trovato")
	}
}
