package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOutputsToFile(t *testing.T) {
	mockData := `line1
	line2
	line3
	`
	tempFile, _ := os.CreateTemp("", "test-inputs.txt")
	defer os.Remove(tempFile.Name())

	if err := handleOutput(bytes.NewReader([]byte(mockData)), tempFile); err != nil {
		t.Error(err)
	}

	fileContent, _ := os.ReadFile(tempFile.Name())
	contentString := string(fileContent)
	if contentString != mockData {
		t.Errorf("Expected %s, but got: %s", mockData, contentString)
	}
}
