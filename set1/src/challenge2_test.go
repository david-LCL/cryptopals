package main

import (
	"encoding/hex"
	"testing"
)

type TestDataItem2 struct {
	input1   string
	input2   string
	result   string
	hasError bool
}

func TestXORBytes(t *testing.T) {

	var dataItem2 TestDataItem2
	dataItem2.input1 = "1c0111001f010100061a024b53535009181c"
	dataItem2.input2 = "686974207468652062756c6c277320657965"
	dataItem2.result = "746865206b696420646f6e277420706c6179"

	result := hex.EncodeToString(XORBytes(HexStringToBytes(dataItem2.input1), HexStringToBytes(dataItem2.input2), len(HexStringToBytes(dataItem2.input2))))

	if result != dataItem2.result {
		t.Errorf("XORBytes() failed, expected %v\ngot %v\n", dataItem2.result, result)
	} else {
		t.Logf("XORBytes() success, expected %v\ngot %v\n", dataItem2.result, result)
	}
}
