package main

import (
	"os"
	"testing"
)

func TestCheckCommonWordsWithString(t *testing.T) {
	args := []string{"exe/main", "Use a for loop to iterate through each element in a list of strings by index"}
	result := checkCommonWords(args)

	expectedResult := "index, iterate, loop, strings"
	if result != expectedResult {
		t.Errorf("Returned non-common words were incorrect, got: %s, want: %s.", result, expectedResult)
	}
}

func TestCheckCommonWordsWithFile(t *testing.T) {
	args := []string{os.Args[0], "-p", "payload.txt"}
	result := checkCommonWords(args)

	expectedResult := "index, iterate, loop, strings"
	if result != expectedResult {
		t.Errorf("Returned non-common words were incorrect, got: %s, want: %s.", result, expectedResult)
	}
}

// TODO: add test for missing file when using -p
// TODO: add test for invalid input
