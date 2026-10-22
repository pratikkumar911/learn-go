package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks{
		fmt.Printf("worker %d processing task %d\n", id, task)
		time.Sleep(1*time.Second)
		fmt.Printf("worker %d finished task %d\n", id, task)
	}
}

func main() {
	const totalTasks = 10

	task := make(chan int, 5)

	var wg sync.WaitGroup

	for i:=1;i<=3;i++{
		wg.Add(1)
		go func(){
			worker(i,task,&wg)
		}()
	}

	for i:=1;i<=totalTasks;i++{
		fmt.Println("Adding task", i)
		task<- i
	}

	close(task)
	wg.Wait()
}