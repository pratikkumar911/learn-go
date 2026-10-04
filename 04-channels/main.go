package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sendMessage(ch chan<- string, num int, wg *sync.WaitGroup) {
	fmt.Printf("Sending message %d...\n", num)
	
	time.Sleep(time.Second * time.Duration(num))	

	ch <- fmt.Sprintf("Message %d send!", num)

	wg.Done()
}

func receiveMessage(msgs <-chan string){
	fmt.Println("Waiting for message")

	for msg := range msgs {
		fmt.Println("Received :", msg)
	}
}

func main() {
	msgs := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go sendMessage(msgs, 2, &wg)
	go sendMessage(msgs, 3, &wg)

	go func() {
		receiveMessage(msgs)
	}()

	wg.Wait()

	close(msgs)

	log.Println("Done")
}
