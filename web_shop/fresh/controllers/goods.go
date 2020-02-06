package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"../models"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type GoodsController struct {
	beego.Controller
}

//使用beego控制器，父类，子类的所有控制器都能调用方法
func GetUser(this *beego.Controller) string{
	userName := this.GetSession("userName")
	if userName == nil{
		this.Data["userName"] = ""
	}else{
		this.Data["userName"] = userName.(string)  //显示用户名
		return userName.(string)
	}
	return ""
}


func PageTool(pageCount int,pageIndex int)[]int{

	var pages []int
	if pageCount <= 5{
		pages = make([]int,pageCount)
		for i,_ := range pages{
			pages[i] = i + 1
		}

		//pages = [1,2,..,pageCount]
	}else if pageIndex <= 3{
		//pages := make([]int,5)
		pages = []int{1,2,3,4,5}
	}else if pageIndex > pageCount - 3 {
		//pages = [6, 7, 8, 9, 10]
		pages = []int{pageCount -4,pageCount - 3,pageCount - 2,pageCount -1 ,pageCount}
	}else {
		pages = []int{pageIndex - 2,pageIndex -1 ,pageIndex,pageIndex + 1, pageIndex + 2}
	}
	return pages

}


//展示首页
func(this *GoodsController) ShowIndex(){
	GetUser(&this.Controller)
	o := orm.NewOrm()
	//获取类型数据
	var goodsTypes []models.GoodsType
	o.QueryTable("GoodsType").All(&goodsTypes)
	this.Data["goodsTypes"] = goodsTypes

	//获取轮播图数据
	var indexGoodsBanner []models.IndexGoodsBanner
	o.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&indexGoodsBanner)
	this.Data["indexGoodsBanner"] = indexGoodsBanner

	//获取促销商品数据
	var promotionGoods []models.IndexPromotionBanner
	o.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&promotionGoods)
	this.Data["promotionsGoods"] = promotionGoods

	//首页展示商品数据
	goods := make([]map[string]interface{},len(goodsTypes))

	//向切片interface中插入类型数据
	for index, value := range goodsTypes{
		//获取对应类型的首页展示商品
		temp := make(map[string]interface{})
		temp["type"] = value
		goods[index] = temp
	}
	//商品数据


	for _,value := range goods{
		var textGoods []models.IndexTypeGoodsBanner
		var imgGoods []models.IndexTypeGoodsBanner
		//获取文字商品数据
		o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType","GoodsSKU").OrderBy("Index").Filter("GoodsType",value["type"]).Filter("DisplayType",0).All(&textGoods)
		//获取图片商品数据
		o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType","GoodsSKU").OrderBy("Index").Filter("GoodsType",value["type"]).Filter("DisplayType",1).All(&imgGoods)

		value["textGoods"] = textGoods
		value["imgGoods"] = imgGoods
	}
	this.Data["goods"] = goods



	this.TplName = "index.html"
}

func ShowLaout(this*beego.Controller){
	//查询类型
	o := orm.NewOrm()
	var types []models.GoodsType
	o.QueryTable("GoodsType").All(&types)
	this.Data["types"] = types
	//获取用户信息
	GetUser(this)
	//指定layout
	this.Layout = "goodsLayout.html"
}

//展示商品详情
func(this*GoodsController)ShowGoodsDetail(){
	//获取数据
	id,err := this.GetInt("id")
	//校验数据
	if err != nil{
		beego.Error("浏览器请求错误")
		this.Redirect("/",302)
		return
	}
	//处理数据
	o := orm.NewOrm()
	var goodsSku models.GoodsSKU
	goodsSku.Id = id
	//o.Read(&goodsSku)
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType","Goods").Filter("Id",id).One(&goodsSku)

	//获取同类型时间考前的两条商品数据
	var goodsNew []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType",goodsSku.GoodsType).OrderBy("Time").Limit(2,0).All(&goodsNew)
	this.Data["goodsNew"] = goodsNew

	//返回视图
	this.Data["goodsSku"] = goodsSku

	//添加历史浏览记录
	//判断用户是否登录
	userName := this.GetSession("userName")
	if userName != nil{
		//查询用户信息
		o := orm.NewOrm()
		var user models.User
		user.Name = userName.(string)
		o.Read(&user,"Name")
		//添加历史记录,用redis存储
		conn,err := redis.Dial("tcp","192.168.110.81:6379")
		defer conn.Close()
		if err != nil{
			beego.Info("redis链接错误")
		}
		//把以前相同商品的历史浏览记录删除
		conn.Do("lrem","history_"+strconv.Itoa(user.Id),0,id)
		//添加新的商品浏览记录
		conn.Do("lpush","history_"+strconv.Itoa(user.Id),id)


	}



	ShowLaout(&this.Controller)
	cartCount := GetCartCount(&this.Controller)
	this.Data["cartCount"] = cartCount
	this.TplName = "detail.html"

}


