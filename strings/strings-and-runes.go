package main

import (
	"fmt"
	"unicode/utf8"
)

type Card struct {
	uuid           string
	unicode        string
	fullName       string
	suit           string
	cardName       string
	primaryValue   int
	secondaryValue int
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '\u0E2A' {
		fmt.Println("found so sua")
	}
}

func main() {
	const s = "\u0E2A\u0E35\u0E2A\u0E4C"
	fmt.Println("Len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
	fmt.Println("\u2660")
	fmt.Println("\u2665")
	fmt.Println("\u2666")
	fmt.Println("\u2663")
}
