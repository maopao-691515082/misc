package main

import (
	"os"
	"fmt"
)

func exitWithErrMsg(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
	os.Exit(1)
}

func listToSet(l []string) (s map[string]bool) {
	s = map[string]bool{}
	for _, v := range l {
		s[v] = true
	}
	return
}
