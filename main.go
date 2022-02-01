package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Input string missing!\n")
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

	input := strings.Fields(os.Args[1])
	wordsMap := make(map[string]string)
	for _, s := range input {
		wordsMap[s] = strings.ToLower(s)
	}

	result := strings.Join(difference(wordsMap, commonWordsMap), ", ")
	fmt.Printf("Following words are not in the 1000 most common list:\n%s\n", result)
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
