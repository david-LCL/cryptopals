package main

import (
	"encoding/hex"
	"strings"
	"unicode"
)

//Guess for storing scores and cipher for a message
type Guess struct {
	message []byte
	cipher  byte
	score   float64
}

var engMap = map[string]float64{

	//top english chars by percent wikipedia https://en.wikipedia.org/wiki/Letter_frequency

	"a": .08,
	"b": .02,
	"c": .03,
	"d": .05,
	"e": .13,
	"f": .03,
	"g": .02,
	"h": .06,
	"i": .07,
	"j": .00153,
	"k": .01,
	"l": .05,
	"m": .03,
	"n": .07,
	"o": .07,
	"p": .02,
	"q": 0.00095,
	"r": .06,
	"s": .07,
	"t": .09,
	"u": .03,
	"v": .01,
	"w": .03,
	"x": 0.0015,
	"y": .02,
	"z": 0.0077,
}

func guessEnglish(input map[int][]byte, langMap map[string]float64) Guess {
	//check map of strings and return string with highest score based on language map
	var myGuesses [256]Guess
	var topGuess Guess
	topGuess.score = -999
	for i := 0; i < len(input); i++ {
		//	myGuesses[i].message = input[i]
		//	myGuesses[i].cipher = byte(i)
		myGuesses[i].score = scoreString(string(input[i]), langMap)
		if myGuesses[i].score > topGuess.score {
			topGuess.message = input[i] //myGuesses[i].message
			topGuess.cipher = byte(i)   //myGuesses[i].cipher
			topGuess.score = myGuesses[i].score
		}
	}
	return topGuess
}

//scoreString based on top character occurence passed as language map
func scoreString(input string, langMap map[string]float64) float64 {

	var score float64 = 0
	input = strings.ToLower(input)

	//var tgraphs = []string{"the", "and", "for", "are", "but", "not", "you", "all", "any", "can", "had", "her", "was", "one", "our", "out", "day", "get", "has", "him", "his", "how", "man", "new", "now", "old", "see", "two", "way", "who", "boy", "did", "its", "let", "put", "say", "she", "too", "use", "that", "with", "have", "this", "will", "your", "from", "they", "know", "want", "been", "good", "much", "some", "time"}
	var engDMap = map[string]float64{
		//top english digraphs by percent http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/digraphs.html
		"th":   .0152,
		"he":   .0128,
		"in":   .0094,
		"er":   .0094,
		"an":   .0082,
		"re":   .0068,
		"nd":   .0063,
		"at":   .0059,
		"on":   .0057,
		"nt":   .0056,
		"ha":   .0058,
		"es":   .0056,
		"st":   .0055,
		"en":   .0055,
		"ed":   .0053,
		"to":   .0052,
		"it":   .0050,
		"ou":   .0050,
		"ea":   .0047,
		"hi":   .0046,
		"is":   .0046,
		"or":   .0043,
		"ti":   .0034,
		"as":   .0033,
		"te":   .0027,
		"et":   .0019,
		"ng":   .0018,
		"of":   .0016,
		"al":   .0009,
		"de":   .0009,
		"se":   .0008,
		"le":   .0008,
		"sa":   .0006,
		"si":   .0005,
		"ar":   .0004,
		"ve":   .0004,
		"ra":   .0004,
		"ld":   .0002,
		"ur":   .0002,
		"the":  .0025,
		"and":  .0025,
		"for":  .0025,
		"are":  .0025,
		"but":  .0025,
		"not":  .0025,
		"you":  .0025,
		"all":  .0025,
		"any":  .0025,
		"can":  .0025,
		"had":  .0025,
		"her":  .0025,
		"was":  .0025,
		"one":  .0025,
		"our":  .0025,
		"out":  .0025,
		"day":  .0025,
		"get":  .0025,
		"has":  .0025,
		"him":  .0025,
		"his":  .0025,
		"how":  .0025,
		"man":  .0025,
		"new":  .0025,
		"now":  .0025,
		"old":  .0025,
		"see":  .0025,
		"two":  .0025,
		"way":  .0025,
		"who":  .0025,
		"boy":  .0025,
		"did":  .0025,
		"its":  .0025,
		"let":  .0025,
		"put":  .0025,
		"say":  .0025,
		"she":  .0025,
		"too":  .0025,
		"use":  .0025,
		"that": .0025,
		"with": .0025,
		"have": .0025,
		"this": .0025,
		"will": .0025,
		"your": .0025,
		"from": .0025,
		"they": .0025,
		"know": .0025,
		"want": .0025,
		"been": .0025,
		"good": .0025,
		"much": .0025,
		"some": .0025,
		"time": .0025,
	}

	sc := []string{"<", ">", "/", "{", "}", "|", "(", ")", "@", "#", "$", "%", "^", "&", "*", "_", "-", "+", "=", "\\", "`", "~"}
	//	zp := regexp.MustCompile(`\w`)
	//	res := zp.Split(input, -1)
	res := strings.Fields(input)

	for _, v := range sc {
		for _, w := range res {
			count := strings.Count(w, v)
			if count > 0 {
				score -= float64(count) * .125
			}
		}
	}

	for j, s := range engDMap {

		for _, z := range res {

			count := strings.Count(z, j)
			if count > 0 {
				score += s * float64(count)
			}

		}
	}

	for j, s := range langMap {

		for _, v := range res {
			count := strings.Count(v, j)
			if count > 0 {
				score += s * float64(count)
			}

		}
		for _, v := range input {
			/*	if !unicode.IsPrint(v) {
				score--

			}*/
			if v > unicode.MaxASCII {
				score-- //5
			}
		}

		if strings.Count(input, " ") < 3 {
			score -= 100
		}

	}

	return score
}

//XOREncipher xor input using a single byte cipher
func XOREncipher(input []byte, cipher byte) []byte {

	var buff = make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		buff[i] = input[i] ^ cipher
	}
	return buff
}

//join using strinbuilder https://www.calhoun.io/concatenating-and-building-strings-in-go/
func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func challenge3() string {

	var input, _ = hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	var answermap map[int][]byte
	answermap = make(map[int][]byte)

	var myGuess Guess

	for m := 0; m < 255; m++ {
		answermap[m] = XOREncipher(input, byte(m))
	}
	myGuess = guessEnglish(answermap, engMap)
	return prettyPrint(string(myGuess.message))
}
