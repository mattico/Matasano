package main

import (
	"io/ioutil"
	"os"
	"sort"
)

const filename = "encrypted-text.txt"

// returns the number of different bits in two byte arrays
func hamming_distance(a, b []byte) (diff int) {
	for i, _ := range a {
		var j uint
		for j = 0; j < 8; j++ {
			abit := (a[i] & (1 << j)) >> j
			bbit := (b[i] & (1 << j)) >> j
			if abit != bbit {
				diff++
			}
		}
	}
	return
}

type hamArray []ham

type ham struct {
	key              int
	hamming_distance float64
}

func (h hamArray) Len() int {
	return len(h)
}

func (h hamArray) Less(i, j int) bool {
	return h[i].hamming_distance < h[j].hamming_distance
}

func (h hamArray) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}

// returns an array of keylengths sorted by edit distance
func find_key_lengths() (lengths []ham) {
	file, _ := os.Open(filename)
	data := make([]byte, 80)
	file.Read(data)

	ham_dists := make(hamArray, 40)

	for keysize := 1; keysize <= 40; keysize++ {
		ham_dists[keysize-1] = ham{
			keysize,
			float64(hamming_distance(data[:keysize], data[keysize:2*keysize])) / float64(keysize),
		}
	}

	sort.Sort(ham_dists)

	return ham_dists
}

func main() {

	key_lengths := find_key_lengths()

	ciphertext, _ := ioutil.ReadFile(filename)

	// test first 5 keylengths
	for k := 0; k < 5; k++ {
		keylength := key_lengths[k].key
		cipherblocks := make([][]byte, keylength)
		for i, _ := range cipherblocks {
			cipherblocks[i] = make([]byte, len(ciphertext)/3+1)
		}
		for i, v := range ciphertext {
			cipherblocks[i%keylength][i/3] = v
		}

		// solve each cipherblocks[i] as if it were a single-char xor

	}

}
