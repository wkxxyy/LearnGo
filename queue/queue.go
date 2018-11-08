package queue

// An FIFO queue
type Queue []interface{} //这里就是等价代换，相当于Queue是个[]intt？
//interface{}就是指任何类型

//Pushes the element into the queue
//			e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) //这里限制进去的只能是int
}

//Pops element from head
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) //把head里面的值强行转换成Int
}

//Return if the queue is empty or not

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
