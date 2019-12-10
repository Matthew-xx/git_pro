package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main()  {
	//创建或打开数据库
	db, err := bolt.Open("my.db",0600,nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//查看数据
	err = db.View(func(tx *bolt.Tx) error {
		//获取表对象
		b := tx.Bucket([]byte("BlockBucket"))

		//读取数据
		if b != nil {
			data := b.Get([]byte("l"))
			fmt.Printf("%s\n",data)
		}

		return nil  //返回nil,以便数据库操作
	})
	
	if err != nil {
		log.Panic(err)
	}

}
