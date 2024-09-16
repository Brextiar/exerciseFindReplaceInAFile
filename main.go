package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceInFile(src, old, new, newFileName string) (occurrences int, lines []int, err error) {
	data, err := os.ReadFile(src)

	if err != nil {
		return 0, nil, err
	}

	if len(data) == 0 {
		return 0, nil, fmt.Errorf("empty content (filename=\"%v\")", src)
	}
	newFileName += ".txt"
	outputFile, err := os.Create(newFileName)
	if err != nil {
		return 0, nil, err
	}
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		found, result, occurences := ProcessLine(line, old, new)
		if !found {
			i++

			writer.WriteString(line)
			writer.WriteString("\n")
			fmt.Println(line)

			continue
		}
		lines = append(lines, i)
		occurrences += occurences
		fmt.Println(result)
		writer.WriteString(result)
		writer.WriteString("\n")
		i++
	}

	return occurrences, lines, err
}

func ProcessLine(line, old, new string) (found bool, result string, occurences int) {
	found = strings.Contains(line, old)
	if !found {

		return false, line, 0
	}
	occurences = strings.Count(line, old)
	result = strings.ReplaceAll(line, old, new)

	return found, result, occurences
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path to the file: ")
	PathToTheFile, _ := reader.ReadString('\n')
	PathToTheFile = strings.TrimSpace(PathToTheFile)
	fmt.Print("Enter the word to replace: ")
	wordToReplace, _ := reader.ReadString('\n')
	wordToReplace = strings.TrimSpace(wordToReplace)
	fmt.Print("Enter the new word: ")
	newWord, _ := reader.ReadString('\n')
	newWord = strings.TrimSpace(newWord)
	fmt.Println("Enter the new file name: ")
	newFileName, _ := reader.ReadString('\n')
	newFileName = strings.TrimSpace(newFileName)
	fmt.Println("--------------------------------------------")
	fmt.Println("New Content---------------------------------")
	occurrences, lines, err := FindReplaceInFile(PathToTheFile, wordToReplace, newWord, newFileName)
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)

		return
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("STATS---------------------------------------")
	fmt.Printf("Number of occurrences of %v : %d\n", wordToReplace, occurrences)
	fmt.Println("Numbers of lines : ", len(lines))
	fmt.Println("Lines:", lines)
}
