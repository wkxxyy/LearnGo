package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string { //有点像是toString
	return fmt.Sprintf(
		"Retriever:{Contents=%s}", r.Contents)
}

func (r *Retriever) Get(url string) string { //实现了接口的方法
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
