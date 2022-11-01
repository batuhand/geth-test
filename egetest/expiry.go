package main

import (
	"fmt"
	"math/big"
	"sync"
)

//arbitrary tx structure and testing expiring tx

var encodeBufferPool = sync.Pool{
	New: func() interface{} { return new(*Transaction) },
}

// assuming one type of transaction
type Transaction struct {
	value1          uint64
	value2          *big.Int
	expiryTimestamp uint
}

/*
type TxData interface {
	value1() uint64
	value2() *big.Int
	expiryTimestamp() uint
}

func (tx *Transaction) Value1() uint64        { return tx.inner.value1() }
func (tx *Transaction) Value2() *big.Int      { return new(big.Int).Set(tx.inner.value2()) }
func (tx *Transaction) ExpiryTimestamp() uint { return tx.inner.expiryTimestamp() }
*/

func createTx(value1 uint64, expiry uint) {
	//buftx := encodeBufferPool.Get().(*Transaction)
	//defer encodeBufferPool.Put(buftx)

	var val2 *big.Int
	val2.SetUint64(946326465324)

	buftx := Transaction{
		value1:          value1,
		value2:          val2,
		expiryTimestamp: expiry,
	}

	fmt.Printf("%+v/n", buftx)
	//return buftx
}
