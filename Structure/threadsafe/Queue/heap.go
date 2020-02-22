package Queue

import "sync"

type Item interface {
	Less (than Item) bool
}

type Int int

func (x Int) Less(than Item) bool {
	return x <than.(Int)
}

//实现一个最小堆
type Heap struct {
	locK *sync.Mutex  //线程安全
	data  []Item  //数组
	min bool
}
//标准堆
func NewHeap() *Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),true}

}
//最大堆
func NewMin() *Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),true}
}
//最小堆
func NewMax() *Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),false}
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}
//求长度
func (h *Heap) Len() int {
	return len(h.data)
}
//抓取数据
func (h *Heap) Get(index int) Item {
	return h.data[index]
}

//插入数据
func (h *Heap) Insert(It Item)  {
	h.locK.Lock()
	defer h.locK.Unlock()
	h.data= append(h.data,It)
	h.SiftUp()  //每次插入后up
	return
}
//根据类型返回比大小
func (h *Heap) Less(a,b Item) bool {
	if h.min {
		return a.Less(b)
	}else {
		return b.Less(a)
	}
}
//压缩(弹出一个
func (h *Heap) Extract() Item {
	h.locK.Lock()
	defer h.locK.Unlock()
	if h.Len() == 0 {
		return nil
	}
	e1 := h.data[0]  //第一个数据
	last := h.data[h.Len()-1]  //最后一个数据
	if h.Len() == 1 {
		h.data = nil  //弹出一个数据,重新分配内存
		return nil
	}
	h.data = append([]Item{last},h.data[1:h.Len()-1]...)
	h.SiftDown()  //压缩的时候down
	return e1
}
//弹出一个极大值
func (h *Heap) SiftUp()  {
	//堆排序循环过程,对n,2n+1取值比较
	for i,parent:=h.Len()-1,h.Len()-1;i>0;i=parent {
		parent = i/2
		if h.Less(h.Get(i),h.Get(parent)) { //处理数据
			h.data[parent],h.data[i]=h.data[i],h.data[parent]
		}else {
			break
		}
	}
}

//弹出一个极小值
func (h *Heap) SiftDown()  {
	//堆排序循环过程,对n,2n+1取值比较
	for i,child:=0,1;i<h.Len() && i*2+1 <h.Len();i=child {
		child = i*2+1
		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1),h.Get(child)){
			child++  //循环左右节点的过程
		}
		if h.Less(h.Get(i),h.Get(child)) {
			break  //i小于child
		}
		h.data[i],h.data[child] = h.data[child],h.data[i]
	}
}

