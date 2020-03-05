package service

import (
	"fmt"
	"log"
	"pld_bitcoin_service/utils"
	"pld_bitcoin_service/utils/db"
)

func GetBlcok(indexDb *db.IndexDb, blockId string) (*utils.Block, error) {
	db.FailIfReindexing(indexDb)
	result, err := db.GetBlockIndexRecordByBigEndianHex(indexDb, blockId)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", result)

	block, err := utils.NewBlockFromFile(db.Datadir, utils.BLOCK_MAGIC_ID_TESTNET, uint32(result.NFile), result.NDataPos)
	if err != nil {
		log.Fatal(err)
	}
	block.Hash = blockId
	fmt.Printf("block:%+v\n", block)
	//fmt.Printf("First Txid: %s\n", hex.EncodeToString(utils.ReverseHex(block.Transactions[0].Txid())))
	return block, err
}
