package utils

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func ReadLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadContent(filename string) string {
	lines := ReadLines(filename)

	return strings.Join(lines, "\n")
}

func ReadCleanContent(filename string, runeListToBeRemoved []rune) string {
	rawContent := ReadContent(filename)

	for _, rune := range runeListToBeRemoved {
		rawContent = RemoveChar(rawContent, rune)
	}

	return rawContent
}
