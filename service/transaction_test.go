package service

import (
	"fmt"
	"pld_bitcoin_service/utils/db"
	"testing"
)

func TestGetTx(t *testing.T) {
	//打开leveldb
	//db, err := leveldb.OpenFile("F:\\Bitcoin\\testnet3\\blocks\\index", opts) // You have got to dereference the pointer to get the actual value
	indexDb, err := db.OpenTxIndexDb(db.Datadir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer indexDb.Close()

	GetTx(indexDb, "1ed6a57a2586efe3e98a82694f9bd190b3664615bf1b14c8612d65d5f6c7eea1")
}
