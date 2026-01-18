package internal

import (
	"os"
	"path/filepath"
	"strings"
)

func appDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}

func SaveUser(username string) error {
	dir, err := appDir()
	if err != nil {
		return err
	}

	file := filepath.Join(dir, "user")
	return os.WriteFile(file, []byte(username), 0644)
}

func LoadUser() (string, error) {
	dir, err := appDir()
	if err != nil {
		return "", err
	}

	file := filepath.Join(dir, "user")
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
