package main

import (
	"strings"
	"testing"
)

func TestChallenge3(t *testing.T) {

	var dataItem3 TestDataItem2
	dataItem3.input1 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	dataItem3.input2 = "686974207468652062756c6c277320657965"
	dataItem3.result = "Cooking MC's like a pound of bacon"

	result := challenge3()

	if !strings.Contains(result, dataItem3.result) {
		t.Errorf("XOREncipher guess failed, expected %s\ngot %s\n", dataItem3.result, result)
	} else {
		t.Logf("XOREnchiper guess success, expected %s\ngot %s\n", dataItem3.result, result)
	}
}
