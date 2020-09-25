package service

import (
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func WriteFile(path string, content []byte) (bool, error) {
	err := ioutil.WriteFile(path, content, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}