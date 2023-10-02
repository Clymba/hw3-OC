package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Ошибка чтения")
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var count int

	regexpPattern := "[a-fA-F]"
	re := regexp.MustCompile(regexpPattern)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		count += len(matches)
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Ошибка чтения")
		os.Exit(2)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(fmt.Sprintf("Число символов %d\n", count))
	if err != nil {
		fmt.Println("Ошибка чтения")
		os.Exit(2)
	}

	fmt.Println("Исполняемый файл завершен")
	os.Exit(0)
}
