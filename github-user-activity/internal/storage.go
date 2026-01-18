package internal

import (
	"os"
	"path/filepath"
	"strings"
)

func SaveUser(username string) error {
	path, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	file := filepath.Join(path, "github-activity", "user")
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}

	return os.WriteFile(file, []byte(username), 0644)
}

func LoadUser() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	file := filepath.Join(path, "github-activity", "user")
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
