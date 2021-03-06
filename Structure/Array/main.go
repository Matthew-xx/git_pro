package main

import (
	"./ArrayList"
	"./CricleQueue"
	"./Queue"
	"./StackArray"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)


func main1()  {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append("a")
	fmt.Println(list)
}


func main2()  {
	//定义接口对象，赋值的对象须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()  //在List接口未实现接口里面定义的全部方法时，无法实现该定义
	list.Append(1)
	list.Append(2)
	list.Append("a")
	fmt.Println(list)
}

func main3()  {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("a")
	list.Insert(1,"b")
	for i:=0;i<15;i++ {
		list.Insert(1,"c")
	}
	fmt.Println(list)
	list.Delete(2)
	fmt.Println(list)
}

func main4()  {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("a")
	list.Insert(1,"b")
	list.Insert(3,"c")
	fmt.Println(list)
	for it:=list.Iterator();it.HasNext(); {
		item,_ := it.Next("1234")
		if item == "a"{
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}

func main5()  {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main6()  {
	mystack := ArrayList.NewArraryListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main7()  {
	mystack := ArrayList.NewArraryListStackX()  //ArrayList.NewArraryListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())

	for it := mystack.Myit;it.HasNext(); {
		item,_ := it.Next("1234")
		fmt.Println(item)
	}
}

func main8()  {
	mystack := StackArray.NewStack()
	mystack.Push(4)
	last := 0  //保存结果
	for !mystack.IsEmpty() {
		data := mystack.Pop()

		if data == 0{
			last += 0
		}else {
			last += data.(int)  //递归取出
			mystack.Push((data.(int)-1))  //递归存数
		}
	}
	fmt.Println(last)
}

//递归读取文件夹（缓存交给了系统
func GetAll(path string,files []string) ([]string,error) {
	read,err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files,errors.New("文件夹不可读取")
	}

	for _,fi := range read{ //循环每个
		if fi.IsDir() {  //判断是否是文件夹
			fulldir := path+"\\"+fi.Name()  //构造新的路径
			files = append(files,fulldir)  //追加路径
			files,_ = GetAll(fulldir,files)  //文件夹递归处理

		}else { //如果是文件
			fulldir := path+"\\"+fi.Name()  //构造新的路径
			files = append(files,fulldir)  //追加路径
		}
	}
	return files,nil
}

//递归读取文件
func main9()  {
	path := "F:\\BaiduNetdiskDownload\\密码学\\day03"  //  "\\"转义字符
	files := []string{}  //数组字符串
	files,_ = GetAll(path,files) //抓取所有文件
	for i:=0;i<len(files);i++ {  //打印路径
		fmt.Println(files[i])
	}
}

//通过栈递归读取文件
func mainx()  {
	path := "F:\\BaiduNetdiskDownload\\密码学\\day03"  //  "\\"转义字符
	files := []string{}  //数组字符串

	mystack := StackArray.NewStack()
	mystack.Push(path)

	for !mystack.IsEmpty() {
		path := mystack.Pop().(string) //取出数据,实例化成字符串
		files = append(files,path)  //加入列表
		read,_ := ioutil.ReadDir(path) //读取

		for _,fi := range read{
			if fi.IsDir() {
				fulldir := path+"\\"+fi.Name()  //构造新的路径
				files = append(files,fulldir)  //追加路径
				mystack.Push(fulldir)  //如果是文件夹继续压栈循环
			}else {
				fulldir := path+"\\"+fi.Name()  //构造新的路径
				files = append(files,fulldir)  //追加路径
			}
		}
	}
	for i:=0;i<len(files);i++ {
		fmt.Println(files[i])
	}
}

//(实现层级,树状）
func GetAllX(path string,files []string,level int) ([]string,error) {
	levelstr := ""
	if level ==1 {
		levelstr = "+"
	}else {
		for ;level>1;level-- {
			levelstr+= "|--"
		}
		levelstr+="+"
	}

	read,err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files,errors.New("文件夹不可读取")
	}

	for _,fi := range read{ //循环每个
		if fi.IsDir() {  //判断是否是文件夹
			fulldir := path+"\\"+fi.Name()  //构造新的路径
			files = append(files,levelstr+fulldir)  //追加路径
			files,_ = GetAllX(fulldir,files,level+1)  //文件夹递归处理,层级+1

		}else { //如果是文件
			fulldir := path+"\\"+fi.Name()  //构造新的路径
			files = append(files,levelstr+fulldir)  //追加路径
		}
	}
	return files,nil
}

func mainx1()  {
	path := "F:\\BaiduNetdiskDownload\\密码学\\day03"  //  "\\"转义字符
	files := []string{}  //数组字符串
	files,_ = GetAllX(path,files,1) //抓取所有文件
	for i:=0;i<len(files);i++ {  //打印路径
		fmt.Println(files[i])
	}
}

func mainl()  {
	myq := Queue.NewQueue()
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(5)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
}

func mainm()  {
	path := "F:\\BaiduNetdiskDownload\\密码学\\day03"  //  "\\"转义字符
	files := []string{}  //数组字符串
	mystack := Queue.NewQueue()
	mystack.EnQueue(path)

	for ;; {  //死循环,不断从队列中取出数据
		path := mystack.DeQueue()
		if path == nil{
			break
		}
		fmt.Println("get",path)
		read,_ := ioutil.ReadDir(path.(string))
		for _,fi := range read{
			if fi.IsDir() {
				fulldir := path.(string)+"\\"+fi.Name()
				fmt.Println("Dir",fulldir)
				mystack.EnQueue(fulldir)
			}else {
				fulldir := path.(string)+"\\"+fi.Name()
				files = append(files,fulldir)
				fmt.Println("file",fulldir)
			}
		}
	}
	for i:=0;i<len(files);i++ {
		fmt.Println(files[i])
	}
}

func main()  {
	var myq CricleQueue.CricleQueue
	CricleQueue.InitQueue(&myq)  //初始化
	CricleQueue.EnQueue(&myq,1)
	CricleQueue.EnQueue(&myq,2)
	CricleQueue.EnQueue(&myq,3)
	fmt.Println(CricleQueue.DnQueue(&myq))
	fmt.Println(CricleQueue.DnQueue(&myq))
}

