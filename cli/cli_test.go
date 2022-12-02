package cli

import (
	"errors"
	"testing"
)

func TestParsingArgs(t *testing.T) {
	t.Run("Error if no args in inputs", func(t *testing.T) {
		args := []string{"command"}

		if _, err := ParseArgs(args); err == nil {
			t.Error("Should return error if no args")
		}
	})

	t.Run("First argument is fetch", func(t *testing.T) {
		args := []string{"command", "fetch", "url"}

		if _, err := ParseArgs(args); err != nil {
			t.Error("First argument as fetch should be valid")
		}
	})

	t.Run("Error if no url passed as second argument", func(t *testing.T) {
		args := []string{"command", "fetch"}

		if _, err := ParseArgs(args); err == nil {
			t.Error("Not passing a third argument should return an error")
		}
	})

	t.Run("Should return url", func(t *testing.T) {
		args := []string{"command", "fetch", "url"}

		url, err := ParseArgs(args)
		if err != nil || url != args[2] {
			t.Errorf("Should return %s, got %s", args[2], url)
		}
	})
}

func TestGrabbingSessionId(t *testing.T) {
	t.Run("Returns error if no cookie found", func(t *testing.T) {
		readFile = mockReadFile([]byte{}, nil)
		getEnv = mockGetEnv("")

		sessionID, err := GrabSessionId()
		if err == nil {
			t.Errorf("Expected an error, got %s", sessionID)
		}
	})

	t.Run("Returns sessionID from file when it exists", func(t *testing.T) {
		mockFile := []byte{'a', 'b', 'c'}
		expected := "abc"

		readFile = mockReadFile(mockFile, nil)
		getEnv = mockGetEnv("")

		sessionID, err := GrabSessionId()
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})

	t.Run("Returns sessionID from environment variable if file doesn't exist", func(t *testing.T) {
		mockValue := "abc"
		expected := mockValue

		readFile = mockReadFile([]byte{}, errors.New("No file"))
		getEnv = mockGetEnv(mockValue)

		sessionID, err := GrabSessionId()
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})

	t.Run("Returns sessionID from file if both file and env variable exist", func(t *testing.T) {
		mockFile := []byte{'a', 'b', 'c'}
		mockValue := "123"
		expected := "abc"

		readFile = mockReadFile(mockFile, nil)
		getEnv = mockGetEnv(mockValue)

		sessionID, err := GrabSessionId()
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})
}

func mockReadFile(file []byte, err error) func(string) ([]byte, error) {
	return func(path string) ([]byte, error) {
		return file, err
	}
}

func mockGetEnv(value string) func(string) string {
	return func(key string) string {
		return value
	}
}
