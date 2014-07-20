package ks

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
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
