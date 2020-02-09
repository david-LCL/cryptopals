package main

import (
	"encoding/hex"
	"log"
	"strings"
)

//HexStringToBytes converts string to bytes
func HexStringToBytes(s string) []byte {
	out, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

//XORBytes xor two bytes
func XORBytes(a []byte, b []byte, c int) []byte {
	//c is the length of the larger byte array
	for i := 0; i < c; i++ {
		a[i] ^= b[i]
	}
	return a
}

//challenge2 XOR equal length inputs
func challenge2() string {
	var input1 string = "1c0111001f010100061a024b53535009181c"
	var input2 string = "686974207468652062756c6c277320657965"
	var answer string = "746865206b696420646f6e277420706c6179"
	var result string = "foo"
	var c int = 0

	var output1 []byte = HexStringToBytes(input1)
	var output2 []byte = HexStringToBytes(input2)
	if len(output1) < len(output2) {
		c = len(output2)
	} else {
		c = len(output1)
	}
	var output []byte = XORBytes(output1, output2, c)
	result = hex.EncodeToString(output)
	if strings.Contains(answer, result) {
		data, _ := hex.DecodeString(result)

		return prettyPrint(string(data))
	}
	return "XOR encoding error \n"

}
