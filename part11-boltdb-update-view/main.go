package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	//创建打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 更新表数据
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b != nil {
			err := b.Put([]byte("ll"), []byte("Send 100 BTC To feifei..."))
			if err != nil {
				log.Panic("数据库更新失败")
			}
		}
		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
