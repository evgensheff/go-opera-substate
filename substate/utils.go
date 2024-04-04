package substate

import (
	"github.com/Fantom-foundation/Substate/substate"
	stypes "github.com/Fantom-foundation/Substate/types"
	"github.com/Fantom-foundation/Substate/types/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// Utils to convert Geth types to Substate types

// HashGethToSubstate converts map of geth's common.Hash to Substate hashes map
func HashGethToSubstate(g map[uint64]common.Hash) map[uint64]stypes.Hash {
	res := make(map[uint64]stypes.Hash)
	for k, v := range g {
		res[k] = stypes.Hash(v)
	}
	return res
}

// AccessListGethToSubstate converts geth's types.AccessList to Substate types.AccessList
func AccessListGethToSubstate(al types.AccessList) stypes.AccessList {
	st := stypes.AccessList{}
	for _, tuple := range al {
		var keys []stypes.Hash
		for _, key := range tuple.StorageKeys {
			keys = append(keys, stypes.Hash(key))
		}
		st = append(st, stypes.AccessTuple{Address: stypes.Address(tuple.Address), StorageKeys: keys})
	}
	return st
}

// LogsGethToSubstate converts slice of geth's *types.Log to Substate *types.Log
func LogsGethToSubstate(logs []*types.Log) []*stypes.Log {
	var ls []*stypes.Log
	for _, log := range logs {
		l := new(stypes.Log)
		l.BlockHash = stypes.Hash(log.BlockHash)
		l.Data = log.Data
		l.Address = stypes.Address(log.Address)
		l.Index = log.Index
		l.BlockNumber = log.BlockNumber
		l.Removed = log.Removed
		l.TxHash = stypes.Hash(log.TxHash)
		l.TxIndex = log.TxIndex
		for _, topic := range log.Topics {
			l.Topics = append(l.Topics, stypes.Hash(topic))
		}

		ls = append(ls, l)
	}
	return ls
}

// NewEnv prepares *substate.Env from ether's Block
func NewEnv(etherBlock *types.Block, statedb *state.StateDB) *substate.Env {
	return substate.NewEnv(
		stypes.Address(etherBlock.Coinbase()),
		etherBlock.Difficulty(),
		etherBlock.GasLimit(),
		etherBlock.NumberU64(),
		etherBlock.Time(),
		etherBlock.BaseFee(),
		HashGethToSubstate(statedb.SubstateBlockHashes))
}

// NewMessage prepares *substate.Message from ether's Message
func NewMessage(msg *types.Message) *substate.Message {
	to := stypes.Address(msg.To().Bytes())
	dataHash := hash.Keccak256Hash(msg.Data())

	return substate.NewMessage(
		msg.Nonce(),
		msg.IsFake(),
		msg.GasPrice(),
		msg.Gas(),
		stypes.Address(msg.From()),
		&to,
		msg.Value(),
		msg.Data(),
		&dataHash,
		AccessListGethToSubstate(msg.AccessList()),
		msg.GasFeeCap(),
		msg.GasTipCap())
}

// NewResult prepares *substate.Result from ether's Receipt
func NewResult(receipt *types.Receipt) *substate.Result {
	return substate.NewResult(
		receipt.Status,
		receipt.Bloom.Bytes(),
		LogsGethToSubstate(receipt.Logs),
		stypes.Address(receipt.ContractAddress),
		receipt.GasUsed)
}
