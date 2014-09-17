package main

import (
	"fmt"
	"math"
	"sort"
)

func decrypt(cyphertext []byte, key byte) []byte {
	retval := make([]byte, len(cyphertext))
	for i, _ := range cyphertext {
		retval[i] = cyphertext[i] ^ key
	}
	return retval
}

type frequency_tuple struct {
	data []byte
	freq map[byte]int
	key  byte
}

type by_E []frequency_tuple

func (t by_E) Len() int {
	return len(t)
}

func (t by_E) Less(i, j int) bool {
	return t[i].freq['E']+t[i].freq['e'] < t[j].freq['E']+t[j].freq['e']
}

func (t by_E) Swap(i, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
}

func count_chars(decrypted []byte) (ret map[byte]int) {
	ret = make(map[byte]int)
	for _, v := range decrypted {
		ret[v]++
	}

	return
}

// http://en.wikipedia.org/wiki/Letter_frequency
func main() {
	decrypted := make([]frequency_tuple, 256)
	for i := 0; i < math.MaxUint8; i++ {
		decrypted[byte(i)].data = decrypt(cyphertext, byte(i))
		decrypted[byte(i)].freq = count_chars(decrypted[byte(i)].data)
		decrypted[byte(i)].key = byte(i)
	}

	sort.Sort(by_E(decrypted))

	for _, v := range decrypted {
		if v.freq['E'] > 0 {
			fmt.Println(string(v.data), v.key)
		}
	}

}
