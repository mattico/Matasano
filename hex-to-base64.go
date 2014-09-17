package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

const hex_string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {

	bytes, _ := hex.DecodeString(hex_string)

	base64_string := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println(base64_string)
}