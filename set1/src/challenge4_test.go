package main

import (
	"strings"
	"testing"
)

func TestChallenge4(t *testing.T) {

	var dataItem4 TestDataItem2
	dataItem4.input1 = "foo"
	dataItem4.input2 = "bar"
	dataItem4.result = "Now that the party is jumping"

	result := challenge4()

	if !strings.Contains(result, dataItem4.result) {
		t.Errorf("XOR search failed, expected %v\ngot %v\n", dataItem4.result, result)
	} else {
		t.Logf("XOR search success, expected %v\ngot %v\n", dataItem4.result, result)
	}
}
