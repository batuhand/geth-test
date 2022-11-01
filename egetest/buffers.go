package main

import (
	"github.com/ethereum/go-ethereum/rlp"
)

//Testing encode bytes function and its output

func testEncodeBytes() ([]byte, error) {
	//var buf bytes.Buffer
	list := [3]string{"Ege", "Ahmet", "Mehmet"}
	return rlp.EncodeToBytes(list)
}
