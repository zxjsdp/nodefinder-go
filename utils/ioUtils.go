package utils

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"io/ioutil"
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

func CheckFileExists(fileName, description, usage string) {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		log.Fatal(fmt.Sprintf("[ %s ] No file name provided! %s",
			description, usage))
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("[ %s ] File does not exist: %s. %s",
				description, fileName, usage))
		}
	}
}

func WriteContent(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
