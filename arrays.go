package tree

//go语言传的是值，不是引用
import "fmt"

func printArray(arr []int) {
	arr[0] = 100 //直接不用取地址，直接改
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10} //自己数数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")

	printArray(arr1[:])

	fmt.Println("printArray(arr3)")

	printArray(arr3[:])
	//printArray(arr2) [3]int和[5]int是两个东西
	fmt.Println("printArray(arr1,arr3)")
	fmt.Println(arr1, arr3)

}
