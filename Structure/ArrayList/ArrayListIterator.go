package ArrayList

import "github.com/pkg/errors"

//迭代器

type Iterator interface {
	HasNext()  bool   //是否有下一个
	Next(password string) (interface{},error)  //下一个
	Remove() //删除
	GetIndex() int  //得到索引
}

type Iterable interface {
	Iterator() Iterator  //构造时初始化接口(迭代
}

//构造指针访问数组
type ArrayListIterator struct {
	list *ArrayList   //数组指针
	currentindex int  //当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)  //构造迭代器
	it.currentindex = 0
	it.list = list

	return it
}

func (it *ArrayListIterator) HasNext()  bool {
	return it.currentindex < it.list.TheSize   //是否有下一个
}

func (it *ArrayListIterator) Next(password string) (interface{},error) {
	if password == "1234"{
		if !it.HasNext() {
			return nil,errors.New("没有下一个")
		}
		value,err := it.list.Get(it.currentindex)  //抓取当前数据
		it.currentindex++
		return value,err
	}else {
		return nil,nil
	}


} //下一个

func (it *ArrayListIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex)  //删除一个元素
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentindex   //得到索引
}

