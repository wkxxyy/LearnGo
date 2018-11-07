package queue

type Queue []interface{} //这里就是等价代换，相当于Queue是个[]intt？
//interface{}就是指任何类型

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) //这里限制进去的只能是int
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) //把head里面的值强行转换成Int
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
