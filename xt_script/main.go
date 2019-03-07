package main

import (
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		exitWithErrMsg("usage: xt_script XTS_FILE")
	}
	fn := os.Args[1]
	src, err := ioutil.ReadFile(fn)
	if err != nil {
		exitWithErrMsg("read file [%s] failed [%s]", fn, err.Error())
	}

	tokenList := parseTokenList(string(src))
	stmtList := parseStmtList(tokenList, 0)
	stmtList.exec()
}
