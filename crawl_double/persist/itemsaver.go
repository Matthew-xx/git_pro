package persist

import (
	"../engine"
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
)

//存储功能

func ItemSaver(index string) (chan engine.Item ,error){
	client,err := elastic.NewClient( elastic.SetSniff(false)) //因为是在docker内网，外网无法sniff，所以直接false
	if err != nil {
		//panic(err)
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("item save:got item #%d: %v",itemCount,item)
			itemCount++

			err := Save(client,index,item)
			if err != nil {
				log.Printf("item saver:error"+"saveing item %v:%v",item,err)
			}
		}

	}()
	return out,nil    //做成的话return out 和nil
}

func Save(client *elastic.Client,index string,item engine.Item) error{


	if item.Type == "" {
		return  errors.New("must supply Type")  //type为空
	}

	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)       //存数据  ,（表和id）type和ID可由parser给到，而index可以是系统，也即全部存到里面（数据库）
	//第一个index可创建可修改，后面三元组表文档的位置（index，type，ID（可不写由系统分配））
	if item.Id != "" {
		indexService.Id(item.Id)    //id不为空时才传进去
	}
	_,err := indexService.Do(context.Background())   //

	if err != nil {
		//panic(err)
		return err
	}
	//fmt.Printf("%+v",resp) // %+v 打印结构体的时候可以把字段名也打印出来
	return nil
}




