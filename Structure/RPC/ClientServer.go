package main

import (
	"fmt"
	"net/rpc"
)

type Args1 struct {
	A,B int  //两个数据
}

type Query1 struct {
	X,Y int
}

func main()  {
	serverip := "127.0.0.1:1234"
	client,err := rpc.DialHTTP("tcp",serverip)
	if err != nil {
		panic(err)
	}
	i1 := 12
	i2 := 5
	args := Args1{i1,i2}

	var reply int
	err = client.Call("Last.Multiply",args,&reply)  //大写
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d * %d = %d\n",args.A ,args.B,reply)

	var query Query1
	err = client.Call("Last.Divide",args,&query)  //大写
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d / %d = %d mod %d",args.A ,args.B,query.Y,query.Y)
}

