package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

//HexStringToBase64 converts hexstring base64 string
func HexStringToBase64(s string) string {
	hs, err0 := hex.DecodeString(s)
	if err0 != nil {
		log.Fatal(err0)
	}
	out := base64.StdEncoding.EncodeToString(hs)
	return out
}

func checkAnswer(exAnswer string, myGuess string) bool {
	myGuess = base64.StdEncoding.EncodeToString([]byte(myGuess))
	check := false
	if strings.Contains(myGuess, exAnswer) {
		check = true
	}
	return check
}

func prettyPrint(myString string) string {
	myString = base64.StdEncoding.EncodeToString([]byte(myString))

	raw, err := base64.StdEncoding.DecodeString(myString)
	if err != nil {
		panic(err)
	}
	return join("base64 ", myString, "\n", "ascii ", string(raw), "\n")
}

//challenge1 base64 encoding hex string
func challenge1() string {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var answer string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	output := HexStringToBase64(input)
	if strings.Contains(answer, output) {
		//return prettyPrint(output)
		raw, err := base64.StdEncoding.DecodeString(output)
		if err != nil {
			panic(err)
		}
		return (join("base64 ", output, "\n", string(raw), "\n"))

	}
	return "base64 encoding error \n"

}

func main() {
	fmt.Println("Running Challenge 1")
	fmt.Println(challenge1())
	fmt.Println("Running Challenge 2")
	fmt.Println(challenge2())
	fmt.Printf("Running Challenge 3\n")
	fmt.Println(challenge3())
	fmt.Printf("Running Challenge 4\n")
	fmt.Printf(challenge4())
	fmt.Printf("\nRunning Challenge 5\n")
	fmt.Printf(challenge5())
	fmt.Printf("\nRunning Challenge 6\n")
	fmt.Printf(challenge6())
}
