package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return content
}
