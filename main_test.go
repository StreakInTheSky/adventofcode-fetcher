package main

import (
	"bytes"
	"os"
	"strings"
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

func TestFileNotExist(t *testing.T) {
	exists, err := checkFileExist("test")
	if err != nil {
		t.Error(err)
	}
	if exists {
		t.Error("File should not exist")
	}
}

func TestErrorOutputFileExists(t *testing.T) {
	testFile, _ := os.CreateTemp("", "test-file")
	defer os.Remove(testFile.Name())

	exists, err := checkFileExist(testFile.Name())
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Error("File should exist")
	}
}

func TestCreateFile(t *testing.T) {
	fileName := "test"
	createFile = mockCreateFile
	file, err := createOutputFile(fileName)
	if err != nil {
		t.Errorf("Should not have error, got error: %s", err.Error())
	}
	defer os.Remove(file.Name())

	if !strings.Contains(file.Name(), fileName) {
		t.Errorf("Should have created a file with name: %s, got %s", fileName, file.Name())
	}
}

func TestErrorCreatingFileIfExists(t *testing.T) {
	testFile, _ := os.CreateTemp("", "test-file")
	defer os.Remove(testFile.Name())

	createFile = mockCreateFile

	expectedError := "inputs.txt already exists in the directory"

	createdFile, err := createOutputFile(testFile.Name())
	if err == nil {
		os.Remove(createdFile.Name())
		t.Log("Should have error")
		t.FailNow()
	}
	if err.Error() != expectedError {
		t.Errorf("Expectd error: %s, got error: %s", expectedError, err.Error())
	}
}

func mockCreateFile(name string) (file *os.File, err error) {
	filepath := strings.Split(name, "/")
	filename := filepath[len(filepath)-1]

	file, err = os.CreateTemp("", filename)
	if err != nil {
		return file, err
	}
	return file, nil
}
