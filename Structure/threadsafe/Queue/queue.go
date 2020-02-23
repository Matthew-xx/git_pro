package Queue

import (
	"sync"
)

type Queue struct {
	queue []interface{}
	len int
	lock *sync.Mutex
}

//新建队列
func NewQueue() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{},0)
	queue.len = 0
	queue.lock = new(sync.Mutex) //初始化

	return queue
}

//解决线程安全
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) Shift() (e1 interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	e1,q.queue = q.queue[0],q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(e1 interface{})  {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue,e1)
	q.len++

	return
}

func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue[0]
}
