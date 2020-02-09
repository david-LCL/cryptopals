package main

import "encoding/hex"

//keyedXOR xor data with key
func keyedXOR(key []byte, data []byte) []byte {
	//pad data
	data = addPad(len(key), len(data), data)

	//repeating XOR on fixed key
	for j := 0; j < len(data); j += len(key) {
		for k := 0; k < len(key); k++ {
			data[k+j] = data[k+j] ^ key[k]
		}
	}
	//trim padding from keyed XOR data
	data = trimPad(len(key), len(data), data)
	return data
}

//addPad introduces padding to input data if length does not align with keylength
func addPad(klen int, dlen int, data []byte) []byte {
	padding := dlen % klen
	if padding != 0 {
		padding = klen - padding
	}

	for i := 0; i < padding; i++ {
		data = append(data, byte(5))
	}
	return data
}

func trimPad(klen int, dlen int, ct []byte) []byte {
	pad := dlen % klen
	if pad != 0 {
		pad = klen - pad
	}
	return ct[:(dlen - pad)]
}

func challenge5() string {
	i1 := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	c1 := []byte(keyedXOR(key, i1))

	a1 := hex.EncodeToString(c1)

	return prettyPrint(string(a1))

}
