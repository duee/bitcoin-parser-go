package service

import (
	"fmt"
	"pld_bitcoin_service/utils/db"
	"testing"
)

func TestGetBlcok(t *testing.T) {
	//open leveldb
	//db, err := leveldb.OpenFile("F:\\Bitcoin\\testnet3\\blocks\\index", opts) // You have got to dereference the pointer to get the actual value
	indexDb, err := db.OpenIndexDb(db.Datadir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer indexDb.Close()

	GetBlcok(indexDb, "0000000000003dd1ca5172cff71f62eda0bb7a986ea32b76be615e30e2bc3c49")
	//GetBlcok(indexDb,"000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943")
	//GetBlcok(indexDb,"00000000d1145790a8694403d4063f323d499e655c83426834d4ce2f8dd4a2ee")
}
