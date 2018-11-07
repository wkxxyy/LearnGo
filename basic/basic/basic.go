package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa int    = 3
	ss string = "kkk"
	bb bool   = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d  %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, s, c)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	//var c complex128=3+4i
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	//cmplx.Pow(math.E,1i*math.Pi)+1)//底数是E，

}

func triangle() {
	var a, b int = 3, 4

	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename string = "abc.txt"
		a        int    = 3
		b        int    = 4
	)
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
		javascript
	)
	//b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
func main() {
	fmt.Println("hello,world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()

}
