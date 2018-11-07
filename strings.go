package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "Yes我爱慕课网!"    //每个中文字节占3位，utf-8
	fmt.Println(len(s)) //获得的是字节数量
	//使用[]byte获得所有的字节
	//使用[]rune可以获得转换好的字符数组
	//其他方法都在strings里面

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune= int32会转换成unicode
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //或者真的字符数量

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c) ", i, ch)
	}
	fmt.Println()
	c := strings.LastIndex(s, "课")
	fmt.Println(c)
}
