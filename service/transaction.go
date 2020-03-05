package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"pld_bitcoin_service/utils"
	"pld_bitcoin_service/utils/db"
)

//解析tx
func GetTx(indexDb *db.IndexDb, txHash string) (*utils.Transaction, error) {
	f, _ := db.GetFlag(indexDb, []byte("txindex"))
	if !f {
		log.Fatal(errors.New("txindex is not enabled for your bitcoind"))
	}
	result, err := GetTxIndexRecordByBigEndianHex(indexDb, txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	//utils.BLOCK_MAGIC_ID_TESTNET 表示testnet3环境魔数
	tx, err := utils.NewTxFromFile(db.Datadir, utils.BLOCK_MAGIC_ID_TESTNET, uint32(result.NFile), result.NDataPos, result.NTxOffset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", tx)
	return tx, err
}

func GetTxIndexRecordByBigEndianHex(indexDb *db.IndexDb, txHash string) (*db.TxIndexRecord, error) {
	txHashInBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}
	// reverse hex to get the LittleEndian order
	txHashInBytes = utils.ReverseHex(txHashInBytes)

	return GetTxIndexRecord(indexDb, txHashInBytes)
}

func GetTxIndexRecord(indexDb *db.IndexDb, txHash []byte) (*db.TxIndexRecord, error) {
	db.FailIfReindexing(indexDb)
	fmt.Printf("tx: %v, %d bytes\n", txHash, len(txHash))

	// Get data
	data, err := indexDb.Get(append([]byte("t"), txHash...), nil)
	if err != nil {
		return nil, err
	}

	// Parse the raw bytes
	txRecord := db.NewTxIndexRecordFromBytes(data)

	return txRecord, nil
}
