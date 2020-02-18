package CricleQueue

import "github.com/pkg/errors"

//循环队列，避免重复计算浪费资源
const QueueSize = 100 //最多存储（QueueSize-1）个数据，空一个空格表示满格
type CricleQueue struct {
	date [QueueSize]interface{} //存储数据的结构
	front int  //头部位置
	rear  int  //尾部位置
}

//初始化
func InitQueue(q *CricleQueue)  { //头部尾部重合，为空
	q.front = 0
	q.rear = 0
}

func Queuelength(q *CricleQueue) int {
	return (q.rear-q.front+QueueSize) % QueueSize
}

func EnQueue(q *CricleQueue,data interface{}) (err error) {
	if (q.rear + 1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已经满了")
	}
	q.date[q.rear] = data  //入队
	q.rear = (q.rear + 1) %QueueSize  //归位

	return nil
}

func DnQueue(q *CricleQueue) (data interface{},err error) {
	if q.rear == q.front {
		return nil,errors.New("队列为空")
	}
	res := q.date[q.front] //取出第一个数据
	q.date[q.front] = 0  //清空数据
	q.front = (q.front + 1) % QueueSize  //大于100时取模回到首位
	return res,nil
}

