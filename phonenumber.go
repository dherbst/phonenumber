package main

import (
	"fmt"
	"os"
	"strings"
)

var letterMap = map[string]string{
	"A": "2",
	"B": "2",
	"C": "2",
	"D": "3",
	"E": "3",
	"F": "3",
	"G": "4",
	"H": "4",
	"I": "4",
	"J": "5",
	"K": "5",
	"L": "5",
	"M": "6",
	"N": "6",
	"O": "6",
	"P": "7",
	"Q": "7",
	"R": "7",
	"S": "7",
	"T": "8",
	"U": "8",
	"V": "8",
	"W": "9",
	"X": "9",
	"Y": "9",
	"Z": "9",
}

func isPhoneNumber(c string) bool {
	list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", ".", "(", ")"}
	for _, check := range list {
		if c == check {
			return true
		}
	}
	return false
}

func convert(c string) string {
	v := letterMap[c]
	return v
}

func main() {
	word := os.Args[1]
	fmt.Println(word)

	number := ""

	for i := 0; i < len(word); i++ {
		// assume you are on a terminal typing characters
		c := strings.ToUpper(string(word[i]))
		if isPhoneNumber(c) {
			number += c
		} else {
			x := convert(c)
			number += x
		}
	}
	fmt.Println(number)
}
