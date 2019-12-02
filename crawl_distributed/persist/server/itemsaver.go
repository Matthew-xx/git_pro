package main

import (
	"../../persist"
	"../../rpcsupport"
	"flag"
	"log"
	"github.com/olivere/elastic"
	"../../config"
	"fmt"
)

var port = flag.Int("port",0,"the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}

	log.Fatal(serveRpc(fmt.Sprintf(": %d",*port) ,config.ElasticIndex))

}

//为了测试方便，另外包装
func serveRpc(host,index string) error {
	client,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return  rpcsupport.ServeRpc(host,&persist.ItemServerService{
		Client: client,
		Index:  index,
	})
}
