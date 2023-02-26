package main

import (
	"strings"
	"regexp"
	"os"
)

func createMap(name string) (map[rune][]string, error) {
	asciiMap := make(map[rune][]string)
	file, err := os.ReadFile("banners/" + name)
	if err != nil {
		return asciiMap, err
	}
	asciiKey := rune(31)
	a := regexp.MustCompile(`\r\n|\n`)
	for _, c := range a.Split(string(file), -1) {
		if c == "" {
			asciiKey++
		} else {
			asciiMap[asciiKey] = append(asciiMap[asciiKey], c)
		}
	}
	return asciiMap, err
}

func printAsciiArt(text string, name map[rune][]string) string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	readSlice := strings.Split(text, "\n")
	result := ""
	for _, value := range readSlice {
		if value == "" {
			result = result + "\n"
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range value {
					result = result + name[char][i]
				}
				result = result + "\n"
			}
		}
	}
	return result
}
