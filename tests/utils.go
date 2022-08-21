package tests

import (
	"os"
	"strings"
)

func AreRunning() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}
