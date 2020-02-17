package ArrayList


type StackArray interface {
	Clear()
	Size()  int
	Pop()   interface{}  //弹出
	Push(data interface{})  //压入
	IsFull()  bool  //是否满
	IsEmpty()  bool  //是否为空
}

type Stack struct {
	myarray *ArrayList
	capsize int  //最大范围
}

func NewArraryListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList()
	mystack.capsize = 10

	return mystack
}

func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
	mystack.capsize = 10
}

func (mystack *Stack) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize-1)
		return last
	}
	return nil
} //弹出

func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}

} //压入

func (mystack *Stack) IsFull() bool {
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	}else {
		return false
	}
} //是否满

func (mystack *Stack) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	}else {
		return false
	}
} //是否为空

