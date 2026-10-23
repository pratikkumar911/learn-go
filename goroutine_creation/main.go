package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type WorkerPool struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func NewWorkerPool(workerCount int, jobBuffer int) *WorkerPool {
	pool := &WorkerPool{
		jobs: make(chan Job, jobBuffer),
	}
	pool.start(workerCount)
	return pool
}

func (p *WorkerPool) start(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		p.wg.Add(1)
		go func(workerId int) {
			defer p.wg.Done()
			for job := range p.jobs {
				process(job, workerId)
			}
		}(i)
	}
}

func (p *WorkerPool)submit(job Job){
	p.jobs<- job
}

func (p *WorkerPool) shutdown(){
	close(p.jobs)
	p.wg.Wait()
}

func process(job Job, workerId int) {
	fmt.Printf("Worker %d processing %d\n", workerId, job.ID)
	time.Sleep(time.Second)
}

func main() {
	pool := NewWorkerPool(3, 10)
	for i := 1; i <= 10; i++ {
		pool.submit(Job{ID: i})
	}
	pool.shutdown()
}
