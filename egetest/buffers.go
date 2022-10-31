package main

import (
	"github.com/ethereum/go-ethereum/rlp"
)

func testEncodeBytes() ([]byte, error) {
	//var buf bytes.Buffer
	list := [3]string{"Ege", "Ahmet", "Mehmet"}
	return rlp.EncodeToBytes(list)
}
