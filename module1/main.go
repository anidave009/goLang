package main

import (
	"fmt"
	"gopractice/module1/anagramcheck"
)

func main() {
	// fmt.Println(commaformatter.AddCommasToNumber("1a23456789"))
	// fmt.Println(commaformatter.AddCommasToNumber("11.456.21"))
	// fmt.Println(commaformatter.AddCommasToNumber("123456789"))
	// fmt.Println(commaformatter.AddCommasToNumber("-123456789"))
	// fmt.Println(commaformatter.AddCommasToNumber("+123456789"))
	// fmt.Println(commaformatter.AddCommasToNumber("1234567.89"))
	// fmt.Println(commaformatter.AddCommasToNumber("-1234567.89"))
	// fmt.Println(commaformatter.AddCommasToNumber("11234567"))
	// fmt.Println(commaformatter.AddCommas("11234567"))

	fmt.Println(anagramcheck.IsAnagram("世界", "界世"))

	// if len(os.Args) < 2 {
	// 	fmt.Println("usage: wordfreq <filename>")
	// 	os.Exit(1)
	// }
	// filename := os.Args[1]
	// counts, err := wordfreqcounter.WordFrequencyCounter(filename)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "error:", err)
	// 	os.Exit(1)
	// }
	// wordfreqcounter.PrintFrequencies(counts)

	// elements := []string{"listen", "silent", "enlist", "google", "glooge", "evil", "vile", "live"}
	// results := strings.Join(elements, " - - - ")
	// fmt.Println("Elements:", results)

	// fmt.Println(joinstringsvariadic.MyJoin(",", "Go", "Java", "Python"))
	// fmt.Println(joinstringsvariadic.MyJoin(",", "Go"))
	// fmt.Println(joinstringsvariadic.MyJoin(","))
	// fmt.Println(joinstringsvariadic.MyJoin("", "Go", "Java"))
	// fmt.Println(joinstringsvariadic.MyJoin("", ""))
	// fmt.Println(joinstringsvariadic.MyJoin("", "", ""))
	// fmt.Println(joinstringsvariadic.MyJoin(",", "", ""))
	// fmt.Println(joinstringsvariadic.MyJoin(""))

	// iota.PrintIotaConstants()
}

// Qn 1 text
// 123,456,789
// -123,456,789
// +123,456,789
// 1,234,567.89
// -1,234,567.89
// 11,234,567
