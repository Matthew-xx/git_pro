package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../../services"
	"../../datasource"
	"../../models"
	"log"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.ProductService
}

// http://localhost:8080/
func (c *IndexController) Get() mvc.Result {
	//c.Service.GetAll()
	//return mvc.Response{
	//	Text:"ok\n",
	//}
	datalist := c.Service.GetAll()
	//var datalist []models.StarInfo
	//set the model and render the view template.
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "商铺",
			"Datalist": datalist,
		},
	}
}

// http://localhost:8080/{id}
func (c *IndexController) GetBy() mvc.Result {
	id,err := c.Ctx.URLParamInt64("id")
	if err != nil {
		return nil
	}
	P_name := c.Ctx.URLParam("productname")
	datalist := c.Service.Get(id,P_name)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "商品信息",
			"info":  datalist,
		},
	}
}

// http://localhost:8080/search?country=巴西
func (c *IndexController) GetSearch() mvc.Result {
	P_name := c.Ctx.URLParam("productname")
	datalist := c.Service.Search(P_name)
	// set the model and render the view template.
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "商品",
			"Datalist": datalist,
		},
	}
}

// 集群多服务器的时候，才用得上这个接口
// 性能优化的时候才考虑，加上本机的SQL缓存
// http://localhost:8080/clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.Product{})
	if err != nil {
		log.Fatal(err)
	}
	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}

