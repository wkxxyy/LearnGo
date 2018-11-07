package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(post Poster) {
	post.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc . com",
	})

	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) { //可以这么查看接口类型
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents：", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent：", v.UserAgent)

	}
	fmt.Println()
}

func main() {

	var r Retriever

	retriever := mock.Retriever{"this is a fake imooc.com"}

	r = &retriever

	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	//可以这么取得他肚子里的一个类型，这个r必须是和括号里面的是同；类型的,如果r是有*号的，那么括号里面也要加*
	//也可以这么查看接口类型
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")

	fmt.Println(session(&retriever))
}
