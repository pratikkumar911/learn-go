package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch1
			fmt.Println(i)
			if i < 10 {
				ch2 <- true
			}
		}
	}()


	ch2 <- true

	wg.Wait()
}