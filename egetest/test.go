package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func testHex(address string) ([]byte, string) {
	decoded, err := hexutil.Decode(address)
	decodeduint, erru := hexutil.DecodeBig(address)

	fmt.Println(decoded, err)
	fmt.Println(decodeduint, erru)

	encoded := hexutil.Encode(decoded)

	fmt.Println(encoded)

	return decoded, encoded
}

func testDecodeFunc(address string) []byte {
	src := []byte(address)

	fmt.Println(src)
	return src
}

func main() {
	//a := encodeArray()
	//fmt.Println(a)

	b, _ := testEncodeBytes()
	fmt.Println(b)

	//tx := createTx(100, 1667311668)
	tx := createTx(100, 1667311668)
	fmt.Printf("%+v\n", tx)
}
