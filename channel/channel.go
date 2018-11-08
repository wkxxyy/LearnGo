package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		//n,ok:=<-c//用这个来判断是否接受完了
		//if !ok{
		//	break
		//}
		fmt.Printf("Work %d received %d\n", id, n)

	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)

	go worker(id, c)

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}

	//var c chan int // C == nil
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //这个3是个缓冲数，可以发3个但是不真的发出去
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)

}

func channleClose() {
	//发送方来发送close
	c := make(chan int, 3) //这个3是个缓冲数，可以发3个但是不真的发出去
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //关闭c，告诉接收方我没有要发的了
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channleClose()
}
