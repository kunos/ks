package ks

import (
	"fmt"
)

func TimeToString(t uint64) string {

	tr := t
	min := t / (60000)
	tr -= min * (60000)
	sec := tr / 1000
	tr -= sec * 1000
	mils := tr

	return fmt.Sprintf("%d:%02d:%03d", min, sec, mils)

}
