package queue

type Node struct {
	value int
}

type Queue struct {
	queue []Node
}

func (q *Queue) Enqueue(value int) {
		q.queue = append(q.queue , Node{value})
}
func (q *Queue) Dequeue() (int,bool) {
	if q.Size() == 0 {
		return 0, false
	}
	node := q.queue[0]
	q.queue = append(q.queue[1:],q.queue[len(q.queue):]...)
	return node.value,true
}
func (q Queue) Size() int {
	return len(q.queue)
}
