package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) //如果有一个不用，就用下划线来接收
		return q, nil
	default:
		return 0, fmt.Errorf("unsupport operation：%s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int { //op为函数名，后面的是他的类型，func(int,int) int
	p := reflect.ValueOf(op).Pointer() //反射
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)

}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a

}
func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(eval(3, 4, "x"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	a := apply(func(a int, b int) int { return int(math.Pow(float64(a), float64(b))) }, 3, 4)
	fmt.Println(a)
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
