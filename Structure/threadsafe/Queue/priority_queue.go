package Queue

//优先队列
type PriItem struct {
	value interface{}  //值
	priority int   //优先级
}
//优先队列,基于堆实现
type PriQueue struct {
	data *Heap
}

//队列中间的元素
func NewPriItem(value interface{},priority int) *PriItem {
	return &PriItem{value,priority}
}

func (x PriItem) Less(than Item) bool {
	return x.priority<than.(PriItem).priority //对比优先级大小
}


func NewMaxPriQueue() *PriQueue {
	return &PriQueue{NewMax()}
}
func NewMinPriQueue() *PriQueue {
	return &PriQueue{NewMin()}
}

func (pq *PriQueue) Len() int {
	return pq.data.Len()  //队列的长度
}

func (pq *PriQueue) Insert(e1 PriItem) {
	pq.data.Insert(e1)
}
//弹出
func (pq *PriQueue) Extract() PriItem{
	return pq.data.Extract().(PriItem)
}
//修改优先级
func (pq *PriQueue) ChangePri(val interface{},priority int){
	var storage = NewQueue()  //队列备份数据
	poped := pq.Extract()  //拿出最小数值
	for val != poped.value {  //poped实例化取其值
		if pq.Len() == 0 {
			return
		}
		storage.Push(poped)  //压入数据
		poped = pq.Extract()
	}
	poped.priority = priority  //修改优先级
	pq.data.Insert(poped)  //插入数据

	for storage.Len() > 0 {
		pq.data.Insert(storage.Shift().(Item))  //其余数据重新放入优先队列
	}
}


