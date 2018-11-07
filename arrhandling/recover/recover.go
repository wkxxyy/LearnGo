package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover() //获取一个错误
		if err, ok := r.(error); ok {
			fmt.Println("Error occured: ", err)
		} else {
			panic(fmt.Sprintf("I don't know waht to do：%v", r))
		}
	}()
	//panic(errors.New("This is an error"))
	panic(123)
}

func main() {
	tryRecover()
}
