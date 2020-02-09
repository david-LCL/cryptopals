package main

import "testing"

type TestDataItem struct {
	input    string
	result   string
	hasError bool
}

func TestHexStringToBase64(t *testing.T) {

	var dataItem TestDataItem
	dataItem.input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	dataItem.result = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result := HexStringToBase64(dataItem.input)

	if result != dataItem.result {
		t.Errorf("\nHexStringToBase64(\"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d\")\nfailed, expected %v\ngot %v\n", dataItem.result, result)
	} else {
		t.Logf("\nHexStringToBase64(\"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d\")\nsuccess, expected %v\ngot %v\n", dataItem.result, result)
	}
}
