package main

import (
	"regexp"
	"strings"
	"fmt"
	"strconv"
)

//解析三大类token的正则
var tokenRegExpTable = map[string]*regexp.Regexp{
	"sym":  regexp.MustCompile(`!=|==|&&|\|\||\W`),
	"int":  regexp.MustCompile(`\d\w*`),
	"name": regexp.MustCompile(`[a-zA-Z_]\w*`),
}

type tokenType struct {
	lineNo int
	tp     string
	v      interface{}
}

func (t tokenType) raise(format string, args ...interface{}) {
	exitWithErrMsg("line[%d]: %s", t.lineNo, fmt.Sprintf(format, args...))
}

func (t tokenType) isSym(expSym interface{}) bool {
	return t.tp == "sym" && (expSym == nil || t.v.(string) == expSym.(string))
}

func (t tokenType) isReserved(expReserved interface{}) bool {
	return t.tp == "reserved" && (expReserved == nil || t.v.(string) == expReserved.(string))
}

func (t tokenType) isName() bool {
	return t.tp == "name"
}

func (t tokenType) isInt() bool {
	return t.tp == "int"
}

func (t tokenType) isStr() bool {
	return t.tp == "str"
}

type tokenListType struct {
	tl []tokenType
	i  int
}

func (tl *tokenListType) add(t tokenType) {
	tl.tl = append(tl.tl, t)
}

func (tl *tokenListType) len() int {
	return len(tl.tl) - tl.i
}

func (tl *tokenListType) assertNotEnd() {
	if tl.len() <= 0 {
		exitWithErrMsg("unexpected end of code")
	}
}

func (tl *tokenListType) revert() {
	if tl.i <= 0 {
		panic("bug")
	}
	tl.i --
}

func (tl *tokenListType) revertTo(i int) {
	if i < 0 || tl.i <= i {
		panic("bug")
	}
	tl.i = i
}

func (tl *tokenListType) peek() tokenType {
	tl.assertNotEnd()
	return tl.tl[tl.i]
}

func (tl *tokenListType) pop() tokenType {
	tl.assertNotEnd()
	tl.i ++
	return tl.tl[tl.i - 1]
}

func (tl *tokenListType) popSym(expSym interface{}) string {
	t := tl.pop()
	if !t.isSym(expSym) {
		if expSym == nil {
			t.raise("expect sym")
		}
		t.raise("expect sym '%s'", expSym)
	}
	return t.v.(string)
}

func (tl *tokenListType) popName() string {
	t := tl.pop()
	if !t.isName() {
		t.raise("expect name")
	}
	return t.v.(string)
}

//常用转义码
var escapeCharMap = map[byte]byte{
	'a': '\a',
	'b': '\b',
	'f': '\f',
	'n': '\n',
	'r': '\r',
	't': '\t',
	'v': '\v',
}

func parseStr(lineNo int, line string, endQuota byte) (s string, consumeLen int) {
	b := []byte{}
	defer func () {
		s = string(b)
	}()
	for {
		if len(line) <= 0 {
			exitWithErrMsg("unexpected end of line at line [%d]", lineNo + 1)
		}
		c := line[0]
		switch c {
		case endQuota:
			consumeLen ++
			return
		case '\\':
			if len(line) < 2 {
				exitWithErrMsg("unexpected end of line at line [%d]", lineNo + 1)
			}
			ec := line[1]
			switch ec {
			case 'x':
				if len(line) < 4 {
					exitWithErrMsg("unexpected end of line at line [%d]", lineNo + 1)
				}
				h := line[2 : 4]
				n, e := strconv.ParseUint(h, 16, 8)
				if e != nil {
					exitWithErrMsg("invalid hex escape [\\x%s] at line [%d]", h, lineNo + 1)
				}
				consumeLen += 4
				line = line[4 :]
				b = append(b, byte(n))
			default:
				v, ok := escapeCharMap[ec]
				if !ok {
					v = ec
				}
				consumeLen += 2
				line = line[2 :]
				b = append(b, v)
			}
		default:
			consumeLen ++
			line = line[1 :]
			b = append(b, c)
		}
	}
	panic("bug")
}

//是否关键字
func isReserved(s string) (ok bool) {
	switch s {
	case "loop", "break", "continue", "if", "else":
		ok = true
	}
	return
}

func parseTokenList(src string) (tl *tokenListType) {
	tl = &tokenListType{}
	for lineNo, line := range strings.Split(src, "\n") {
		for {
			line = strings.Trim(line, "\x20\t")
			if len(line) <= 0 {
				break
			}
			ok := false
			for tp, re := range tokenRegExpTable {
				loc := re.FindStringIndex(line)
				if loc != nil && loc[0] == 0 {
					ok = true
					s := line[: loc[1]]
					line = line[loc[1] :]
					t := tokenType{
						lineNo: lineNo + 1,
						tp:     tp,
					}
					switch tp {
					case "sym":
						if s == "'" || s == `"` {
							t.tp = "str"
							v, consumeLen := parseStr(lineNo, line, s[0])
							t.v = v
							line = line[consumeLen :]
						} else {
							t.v = s
						}
					case "int":
						n, err := strconv.ParseInt(s, 0, 64)
						if err != nil {
							exitWithErrMsg("invalid int literal at line [%d]", lineNo + 1)
						}
						t.v = n
					case "name":
						if isReserved(s) {
							t.tp = "reserved"
						}
						t.v = s
					default:
						panic("bug")
					}
					tl.add(t)
					break
				}
			}
			if !ok {
				exitWithErrMsg("invalid token at line [%d]", lineNo + 1)
			}
		}
	}

	return
}
