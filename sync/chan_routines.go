package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i,&wg, ch)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println()
		fmt.Println(v)

	}
	fmt.Println("la ejecución termino")
}

func worker(va int, wg *sync.WaitGroup,  ch chan int) {
	wg.Done()
	fmt.Println("working...")
	ch <- va
	time.Sleep(1 * time.Second)
}
