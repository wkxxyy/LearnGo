package main

import "fmt"

type cmd interface {
	good() string
}

type real struct {
	Contents string
}

func (c real) good() string {
	return c.Contents
}

func main() {

	cmos := real{}
	cmos.Contents = "123456"

	fmt.Println(cmos.good())

}
