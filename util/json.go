package util

import (
	"encoding/json"
	"os"
)

func ParseJSONFile(filename string, target interface{}) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(f, target); err != nil {
		return err
	}

	return nil
}
