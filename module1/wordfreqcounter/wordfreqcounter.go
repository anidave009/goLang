package wordfreqcounter

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func WordFrequencyCounter(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	frequency := make(map[string]int)

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords) //tokenize by words

	for input.Scan() {
		word := strings.ToLower(input.Text())

		word = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return -1
		}, word)

		if word != "" {
			frequency[word]++
		}
	}

	if err := input.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil, err
	}

	return frequency, nil
}

func PrintFrequencies(frequency map[string]int) {
	words := make([]string, 0, len(frequency))
	for word := range frequency {
		words = append(words, word)
	}
	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s\t%d\n", word, frequency[word])
	}
}
