package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waiGroup sync.WaitGroup

	waiGroup.Add(2) // quantidade de go routines

	go func() {
		escrever("Olá Mundo")
		waiGroup.Done() // -1 no count
	}()

	go func() {
		escrever("Programando em Go!")
		waiGroup.Done()
	}()

	waiGroup.Wait() // esperar a quantidade das go routines chegar em zero
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
