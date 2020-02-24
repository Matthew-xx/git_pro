package main

import (
	"fmt"
	"github.com/pkg/errors"

	"net/http"
	"net/rpc"
)

type Args struct {
	A,B int  //两个数据
}

type Query struct {
	X,Y int
}

type Last int

func (t *Last) Multiply(args *Args,reply *int) error{
	*reply = args.A * args.B
	fmt.Println("乘法结果",reply)
	return nil
}

func (t *Last) Divide(args *Args,query *Query) error{
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	query.X = args.A / args.B
	query.Y = args.A % args.B
	fmt.Println("除法结果",query)

	return nil
}

func main()  {
	la := new(Last)
	fmt.Println(la,"la")
	rpc.Register(la) //注册类型
	rpc.HandleHTTP()  //设定http类型
	err := http.ListenAndServe("127.0.0.1:1234",nil)  //监听及服务
	//listen,err := net.Listen("tcp","127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	//http.Serve(listen,nil)
}
