package main

import (
	"../frontend/controller"
	"net/http"
)

func main()  {
	http.Handle("/",http.FileServer(http.Dir("crawl_double/fontend/view")))
	//只要不是search就展示文件,这样就可以访问view下面的静态文件

	http.Handle("/search",controller.CreateSearchResultHandler("crawl_double/frontend/view/template.html"))
	err := http.ListenAndServe(":8888",nil)
	if err!= nil {
		panic(err)
	}
}
