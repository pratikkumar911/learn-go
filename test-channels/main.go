package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.IntN(100)
}

func main(){
	datachan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i:=0;i<999;i++{
			wg.Add(1);
			go func() {
				defer wg.Done()
				result:= DoWork()
				datachan <- result
			}()
		}
		wg.Wait()
		close(datachan)
	}()

	for n := range datachan {
		fmt.Printf("n = %d\n", n)
	}
}