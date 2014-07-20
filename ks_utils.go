package ks

import (
	"os"
)

func SystemPause() bool {
	buf := make([]byte, 1)
	for {
		c, _ := os.Stdin.Read(buf)
		if c != 1 {
			return false
		}
		if buf[0] == '\n' {
			break
		}
	}
	return true
}
