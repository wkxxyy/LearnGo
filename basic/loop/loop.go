package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	} else {
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
		return result
	}

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		printFileContents(file)
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { //省略了初始条件和递增条件
		fmt.Println(scanner.Text())

	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {

	fmt.Println(
		convertToBin(5),
		convertToBin(13), //1001--->1101
		convertToBin(72387885),
		convertToBin(0))
	printFile("basic/branch/abc.txt")

	s := `abc"d"
	kkk
	123
	p`

	printFileContents(strings.NewReader(s))
	printFileContents(bytes.NewReader([]byte(s)))

	//forever()

}
