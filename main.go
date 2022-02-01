package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	results := checkCommonWords(os.Args)
	fmt.Printf("Following words are not in the 1000 most common list:\n%s\n", results)
}

func checkCommonWords(args []string) string {
	if len(args) <= 1 {
		fmt.Printf("Input is missing. Either pass string or use '-p' to pass file with payload\n")
		os.Exit(0)
	}

	file, err := os.Open("1000commonwords.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var commonWords []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commonWords = append(commonWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commonWordsMap := make(map[string]string)
	for _, s := range commonWords {
		commonWordsMap[s] = s
	}

	input := args[1]

	var inputWords []string
	if strings.HasPrefix(input, "-p") {
		if len(args) <= 2 {
			fmt.Println("Expecing a file with -p")
			os.Exit(0)
		}
		inputWords = openPayload(args[2])
	} else {
		inputWords = strings.Fields(input)
	}

	wordsMap := make(map[string]string)
	for _, s := range inputWords {
		wordsMap[s] = strings.ToLower(s)
	}

	return strings.Join(difference(wordsMap, commonWordsMap), ", ")
}

func difference(map1, map2 map[string]string) []string {
	var diff []string

	for _, k := range map1 {
		// Check if key from first map exists in second one
		if _, ok := map2[k]; !ok {
			diff = append(diff, k)
		}
	}

	return diff
}

func openPayload(filePath string) []string {
	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words := strings.Fields(string(file))
	return words
}
