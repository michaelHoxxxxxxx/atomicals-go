package atomicals

import (
	"time"

	"github.com/atomicals-go/atomicals-indexer/atomicals-core/witness"
	"github.com/atomicals-go/pkg/log"
	"github.com/atomicals-go/repo/postsql"
	"github.com/atomicals-go/utils"
	"github.com/btcsuite/btcd/btcjson"
)

func (m *Atomicals) Run() {
	startTime := time.Now()

	location, err := m.Location()
	if err != nil {
		log.Log.Panicf("Location err:%v", err)
	}
	maxBlockHeight, err := m.GetBlockCount()
	if err != nil {
		log.Log.Panicf("GetBlockCount err:%v", err)
	}
	if location.BlockHeight+utils.SafeBlockHeightInterupt > maxBlockHeight {
		time.Sleep(10 * time.Minute)
	}
	if err := m.TraceBlock(location.BlockHeight, location.TxIndex); err != nil {
		return
	}
	log.Log.Infof("maxBlockHeight:%v, currentHeight:%v, time:%v", maxBlockHeight, location.BlockHeight, time.Since(startTime))
}
func (m *Atomicals) TraceBlock(height, txIndex int64) error {
	block, err := m.GetBlockByHeight(height)
	if err != nil {
		return err
	}
	if txIndex+1 >= int64(len(block.Tx)) {
		height++
		txIndex = -1
		block, err = m.GetBlockByHeight(height)
		if err != nil {
			return err
		}
	}
	for index := int64(txIndex + 1); index < int64(len(block.Tx)); index++ {
		tx := block.Tx[index]
		// parse this tx
		mod, deleteFts, newFts, updateNfts, newUTXOFtInfo,
			updateDistributedFt, newGlobalDistributedFt, newGlobalDirectFt, newUTXONftInfo := m.TraceTx(tx, block.Height)

		op := ""

		// update db if needed
		if mod != nil {
			m.InsertMod(mod)
			op = "mod"
		}
		for _, v := range deleteFts {
			if err := m.DeleteFtUTXO(v.LocationID); err != nil {
				log.Log.Panicf("DeleteFtUTXO err:%v", err)
			}
			op = "transfer_ft"
		}
		for _, v := range newFts {
			if err := m.InsertFtUTXO(v); err != nil {
				log.Log.Panicf("InsertFtUTXO err:%v", err)
			}
			op = "transfer_ft"
		}
		for _, v := range updateNfts {
			if err := m.UpdateNftUTXO(v); err != nil {
				log.Log.Panicf("UpdateNftUTXO err:%v", err)
			}
			op = "transfer_nft"
		}
		if newUTXOFtInfo != nil {
			if err := m.InsertFtUTXO(newUTXOFtInfo); err != nil {
				log.Log.Panicf("InsertFtUTXO err:%v", err)
			}
			op = "dmt"
		}
		if updateDistributedFt != nil {
			if err := m.UpdateDistributedFt(updateDistributedFt); err != nil {
				log.Log.Panicf("UpdateDistributedFtAmount err:%v", err)
			}
			op = "dmt"
		}
		if newGlobalDistributedFt != nil {
			if err := m.InsertDistributedFt(newGlobalDistributedFt); err != nil {
				log.Log.Panicf("InsertDistributedFt err:%v", err)
			}
			op = "dft"
		}
		if newGlobalDirectFt != nil {
			if err := m.InsertDirectFtUTXO(newGlobalDirectFt); err != nil {
				log.Log.Panicf("InsertDirectFtUTXO err:%v", err)
			}
			op = "ft"
		}
		if newUTXONftInfo != nil {
			if err := m.InsertNftUTXO(newUTXONftInfo); err != nil {
				log.Log.Panicf("InsertNftUTXO err:%v", err)
			}
			op = "nft"
		}

		// we don't need save all height-txid in db
		if index == 0 {
			m.DeleteBtcTxUntil(block.Height - utils.MINT_GENERAL_COMMIT_REVEAL_DELAY_BLOCKS)
		}
		if err := m.ExecAllSql(block.Height, index, tx.Txid, op); err != nil {
			log.Log.Panicf("ExecAllSql err:%v", err)
		}
	}
	return nil
}

func (m *Atomicals) TraceTx(tx btcjson.TxRawResult, height int64) (
	mod *postsql.ModInfo,
	deleteFts []*postsql.UTXOFtInfo, newFts []*postsql.UTXOFtInfo,
	updateNfts []*postsql.UTXONftInfo,
	newUTXOFtInfo *postsql.UTXOFtInfo, updateDistributedFt *postsql.GlobalDistributedFt,
	newGlobalDistributedFt *postsql.GlobalDistributedFt,
	newGlobalDirectFt *postsql.GlobalDirectFt,
	newUTXONftInfo *postsql.UTXONftInfo,
) {
	operation := witness.ParseWitness(tx, height)

	// step 1: insert mod
	if operation.Op == "mod" {
		mod = m.operationMod(operation, tx)
	}

	// step 2: transfer nft, transfer ft
	deleteFts, newFts, _ = m.transferFt(operation, tx)
	updateNfts, _ = m.transferNft(operation, tx)

	// step 3: process operation
	userPk := tx.Vout[utils.VOUT_EXPECT_OUTPUT_INDEX].ScriptPubKey.Address
	if operation.Op == "dmt" {
		newUTXOFtInfo, updateDistributedFt, _ = m.mintDistributedFt(operation, tx.Vout, userPk)
	} else {
		switch operation.Op {
		case "dft":
			newGlobalDistributedFt, _ = m.deployDistributedFt(operation, userPk)
		case "ft":
			newGlobalDirectFt, _ = m.mintDirectFt(operation, tx.Vout, userPk)
		case "nft":
			newUTXONftInfo, _ = m.mintNft(operation, userPk)
		case "evt":
		case "dat":
		case "sl":
		default:
		}
	}

	// step 4 check payment

	return mod,
		deleteFts, newFts,
		updateNfts,
		newUTXOFtInfo, updateDistributedFt,
		newGlobalDistributedFt,
		newGlobalDirectFt,
		newUTXONftInfo
}
