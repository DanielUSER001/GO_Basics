package main

import (
	"context"
	"fmt"
	"time"
)
func Operaciones( ctx context.Context){
	for i:= 0; i< 5; i++ {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Doing some work")
	case <-ctx.Done():
		fmt.Println("-Hecho-")
		return
	}
	}
	

}

func Main2() {

	ctxHijo,c := context.WithTimeout(context.Background(), 3*time.Second)
	c() // cancelar func c

	go Operaciones(ctxHijo)
	time.Sleep(6 * time.Second)
}
