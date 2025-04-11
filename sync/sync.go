package main

import (
	"fmt"
	"sync"
)

func Main1() {
	//Usa WaitGroup para esperar 5 goroutines que impriman sus índices.
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i <= 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println()
			fmt.Printf("indice:%v",i)
		}()

	}

	wg.Wait()
}
