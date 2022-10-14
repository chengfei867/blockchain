package main

import (
	"blockchain/publicChain/part13-block-bolt/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	mydb, err := bolt.Open("E:\\WorkSpace\\goWorkspace\\Go\\src\\blockchain\\publicChain\\part13-block-bolt\\my.db", 0600, nil)
	if err != nil {
		log.Panicln(err)
	}
	defer func(mydb *bolt.DB) {
		err := mydb.Close()
		if err != nil {

		}
	}(mydb)
	//createDB(mydb)
	blockBytes := readDB(mydb)
	block := new(BLC.Block)
	block = BLC.Deserialize(blockBytes)
	fmt.Printf("hash:%x", block.Hash)
}

func createDB(mydb *bolt.DB) {
	block := BLC.NewBlock("TestDB", 1, []byte{00000000000000000000000000000000})
	blockBytes, err := block.Serialize()
	if err != nil {
		log.Panicln(err)
	}

	err = mydb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("BLC"))
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put([]byte("LastHash"), block.Hash)
		err = bucket.Put(block.Hash, blockBytes)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func readDB(mydb *bolt.DB) []byte {
	blockBytes := *new([]byte)
	err := mydb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("BLC"))
		lastHash := bucket.Get([]byte("LastHash"))
		blockBytes = bucket.Get(lastHash)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return blockBytes
}
