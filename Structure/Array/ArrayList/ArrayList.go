package ArrayList

import (
	"fmt"
	"github.com/pkg/errors"
)

//接口
type List interface {
	Size()   int     //数组大小
	Get(index int) (interface{},error)  //抓取第几个元素,返回的结果
	Set(index int,newval interface{}) error //修改第几个数据，新数值,返回错误
	Insert(index int,newval interface{}) error  //插入
	Append(newval interface{})  //追加数据
	Clear() //清空
	Delete(index int) error //删除
	String() string  //返回字符串
	Iterator() Iterator  //迭代器
}

//数据结构
type ArrayList struct {
	dataStore []  interface{}  //代表datastore存储的数组结构是一个泛型（整型，字符串..
	TheSize   int     //数组的大小
}

func NewArrayList() *ArrayList  {
	list := new(ArrayList)  //初始化结构体
	list.dataStore = make([]interface{},0,10)  //对数组开辟内存开辟空间10个
	list.TheSize = 0
	return list
}

//返回数据大小
func (list *ArrayList) Size() int  {
	return list.TheSize
}

//抓取数据
func (list *ArrayList) Get(index int) (interface{},error) {
	if index <0 || index >= list.TheSize {
		return nil,errors.New("索引越界")
	}
	return list.dataStore[index],nil
}

//追加数据
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore,newval)
	list.TheSize++
}

//返回数组字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int,newval interface{}) error {
	if index <0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval //设置

	return nil
}

func (list *ArrayList) Insert(index int,newval interface{}) error {
	if index <0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.checkisFull()  //检测内存，如果满了则自动追加
	list.dataStore = list.dataStore[:list.TheSize+1]  //插入数据时需将内存追加(移动)一位
	for i:= list.TheSize;i>index;i-- {  //从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval //插入数据
	list.TheSize++  //索引追加

	return nil
}
//重新设置空间
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{},0,10)  //对数组开辟内存开辟空间10个
	list.TheSize = 0
}

//删除数据
func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index],list.dataStore[index+1:]...)  //通过叠加删除index位置的数据
	list.TheSize--

	return nil
}

//判断空间是否已满
func (list *ArrayList) checkisFull()  {
	if list.TheSize == cap(list.dataStore) {  //判断内存空间的使用
		newdataStore := make([]interface{},2*list.TheSize,2*list.TheSize)  //开辟双倍内存
		//make中间的参数为0代表没开辟新内存
		//copy(newdataStore,list.dataStore)  //拷贝

		for i:=0;i<len(list.dataStore);i++ {
			newdataStore[i]=list.dataStore[i]
		}
		list.dataStore = newdataStore
	}
}

