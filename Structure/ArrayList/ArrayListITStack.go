package ArrayList

type StackArrayX interface {
	Clear()
	Size()  int
	Pop()   interface{}  //弹出
	Push(data interface{})  //压入
	IsFull()  bool  //是否满
	IsEmpty()  bool  //是否为空
}

type StackX struct {
	myarray *ArrayList
	Myit  Iterator
}

func NewArraryListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()  //数组
	mystack.Myit = mystack.myarray.Iterator()  //迭代

	return mystack
}

func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
	mystack.myarray.TheSize = 0
}

func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *StackX) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize-1)
		return last
	}
	return nil
} //弹出

func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}

} //压入

func (mystack *StackX) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	}else {
		return false
	}
} //是否满

func (mystack *StackX) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	}else {
		return false
	}
} //是否为空

