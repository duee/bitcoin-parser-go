package service

import (
	"fmt"
	"testing"
)

func Test_loadUtxo(t *testing.T) {
	LoadUTXO()
}

func Test_getUtxo(t *testing.T) {
	utxos,_:=GetUtxo("myqvLpN5vmcSuwEMV4cfRHAL88Aq5MmofQ")
	fmt.Println("utxos = ", utxos)
}

func Test_getKey(t *testing.T) {
	Getkey("2NBMEXgxXbm9zhXxNbmsumUx2m3vrFdhiyx_71712")
}

/**
用于测试leveldb的前缀查询
*/
func Test_put(t *testing.T) {
	Put("myqvLpN5vmcSuwEMV4cfRHAL88Aq5MmofQ")
}
