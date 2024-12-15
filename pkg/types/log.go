package types

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type Log struct {
	blockNumber uint64
	value       *ethtypes.Log
}

func NewLog(blockNumber uint64, value *ethtypes.Log) *Log {
	return &Log{
		blockNumber: blockNumber,
		value:       value,
	}
}

func (l *Log) BlockNumber() uint64 {
	return l.blockNumber
}

func (l *Log) Value() *ethtypes.Log {
	return l.value
}
