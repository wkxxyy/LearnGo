package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2 //0,1,2,3,4,5,6,7
	fmt.Println("arr=", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1)) //2,3,4,5
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2)) //5,6
	// 能够吧6取出来，因为他们其实知道s1后面还有两个数组，切面有一个值叫做cap
	//s[i]不能超过len(s)，向后拓展不能超过底层数组cap(s)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	//s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)
	//由于值传递的关系，必须要接受append
}