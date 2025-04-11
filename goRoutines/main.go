package main

import (
	"fmt"
	"time"
)

func Main() {
	go worker()

	time.Sleep(time.Second)

}

func worker() {
	fmt.Println("->working.. ")
}
