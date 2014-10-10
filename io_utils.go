package ks

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadBytesFromFile(filename string) ([]byte, error) {
	fin, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(fin)
}

func ReadLines(r io.Reader) []string {

	br := bufio.NewReader(r)
	var res []string

	for {
		s, _, err := br.ReadLine()

		if err != nil {
			return res
		}

		res = append(res, string(s))

	}

}

func ListFiles(folder string) []string {

	var res []string

	filepath.Walk(fmt.Sprintf(folder),
		func(path string, info os.FileInfo, err error) error {

			if info != nil && !info.IsDir() {
				res = append(res, info.Name())
			}

			return nil
		})
	return res
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {

		return false
	}

	return true
}
