package main

import (
	"bufio"
	"encoding/hex"

	//"fmt"
	"log"
	"os"
)

func guessXORCipher(input []byte) Guess {
	var myAnswer Guess
	var mapres map[int][]byte
	mapres = make(map[int][]byte)
	for j := 0; j < 255; j++ {

		mapres[j] = XOREncipher(input, byte(j))

	}
	myAnswer = guessEnglish(mapres, engMap)

	return myAnswer
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func challenge4() string {

	var myAnswer Guess
	var temp Guess
	myAnswer.score = -9999

	lines, err := readLines("4.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	//ans := scoreString("nOW THAT THE PARTY IS JUMPIN*", engMap)
	//fmt.Printf("correct answer is %f\n", ans)
	for _, v := range lines {
		buf, _ := hex.DecodeString(v)
		temp = guessXORCipher(buf)

		if temp.score > myAnswer.score {

			myAnswer.score = temp.score
			myAnswer.cipher = temp.cipher
			myAnswer.message = temp.message
			//fmt.Printf("current score is %f\n message is %s\n", temp.score, string(temp.message))
		}
	}
	return prettyPrint(string(myAnswer.message))

}
