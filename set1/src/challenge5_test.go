package main

import (
	"strings"
	"testing"
)

type TestDataItem5 struct {
	input string

	result string

	key      string
	hasError bool
}

func TestChallenge5(t *testing.T) {

	var dataItem5 TestDataItem5
	dataItem5.input = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	dataItem5.result = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	dataItem5.key = "ICE"

	result := challenge5()

	if !strings.Contains(result, dataItem5.result) {
		t.Errorf("Repeating key XOR encryption failed, expected %v\ngot %v\n", dataItem5.result, result)
	} else {
		t.Logf("Repeating key XOR encryption success, expected %v\ngot %v\n", dataItem5.result, result)
	}
}
