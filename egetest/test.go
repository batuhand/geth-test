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
	testHex("0x203560aCa0Aa5AAc09d9708CE29b60Aa3E4366A8") //20 bytes returned
	testDecodeFunc("0x203560aCa0Aa5AAc09d9708CE29b60Aa3E4366A8")
}
