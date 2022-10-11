package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	//创建打开数据库
	db, err := bolt.Open("D:\\workspace\\go\\blockchain\\part12-boltdb-update-view/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 更新表数据
	err = db.View(func(tx *bolt.Tx) error {
		//获取表对象
		b := tx.Bucket([]byte("BlockBucket"))

		//查看表数据
		if b != nil {
			data := b.Get([]byte("ll"))
			fmt.Printf("data:%s\n", data)
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
