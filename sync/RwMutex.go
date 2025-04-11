package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	shareData = 0
	rwMutex    sync.RWMutex
	wg sync.WaitGroup
)

func Main3() {
	//lanzamos 3 Lectores
	for i:=1; i<= 3; i++{
		wg.Add(1)
		go reader(i, &wg)
	}

	//lanzamos 1 Escrito
	for i:=1; i<= 1; i++{
		wg.Add(1)
		go writer(i, &wg)
	}

	wg.Wait()
	fmt.Println("La ejecucion termino")

}

func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock()
	shareData++
	fmt.Printf("Escritor %d escribió: %d\n", id, shareData)
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()

}
func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Lector %d lee: %d\n", id, shareData)
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()

}
