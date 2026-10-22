package main

import (
	"fmt"
	"sync"
)

func main() {
	const cycle = 5

	chA := make(chan struct{})
	chB := make(chan struct{})
	chC := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)

	go func(){
		defer wg.Done()
		for i:=0;i<cycle;i++{
			<-chA
			fmt.Print("A")
			chB<- struct{}{}
		}
	}()

	go func(){
		defer wg.Done()
		for i:=0;i<cycle;i++{
			<-chB
			fmt.Print("B")
			chC<- struct{}{}
		}
	}()

	go func(){
		defer wg.Done()
		for i:=0;i<cycle;i++{
			<-chC
			fmt.Print("C")
			if i<cycle-1{
				chA<- struct{}{}
			}
		}
	}()

	chA<- struct{}{}
	wg.Wait()
}