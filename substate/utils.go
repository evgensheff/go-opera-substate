package substate

import (
	stypes "github.com/Fantom-foundation/Substate/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//Here I put some utils to convert Geth types to Substate types

func HashGethToSubstate(g map[uint64]common.Hash) map[uint64]stypes.Hash {
	res := make(map[uint64]stypes.Hash)
	for k, v := range g {
		res[k] = stypes.Hash(v)
	}
	return res
}

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
