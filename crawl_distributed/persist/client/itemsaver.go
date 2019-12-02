package client

import (
	"../../../crawl_double/engine"
	"log"
	"../../rpcsupport"
	"../../config"
)

func ItemSaver(host string) (chan engine.Item ,error){
	client,err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("item save:got item #%d: %v",itemCount,item)
			itemCount++

			result := ""
			//每收到一个item会call rpc to save
			client.Call(config.ItemSaverRpc,item,&result)
			if err != nil {
				log.Printf("item saver:error"+"saveing item %v:%v",item,err)
			}
		}

	}()
	return out,nil    //做成的话return out 和nil
}

