package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func convertToBytesFromHash(hash common.Hash) []byte {
	new := make([]byte, 32)
	copy(new, hash[:])
	return new
}

func changeBytesToHash(bytes []byte) *common.Hash {
	var hash common.Hash
	copy(hash[:], bytes)
	return &hash
}

func changeByteContent(bytes []byte) []byte {
	//0x6567650a - 4 bytes
	bytes[0] = 0x65
	bytes[1] = 0x67
	bytes[2] = 0x65
	bytes[3] = 0x0a

	return bytes
}

func changeTxHash() *common.Hash {
	tx_hash := common.HexToHash("0xdf4e14abd904447b9aec870b40e9fd620df39c7a9e740d40a3096de2d644e23e")
	bytes := convertToBytesFromHash(tx_hash)
	changed := changeByteContent(bytes)
	hash := changeBytesToHash(changed)
	fmt.Println(hash)
	return hash
}
