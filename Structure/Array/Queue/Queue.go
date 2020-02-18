package Queue

//队列，先进先出。是广度遍历，先广度再深度
//栈、递归都是深度遍历
type MyQueue interface {
	Size() int
	Front() interface{}  //第一个元素
	End()  interface{}  //最后一个
	IsEmpty() bool
	EnQueue(data interface{})  //入队
	DeQueue() interface{} //出队
	Clear()  //清空
}

type Queue struct {
	dataStore []interface{}  //队列的数据存储
	theSize   int
}

func NewQueue() *Queue {
	myqueue := new(Queue) //初始化开辟结构体
	myqueue.dataStore = make([]interface{},0)
	myqueue.theSize = 0

	return myqueue
}

func (myq *Queue) Size() int{
	return myq.theSize
}

func (myq *Queue) Front() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[0]
} //第一个元素

func (myq *Queue) End()  interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[myq.Size()-1]
} //最后一个

func (myq *Queue) IsEmpty() bool {
	return myq.Size() == 0
}

func (myq *Queue) EnQueue(data interface{}) {
	myq.dataStore = append(myq.dataStore,data)
	myq.theSize++
} //入队

func (myq *Queue) DeQueue() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	data := myq.dataStore[0]  //获取数据
	if myq.Size()>1 {
		myq.dataStore = myq.dataStore[1:myq.Size()]  //截取队列
	}else {
		myq.dataStore = make([]interface{},0)
	}
	myq.theSize--
	return data  //返回数据
}//出队

func (myq *Queue) Clear() {
	myq.dataStore = make([]interface{},0)
	myq.theSize = 0
}

