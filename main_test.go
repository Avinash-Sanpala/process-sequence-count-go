package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestProcessSequenceCountFromFile(t *testing.T) {
	input := "This is a test. \nTest input for processing."
	expectedCounts := map[string]int{
		"this is a":       1,
		"is a test":       1,
		"a test test":     1,
		"test test input": 1,
		"test input for":  1}
	reader := strings.NewReader(input)

	sequenceCounts := make(map[string]int)
	err := processSequenceCountFromFile(reader, sequenceCounts)
	if err != nil {
		t.Errorf("Error while processing sequence count: %s", err)
	}

	for sequence, count := range sequenceCounts {
		fmt.Printf("sequence: %s  count: %d\n", sequence, count)
	}

	for seq, expectedCount := range expectedCounts {
		if count, ok := sequenceCounts[seq]; !ok || count != expectedCount {
			t.Errorf("Expected sequence %s to have count %d, but got %d", seq, expectedCount, count)
		}
	}
}

func TestHeapSortTopSequences(t *testing.T) {
	sequences := []SequenceCount{
		{"sequence1", 5},
		{"sequence2", 2},
		{"sequence3", 8},
		{"sequence4", 3},
	}

	expectedSequences := []SequenceCount{
		{"sequence3", 8},
		{"sequence1", 5},
		{"sequence4", 3},
		{"sequence2", 2},
	}

	heapSortTopSequences(sequences)

	for i := range sequences {
		if sequences[i] != expectedSequences[i] {
			t.Errorf("Expected sequence %v, but got %v", expectedSequences[i], sequences[i])
		}
	}
}

func TestGetMaxSequences(t *testing.T) {
	tests := []struct {
		sequencesCount int
		expected       int
	}{
		{105, 100},
		{50, 50},
		{0, -1},
	}

	for _, test := range tests {
		result := getMaxSequences(test.sequencesCount)
		if result != test.expected {
			t.Errorf("For sequences count %d, expected %d but got %d", test.sequencesCount, test.expected, result)
		}
	}
}

func TestIsFileFromStdin(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"program"}

	if isFileFromStdin() {
		t.Error("Expected true when command line arguments are present, but got false")
	}

	os.Args = []string{"program", "arg1"}

	if !isFileFromStdin() {
		t.Error("Expected false when command line arguments are present, but got true")
	}
}
