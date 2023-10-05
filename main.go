package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/labstack/gommon/log"
)

// SequenceCount represents a sequence and its associated count.
type SequenceCount struct {
	Sequence string
	Count    int
}

const (
	MaxSequences     = 100
	FilePath         = "input.txt"
	ErrFileScanning  = "Error occurred while scanning the file"
	ErrParsingRegexp = "Error occurred while parsing the regular expression"
	ErrFileOpening   = "Error occurred while opening the file"
)

// main is the entry point of the program.
func main() {

	sequences, err := processSequenceCounts()

	if err != nil {
		log.Print(err)
	}

	sortSequencesByCount(sequences)

	printHighestSequences(sequences)
}

// printHighestSequences prints the highest count sequences.
func printHighestSequences(sequences []SequenceCount) {
	for i := 0; i < getMaxSequences(len(sequences)); i++ {
		fmt.Println(sequences[i].Sequence, sequences[i].Count)
	}
}

// processSequenceCounts processes the sequence counts from the input file or stdin.
func processSequenceCounts() ([]SequenceCount, error) {
	sequenceCounts := make(map[string]int)
	var sequences []SequenceCount

	var err error

	if isFileFromStdin() {
		err = processFileFromStdin(sequenceCounts)
	} else {
		err = processFile(FilePath, sequenceCounts)
	}

	if err != nil {
		return sequences, err
	}

	for sequence, count := range sequenceCounts {
		sequences = append(sequences, SequenceCount{sequence, count})
	}

	return sequences, nil
}

// processFileFromStdin processes the sequence count from stdin.
func processFileFromStdin(sequenceCounts map[string]int) error {
	fileName := os.Args[2]
	return processFile(fileName, sequenceCounts)
}

// processFile processes the sequence count from a file.
func processFile(filePath string, sequenceCounts map[string]int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s\n", ErrFileOpening, err))
	}
	defer file.Close()

	return processSequenceCountFromFile(file, sequenceCounts)
}

// processSequenceCountFromFile processes sequence count from a reader.
func processSequenceCountFromFile(file io.Reader, sequenceCounts map[string]int) error {
	scanner := bufio.NewScanner(file)

	sequenceQueue := make([]string, 0, 3)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)

		regexCompile, err := regexp.Compile(`\\[nrt]|[^a-zA-Z\s]`)
		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s\n", ErrParsingRegexp, err))
		}

		lineWithoutSpecialCharacters := regexCompile.ReplaceAllString(line, " ")

		words := strings.Fields(lineWithoutSpecialCharacters)

		for _, word := range words {
			sequenceQueue = append(sequenceQueue, word)

			if len(sequenceQueue) > 3 {
				sequenceQueue = sequenceQueue[1:]
			}

			if len(sequenceQueue) == 3 {
				sequence := strings.Join(sequenceQueue, " ")
				sequenceCounts[sequence]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.New(fmt.Sprintf("%s: %s\n", ErrFileScanning, err))
	}

	return nil
}

// getMaxSequences returns the maximum number of sequences to consider.
func getMaxSequences(sequencesCount int) int {
	if sequencesCount >= MaxSequences {
		return MaxSequences
	} else if sequencesCount == 0 {
		return -1
	}
	return sequencesCount
}

// isFileFromStdin checks if the program is receiving input from stdin.
func isFileFromStdin() bool {
	if len(os.Args) > 1 {
		return true
	}
	return false
}

// ths is to mock and test conflict issue. these changes are pushed to feature
// sortSequencesByCount gets all the sequences sorted by count in descending order
// ths is to mock and test conflict issue. these changes are directly pushed to master
func sortSequencesByCount(sequences []SequenceCount) {
	sort.Slice(sequences, func(i, j int) bool {
		return sequences[i].Count > sequences[j].Count
	})
}
