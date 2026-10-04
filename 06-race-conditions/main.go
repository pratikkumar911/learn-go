package main

import (
	"fmt"
	"sync"
)

type database struct{
	users map[int]string
	sync.RWMutex
}

func (db *database) add(i int){
	db.Lock()
	defer db.Unlock()
	db.users[i] = fmt.Sprintf("user-%d", i)
}

func main() {
	db := database{
		users: make(map[int]string),
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i:=0;i<100;i++ {
			db.add(i)
		}
	}()

	go func(){
		defer wg.Done()
		for i:=0;i<100;i++ {
			db.add(i)
		}
	}()

	wg.Wait()
	fmt.Println(db.users)
}