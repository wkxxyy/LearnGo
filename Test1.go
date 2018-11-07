package main

import "fmt"

func main() {

	var nums []int = []int{2, 7, 11, 15}

	var target int = 9

	var result []int = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}

	}

	fmt.Println(result)

}
