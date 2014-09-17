package main

import (
	"fmt"
	"encoding/hex"
)

const vanilla = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

const ice = "ICE"

func encrypt(plaintext, key []byte) (cyphertext []byte) {

	cyphertext = make([]byte, len(plaintext))

	for i, _ := range plaintext {
		cyphertext[i] = plaintext[i] ^ key[i % len(key)]
	}

	return
}

func main() {
	fmt.Println(hex.EncodeToString(encrypt([]byte(vanilla), []byte(ice))))
}