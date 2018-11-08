package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Work %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int

	tm := time.After(10 * time.Second) //十秒后会给这个tm发送一个时间
	tick := time.Tick(time.Second)     //定时

	for {

		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select { //如果想用非阻塞的channel，就用select加default,谁先来就打印谁

		case n := <-c1: //就是如果n收到了c1就执行
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //这个是超时了
			fmt.Println("Time out")
		case <-tick: //这个设置定时
			fmt.Println("queue len = ", len(values))
		case <-tm: //设置定时关机
			fmt.Println("bye")
			return

		}
	}
}
