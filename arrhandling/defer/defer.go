package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() { //defer里面有个栈，加了defer，之后有return何panic都不怕
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s ,%s \n", pathError.Op, pathError.Path, pathError.Err)
		}
		//fmt.Println("Errot:",err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) ////你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
