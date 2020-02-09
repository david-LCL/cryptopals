package main

import (
	"strings"
	"testing"
)

func TestChallenge6(t *testing.T) {

	var dataItem5 TestDataItem5
	dataItem5.input = "Ciphertext from file"
	dataItem5.result = "Terminator X: Bring the noise"
	dataItem5.key = "Terminator X: Bring the noise"

	result := challenge6()

	if !strings.Contains(result, dataItem5.result) {
		t.Errorf("Breaking XOR key encryption failed, expected %v\ngot %v\n", dataItem5.result, result)
	} else {
		t.Logf("Breaking XOR key encryption success, expected %v\ngot %v\n", dataItem5.result, result)
	}
}
