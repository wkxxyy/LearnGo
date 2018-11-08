package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Work %d received %c\n", id, n)
		//通过通信
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w)

	return w
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup //先生成一个WaitGroup
	wg.Add(20)            //有多少个任务

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)

	}
	for i, worker := range workers {
		worker.in <- 'a' + i

	}
	for i, worker := range workers {
		worker.in <- 'A' + i

	}
	wg.Wait() //等待程序发送给我说，接受完成了，当调用wg.Done()时候完成

}

func main() {
	chanDemo()
}
