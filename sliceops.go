package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

func main() {

	var s []int //zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	var s2 []int = make([]int, 16) //创建一个切片，但是大小不知道，长度16
	s2[0] = 12
	s3 := make([]int, 10, 32) //给大小10，cap为32
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //把s1复制给s2,不替换，之覆盖
	printSlice(s2)
	fmt.Println("Delteing elements from slice")
	s2 = append(s2[:3], s2[4:]...) //因为后面是可变参数，所以要在s2【4：】后面加上三个...来表示为可变参数
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)

	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
