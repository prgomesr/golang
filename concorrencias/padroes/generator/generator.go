package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return c
}
