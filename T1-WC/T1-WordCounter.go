package main

import (
	"fmt"
	"strings"
	"sort"
	"io/ioutil"
	"unicode"
	)

type words struct {
	word string
	frec int
}

var text map[string]int
var wordCount []words


func ReadTxt() string {
	txt, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(txt)
}

func Fields (Text string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(Text, f)
}

func main() {

	var textByWords []string
	var tmpWord = words{}
	
	textByWords = Fields(ReadTxt())
	
	var frecuencyOfWords = make(map[string]int)

	// Frecuency
	for _, value := range textByWords {
		value = strings.ToLower(value)
		_, ok := frecuencyOfWords[value]
		if !(ok){ frecuencyOfWords[value] = 1 } else { frecuencyOfWords[value]++}
	}
	
	// Sort
	for j, value := range frecuencyOfWords {
		tmpWord.word = j
		tmpWord.frec = value
		wordCount = append(wordCount, tmpWord)
	}
	sort.SliceStable(wordCount, func(x, y int) bool{return wordCount[x].word < wordCount[y].word})
	
	// Print
	for _, tmp := range wordCount {
		fmt.Printf("%s (%d) \n", tmp.word, tmp.frec)
	}

}
