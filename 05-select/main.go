package main

import (
	"time"
	"fmt"
)

func main() {
	strings := make(chan string)
	nums := make(chan int)

	go func(){
		for {
			strings <- "Hello"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func(){
		for {
			nums <- 1
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg := <-strings:
			fmt.Println(msg)
		case msg := <-nums:
			fmt.Println(msg)
		}
	}
}