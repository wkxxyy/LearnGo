package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { //省略了初始条件和递增条件
		fmt.Println(scanner.Text())
	}
}

//1,1,2,3,5,8,13
//a,b
//	a,b
//	  a b
func fibonacci() intGen {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int //是个类型就能实现接口

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	//TODO:incorrect if p is too small
	s := fmt.Sprintf("%d\n", next) //要把数据写进p
	return strings.NewReader(s).Read(p)
}

func main() {
	f := fibonacci()

	printFileContents(f)

}
