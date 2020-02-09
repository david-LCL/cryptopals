package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"unicode"

	"github.com/steakknife/hamming"
)

//CountBitsByte from https://github.com/steakknife/hamming Copyright © 2014, 2015, 2016, 2018 Barry Allard

func getHamming(a []byte, b []byte) int {
	count := 0
	for i := 0; i < len(a); i++ {
		a[i] ^= b[i]

		count += hamming.CountBitsByte(a[i])
	}
	return count
}

func testHamming() bool {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")
	var h int = 37
	var c int = 0

	c = getHamming(a, b)

	if c == h {
		return true
	}

	fmt.Printf("Computed %d instead of %d\n", c, h)
	return false

}

//getCipherText decode base64'd data from file
func getCipherText(filepath string) []byte {
	var raw []byte
	var temp []byte
	lines, err := readLines(filepath)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, v := range lines {
		temp, _ = base64.StdEncoding.DecodeString(v)
		raw = append(raw, temp...)
	}
	return raw
}

type keyGuess struct {
	len int
	hd  float64
}

//guessKeySize uses hamming distance to find keylength on ciphertext
func guessKeySize(ksmax int, ct []byte) []int {
	var temp float64
	var kg0, kg1, kg2 keyGuess //top 3 guesses for keysize
	var kg []int               // return value
	kg0.hd = 1000              // initialize high for 1st comparison

	if len(ct) < 2*ksmax { //refactor max key size if there is not enough ciphertext
		ksmax = len(ct) / 2
	}

	for i := 2; i+2 < ksmax; i++ { //min keylenth is 2 bytes ie 1 character
		numBlocks := len(ct) / i
		for j := 1; j+2 < numBlocks; j += i {
			temp += float64(getHamming(ct[j:j+1], ct[j+1:j+2])) // get hd for each block within the ct
		}
		temp = temp / float64(i) //normalize the result

		if temp < kg0.hd {
			kg2.hd = kg1.hd
			kg2.len = kg1.len
			kg1.hd = kg0.hd
			kg1.len = kg0.len
			kg0.hd = temp
			kg0.len = i
		}
		temp = 0.0 //reset for next iteration

	}

	kg = append(kg, kg0.len)
	kg = append(kg, kg1.len)
	kg = append(kg, kg2.len)
	return kg
}

//guessXORKey tries to find the XOR cipher value based on language frequency analysis
func guessXORKey(block []byte, langmap map[string]float64) byte {
	var amap map[int][]byte
	amap = make(map[int][]byte)
	var myGuess Guess
	myGuess.score = -9999.0

	for m := 0; m < 255; m++ {
		amap[m] = XOREncipher(block, byte(m))
		tscore := scoreString(string(amap[m]), engMap)
		if unicode.IsPrint(rune(m)) { //checks that cipher is printable value
			if tscore > myGuess.score {
				myGuess.score = tscore
				myGuess.message = amap[m]
				myGuess.cipher = byte(m)
			}
		}
	}
	return myGuess.cipher
}

//transposeBlocks returns maps from data slice
func transposeBlocks(ks int, ct []byte) map[int][]byte {

	var tmap map[int][]byte
	tmap = make(map[int][]byte)
	numBlocks := len(ct) / ks

	for j := 0; j < numBlocks; j++ {
		for i := 0; i < ks && i+ks*j < len(ct); i++ {
			tmap[i] = append(tmap[i], ct[i+ks*j])
		}
	}
	return tmap

}

//breakRXOR tries to recover key from repeating xor cipher text with known keylength
func breakRXOR(ct []byte, ks int) []byte {
	var key []byte

	buf := transposeBlocks(ks, ct)
	for i := 0; i < len(buf); i++ {
		key = append(key, guessXORKey(buf[i], engMap))
	}

	return key
}

func challenge6() string {

	if !testHamming() {
		fmt.Printf("Hamming Fails\n")

	}

	var ct []byte
	var kg []int
	var pt map[int][]byte
	pt = make(map[int][]byte)

	ks := 40 //max suggested key length, min is 2 bytes hardcoded in guessKeySize
	winner := -999999
	score := -999999.0
	ct = getCipherText("6.txt")
	kg = guessKeySize(ks, ct)

	for i, v := range kg {

		if v == 0 {
			break
		}

		ans := breakRXOR(ct, v)

		pt[i] = ans
		score = scoreString(string(pt[i]), engMap)

		if score > float64(winner) {
			winner = i

		}
	}

	return prettyPrint(string(pt[int(winner)]))
}
