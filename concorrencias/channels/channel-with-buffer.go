package main

import "fmt"

func main() {
	canal := make(chan string, 2) // buffer com capacidade de 2 entradas

	canal <- "Olá Mundo"
	canal <- "Fim do Programa!"

	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
