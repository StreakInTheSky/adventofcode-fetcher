package cli

import "testing"

func TestCli(t *testing.T) {
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