package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op) //相当于是报错
	}
	return result
}

func grade(score int) string {
	var g string = ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score :%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g

}

func main() {
	const filename = "abc.txt"
	//在if的条件里就可以赋值
	//if contens,err:=ioutil.ReadFile(filename);err!=nil{//错误如果不为空，就是有错误
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s",contens)
	//}

	contens, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contens)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(-3),
	)
}
