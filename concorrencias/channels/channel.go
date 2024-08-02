package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}

	// esse codigo faz a mesma coisa que o de cima
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	close(canal)
}
