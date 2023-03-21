package files

import (
	"TrojanUI/internal/paths"
	"os"
)

func ReadConfig() (string, error) {
	file, err := os.ReadFile(paths.ConfigPath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func WriteConfig(config string) error {
	err := os.WriteFile(paths.ConfigPath, []byte(config), 0644)
	if err != nil {
		return err
	}
	return nil
}
