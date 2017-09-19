/*
	The utils package provides commonly used tools.
*/
package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ReadLines reads a file and return all the lines.
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

// ReadContent reads a file and return file content as a string.
func ReadContent(filename string) string {
	lines := ReadLines(filename)

	return strings.Join(lines, "\n")
}

// ReadCleanContent reads a file and return clean content as a string.
func ReadCleanContent(filename string, runeListToBeRemoved []rune) string {
	rawContent := ReadContent(filename)

	for _, rune := range runeListToBeRemoved {
		rawContent = RemoveChar(rawContent, rune)
	}

	return rawContent
}

// CheckFileExists checks whether file exists.
func CheckFileExists(fileName, description, usage string) {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! No file name provided for [ %s ].%s",
			description, usage))
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("ERROR! File does not exist: [ %s ] for [ %s ].%s",
				description, fileName, usage))
		}
	}
}

// WriteContent write string content to a file.
func WriteContent(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
