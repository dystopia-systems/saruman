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