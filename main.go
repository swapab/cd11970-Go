package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func receive(receive bool, anotherDone chan bool) {
	if receive == true {
		fmt.Println("from receive")
		anotherDone <- true
	}
}

func main() {
	start := time.Now()
	done := make(chan bool)
	anotherDone := make(chan bool)
	go hello(done)
	toSend := <-done
	go receive(toSend, anotherDone)
	<-anotherDone
	fmt.Println("main function")
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}
