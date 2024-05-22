package utils

import (
	"os"
)

func WriteToFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
