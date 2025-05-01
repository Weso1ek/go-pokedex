package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var cleanWords []string

	words := strings.Split(strings.TrimSpace(text), " ")

	for i, j := range words {
		fmt.Println(i, j)
		if j != "" {
			cleanWords = append(cleanWords, strings.ToLower(j))
		}
	}

	return cleanWords
}
