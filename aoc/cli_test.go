package main

import (
	"testing"
)

func TestParsingArgs(t *testing.T) {
	t.Run("Error if no args in inputs", func(t *testing.T) {
		args := []string{"command"}
		initArgs = mockFlagArgs(args)

		if _, _, err := run(); err == nil {
			t.Error("Should return error if no args")
		}
	})

	t.Run("First argument is fetch", func(t *testing.T) {
		args := []string{"fetch", "url"}
		initArgs = mockFlagArgs(args)

		if _, _, err := run(); err != nil {
			t.Errorf("First argument as fetch should be valid. Got error: %s", err.Error())
		}
	})

	t.Run("Error if first argument not fetch", func(t *testing.T) {
		var err error
		args := []string{"notfetch"}
		initArgs = mockFlagArgs(args)
		expectedErrMsg := "Did you want to call \"fetch\"?"

		if _, _, err = run(); err == nil {
			t.Error("Expected an error")
		}

		errMsg := err.Error()
		if errMsg != expectedErrMsg {
			t.Errorf("Expected error: %s, but got error: %s", expectedErrMsg, errMsg)
		}
	})

	t.Run("Error if no url passed as second argument", func(t *testing.T) {
		var err error
		args := []string{"fetch"}
		initArgs = mockFlagArgs(args)
		expectedErrMsg := "Please enter a url"

		if _, _, err = run(); err == nil {
			t.Error("Expected and error")
		}

		errMsg := err.Error()
		if errMsg != expectedErrMsg {
			t.Errorf("Expected error: %s, but got error: %s", expectedErrMsg, errMsg)
		}
	})

	t.Run("Should return url", func(t *testing.T) {
		args := []string{"fetch", "url"}
		initArgs = mockFlagArgs(args)

		url, _, err := run()
		if err != nil {
			t.Errorf("Should not have error, got error: %s", err.Error())
		}

		if url != args[1] {
			t.Errorf("Should return %s, got %s", args[1], url)
		}
	})

	t.Run("Should return a session field from session flag", func(t *testing.T) {
		args := []string{"fetch", "url"}
		initArgs = mockFlagArgs(args)

		defaultSessionFlagVal := "./session"

		_, sessionID, err := run()
		if err != nil {
			t.Errorf("Should not have an error, got error: %s", err.Error())
		}

		if sessionID != defaultSessionFlagVal {
			t.Errorf("Should have session: %s, but got session: %s", defaultSessionFlagVal, sessionID)
		}
	})
}

func TestGrabSessionId(t *testing.T) {
	t.Run("Returns error if no cookie found", func(t *testing.T) {
		readFile = mockReadFile([]byte{}, nil)

		sessionID, err := grabSessionID("")
		if err == nil {
			t.Errorf("Expected an error, got %s", sessionID)
		}
	})

	t.Run("Returns sessionID from file when it exists", func(t *testing.T) {
		mockFile := []byte{'a', 'b', 'c'}
		expected := "abc"
		pathToFile := "/path"
		readFile = mockReadFile(mockFile, nil)

		sessionID, err := grabSessionID(pathToFile)
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
		pathToFile := "/path"
		readFile = mockReadFile([]byte(mockFile), nil)

		sessionID, err := grabSessionID(pathToFile)
		if err != nil {
			t.Errorf("Expected no error, got %s", err.Error())
		}
		if sessionID != expected {
			t.Errorf("Expected %s, got %s", expected, sessionID)
		}
	})

	t.Run("Should return session id from params", func(t *testing.T) {
		expectedID := "abc123"

		actualID, err := grabSessionID(expectedID)
		if err != nil {
			t.Errorf("Expected no error, got error: %s", err.Error())
		}

		if actualID != expectedID {
			t.Errorf("Expexted id: %s, got id: %s", expectedID, actualID)
		}
	})

	t.Run("Should grab session if from file if passed a path", func(t *testing.T) {
		expectedId := "abc123"
		readFile = mockReadFile([]byte(expectedId), nil)
		idPath := "/file/path"

		sessionID, err := grabSessionID(idPath)
		if err != nil {
			t.Errorf("Expected no error, got error: %s", err.Error())
		}

		if sessionID != expectedId {
			t.Errorf("Expected sessionID: %s, got sessionID: %s", expectedId, sessionID)
		}
	})
}

func TestCheckingIfPath(t *testing.T) {
	t.Run("Should be false for regular string", func(t *testing.T) {
		input := "abc123"

		if isPath(input) {
			t.Errorf("Expected input: %s to be false", input)
		}
	})

	t.Run("Should be true for absolute path string", func(t *testing.T) {
		input := "/path/to/file"

		if !isPath(input) {
			t.Errorf("Expected input: %s to be true", input)
		}
	})

	t.Run("Should be true for relative path", func(t *testing.T) {
		input := "./path/to/file"

		if !isPath(input) {
			t.Errorf("Expected input: %s to be true", input)
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
