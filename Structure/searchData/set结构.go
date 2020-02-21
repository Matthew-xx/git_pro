package main

import "fmt"

//增删改效率较低，只适合查询，判断是否存在
type Set struct {
	buf []interface{}  //存储数据
	num int
	hash map[interface{}] bool  //借助map实现映射
}

//新建一个可变长的set
func NewSet() *Set {
	return &Set{make([]interface{},0),0,make(map[interface{}] bool)}
}

func (this *Set) Add(value interface{}) bool{
	if this.IsExit(value) {
		return false
	}else {
		this.buf = append(this.buf,value)
		this.hash[value] = true
		this.num++
		return true
	}
}

func (this *Set) IsExit(value interface{}) bool{
	return this.hash[value]
}

func (this *Set) Strings() []interface{} {
	return this.buf
}

func main()  {
	set := NewSet()
	set.Add(1)
	set.Add(2)
	fmt.Println(set)
}
