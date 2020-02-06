package main

import (
	_ "../fresh/models"
	_ "../fresh/routers"
	"github.com/astaxie/beego"
)

//func init()  {
//	models.RegisterDB() //初始化数据库
//}

func main() {
	//
	//orm.Debug = true  //方便调试看数据库是否创建
	//orm.RunSyncdb("default", false, true)

	beego.Run()
}

