package main

import "fmt"

func main() {

	m := map[string]string{ //这是一个hashmap所以无序
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3==nill

	//var m3 map[string]int
	//m3=make(map[string]int)
	//m3["good"]=1

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	//var k string
	//var v string
	for k, i := range m {
		fmt.Println(k, i)
	}

	fmt.Println("Getting values")
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName, ok)
	}

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key does not exist")
	}
	//如果k错误，那就拿到空，就是类型的一个默认值

	fmt.Println("Delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(len(m))

	//map使用哈希表，key必须可以比较相等
	//除了slice,map,function的内建类型都可以作为key
	//struct类型不包含上述字段，也可以作为key
}
