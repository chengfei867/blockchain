package BLC

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

const utxoTableName = "utxoTableName"

type UTXOSet struct {
	Blockchain *Blockchain
}

// ResetUTXOSet 重置数据库表
func (utxoSet *UTXOSet) ResetUTXOSet() {
	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		if b != nil {
			err := tx.DeleteBucket([]byte(utxoTableName))
			if err != nil {
				log.Panic(err)
			}
		}
		b, _ = tx.CreateBucket([]byte(utxoTableName))
		if b != nil {
			//[string]*TXOutputs
			txOutputsMap := utxoSet.Blockchain.FindUTXOMap()
			for keyHash, outs := range txOutputsMap {
				txHash, _ := hex.DecodeString(keyHash)
				err := b.Put(txHash, outs.Serialize())
				if err != nil {
					log.Panicln(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (utxoSet *UTXOSet) findUTXOForAddress(address string) []*UTXO {
	var utxos []*UTXO
	err := utxoSet.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		// 游标
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			txOutputs := DeserializeTXOutputs(v)
			for _, utxo := range txOutputs.UTXOS {
				if utxo.Output.UnLockScriptPubKeyWithAddress(address) {
					utxos = append(utxos, utxo)
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return utxos
}

func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	UTXOS := utxoSet.findUTXOForAddress(address)
	var amount int64
	for _, utxo := range UTXOS {
		amount += utxo.Output.Value
	}
	return amount
}
