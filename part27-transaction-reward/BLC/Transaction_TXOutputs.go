package BLC

import (
	"encoding/json"
	"log"
)

type TXOutputs struct {
	UTXOS []*UTXO
}

// Serialize 将区块序列化成字节数组
func (txOutputs *TXOutputs) Serialize() []byte {
	//var result bytes.Buffer
	//encoder := gob.NewEncoder(&result)
	//err := encoder.Encode(txOutputs)
	//if err != nil {
	//	log.Panic(err)
	//}
	//return result.Bytes()
	txOutputsBytes, err := json.Marshal(txOutputs)
	if err != nil {
		log.Panicln(err)
	}
	return txOutputsBytes
}

// DeserializeTXOutputs 反序列化
func DeserializeTXOutputs(txOutputsBytes []byte) *TXOutputs {
	var txOutputs TXOutputs
	//decoder := gob.NewDecoder(bytes.NewReader(txOutputsBytes))
	//err := decoder.Decode(&txOutputs)
	//if err != nil {
	//	log.Panic(err)
	//}
	//return &txOutputs
	if len(txOutputsBytes) == 0 {
		return nil
	}
	err := json.Unmarshal(txOutputsBytes, &txOutputs)
	if err != nil {
		log.Panicln(err)
	}
	return &txOutputs
}
