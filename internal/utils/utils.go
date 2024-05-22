package utils

import (
	"io/ioutil"
	"os"
)

func WriteToFile(filename, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
