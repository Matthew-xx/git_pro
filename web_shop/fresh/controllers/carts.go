package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"../models"
	"strconv"
)

type CartController struct {
	beego.Controller
}

//获取购物车数量的函数
func GetCartCount(this*beego.Controller)int{
	//从redis中获取购物车数量
	userName :=this.GetSession("userName")
	if userName == nil{
		return 0
	}
	o := orm.NewOrm()
	var user models.User
	user.Name = userName.(string)
	o.Read(&user,"Name")

	conn,err :=redis.Dial("tcp","192.168.110.81:6379")
	if err !=nil{
		return 0
	}
	defer conn.Close()

	rep,err := conn.Do("hlen","cart_"+strconv.Itoa(user.Id))
	cartCount,_:=redis.Int(rep,err)

	return cartCount


	//cart_userId
}

