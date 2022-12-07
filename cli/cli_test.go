package cli

import (
	"errors"
	"testing"
)

func TestParsingArgs(t *testing.T) {
	t.Run("Error if no args in inputs", func(t *testing.T) {
		args := []string{"command"}
		initArgs = mockFlagArgs(args)

		if _, err := Run(); err == nil {
			t.Error("Should return error if no args")
		}
	})

	t.Run("First argument is fetch", func(t *testing.T) {
		args := []string{"fetch", "url"}
		initArgs = mockFlagArgs(args)

		if _, err := Run(); err != nil {
			t.Errorf("First argument as fetch should be valid. Got error: %s", err.Error())
		}
	})

	t.Run("Error if no url passed as second argument", func(t *testing.T) {
		args := []string{"fetch"}
		initArgs = mockFlagArgs(args)

		if _, err := Run(); err == nil {
			t.Error("Not passing a third argument should return an error")
		}
	})

	t.Run("Should return url", func(t *testing.T) {
		args := []string{"fetch", "url"}
		initArgs = mockFlagArgs(args)

		params, err := Run()
		if err != nil || params.Url != args[1] {
			t.Errorf("Should return %s, got %s", args[1], params.Url)
		}
	})
}

func TestGrabbingSessionId(t *testing.T) {
	t.Run("Returns error if no cookie found", func(t *testing.T) {
		readFile = mockReadFile([]byte{}, nil)
		getEnv = mockGetEnv("")

		sessionID, err := GrabSessionID()
		if err == nil {
			t.Errorf("Expected an error, got %s", sessionID)
		}
	})

	t.Run("Returns sessionID from file when it exists", func(t *testing.T) {
		mockFile := []byte{'a', 'b', 'c'}
		expected := "abc"

		readFile = mockReadFile(mockFile, nil)
		getEnv = mockGetEnv("")

		sessionID, err := GrabSessionID()
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

		sessionID, err := GrabSessionID()
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

		sessionID, err := GrabSessionID()
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})

	t.Run("Must return a single line string", func(t *testing.T) {
		mockFile := `line1
		line2
		line3
		`
		expected := "line1"

		readFile = mockReadFile([]byte(mockFile), nil)
		getEnv = mockGetEnv("")

		sessionID, err := GrabSessionID()
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})
}

func mockFlagArgs(args []string) func() []string {
	return func() []string {
		return args
	}
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
