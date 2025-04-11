package main

import (
	"context"
	"fmt"
	"time"
)
func Operacion( ctx context.Context){
	for i:= 0; i< 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work")
		case <-ctx.Done():
			fmt.Println("-Funcion cancelad-")
			return
		}
	}
	

}

func main6() {

	ctxHijo, CancelFunc := context.WithCancel(context.Background())

	go Operacion(ctxHijo)
	time.Sleep(3 * time.Second)
	CancelFunc()
	time.Sleep(3 * time.Second)
}
