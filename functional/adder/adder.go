package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}
} //这个返回的时候，返回的不是一个单纯的函数，而是一个闭包，是一个环境，里面有维护的Sum等

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder { //这里base就是sum，v就是传进来的i

	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v) //最后的adder2返回的就是一个iAdder，所以可以用来返回
	}
}

func main() {
	a := adder2(0) //a现在是返回值，在这里就是func(int) int，这里在a创建的时候，已经维护了个sum了，所以不动，
	for i := 0; i < 10; i++ {

		fmt.Printf("0+1+...+ %d=%d\n", i, a(i))
	}
}
