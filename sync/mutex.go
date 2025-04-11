package main

import (
	"fmt"
	"sync"
)

var Contador = 0

func Main2() {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//Ajustar el recurso Contador con un mutex
			mtx.Lock()
			Contador++
			mtx.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println(Contador)
}
