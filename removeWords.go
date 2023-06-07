package main

import "strings"

func removeWords(text string, wordsToRemove []string) string {
	for _, word := range wordsToRemove {
		text = strings.ReplaceAll(text, word, "")
	}
	return text
}
