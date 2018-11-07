package main

import "fmt"

var lastOccurred = make([]int, 0xffff)

func lenthOfNoneRepeatingSubStr(s string) int {

	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0 //当前这个i就是在检查的那个数
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI >= start {
			start = lastOccurred[ch]
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i + 1 //这里加1是因为，如果有一个数之前没有出现，那么
	}

	return maxLength
}

func main() {

	fmt.Println(lenthOfNoneRepeatingSubStr("abcabcbb"))
	fmt.Println(lenthOfNoneRepeatingSubStr("我爱慕课网"))
	fmt.Println(lenthOfNoneRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
