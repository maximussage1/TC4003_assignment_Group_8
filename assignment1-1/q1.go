package cos418_hw1_1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	words := readStringFromFile(path)
	var wordMap map[string]int = groupWordsIntoAMapWithCount(words, charThreshold)
	var wordCounts []WordCount = transformsWordMapIntoSlice(wordMap)
	sortWordCounts(wordCounts)
	wordCounts = append(wordCounts[:numWords])
	return wordCounts
}

func transformsWordMapIntoSlice(wordMap map[string]int) []WordCount {
	var wordCounts []WordCount
	for k, v := range wordMap {
		wordCounts = append(wordCounts, WordCount{Count: v, Word: k})
	}
	return wordCounts
}

// Groups words into a Map with the amount of repetitions
func groupWordsIntoAMapWithCount(words []string, charThreshold int) map[string]int {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	var wordCount map[string]int = map[string]int{}
	for i := 0; i < len(words); i++ {
		currentVal := reg.ReplaceAllString(strings.ToLower(words[i]), "")
		if len(currentVal) < charThreshold {
			continue
		}
		if val, ok := wordCount[currentVal]; ok {
			val++
			wordCount[currentVal] = val
		} else {
			wordCount[currentVal] = 1
		}
	}
	return wordCount
}

// Tries to open the file and returns a string with its content
func readStringFromFile(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	return strings.Fields(text)
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
