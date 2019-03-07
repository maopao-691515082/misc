package main

const (
	execResult_NORMAL   = 0
	execResult_BREAK    = 1
	execResult_CONTINUE = 2
)

type stmtIntf interface {
	exec() int
}

type stmtListType struct {
	sl []stmtIntf
}

func (sl *stmtListType) add(s stmtIntf) {
	sl.sl = append(sl.sl, s)
}

func (sl *stmtListType) exec() int {
	for _, s := range sl.sl {
		r := s.exec()
		if r != execResult_NORMAL {
			return r
		}
	}
	return execResult_NORMAL
}

func parseStmtList(tokenList *tokenListType, loopDeep int) (stmtList *stmtListType) {
	stmtList = &stmtListType{}

	for tokenList.len() > 0 && !tokenList.peek().isSym("}") {
		t := tokenList.pop()

		if t.isSym(";") {
			continue
		}
		if t.isSym("{") {
			stmtList.add(blockStmt{
				sl: parseStmtList(tokenList, loopDeep),
			})
			tokenList.popSym("}")
			continue
		}
		if t.isReserved(nil) {
			r := t.v.(string)
			switch r {
			case "loop":
				tokenList.popSym("{")
				stmtList.add(loopStmt{
					sl: parseStmtList(tokenList, loopDeep + 1),
				})
				tokenList.popSym("}")
			case "break", "continue":
				if loopDeep == 0 {
					t.raise("'%s' outside of loop", r)
				}
				stmtList.add(bocStmt{
					boc: r,
				})
				tokenList.popSym(";")
			case "if":
				var (
					ifStmtIfPartList []ifStmtIfPartType
					ifStmtElsePart   *stmtListType
				)
				for {
					tokenList.popSym("(")
					e := parseExpr(tokenList)
					tokenList.popSym(")")
					tokenList.popSym("{")
					sl := parseStmtList(tokenList, loopDeep)
					tokenList.popSym("}")
					ifStmtIfPartList = append(ifStmtIfPartList, ifStmtIfPartType{
						e:  e,
						sl: sl,
					})
					if !tokenList.peek().isReserved("else") {
						break
					}
					tokenList.pop()
					t = tokenList.peek()
					if t.isReserved("if") {
						tokenList.pop()
						continue
					}
					tokenList.popSym("{")
					ifStmtElsePart = parseStmtList(tokenList, loopDeep)
					tokenList.popSym("}")
					break
				}
				stmtList.add(ifStmt{
					ifPartList: ifStmtIfPartList,
					elsePart:   ifStmtElsePart,
				})
			case "else":
				t.raise("single else")
			default:
				panic("bug")
			}
			continue
		}

		tokenList.revert()
		e := parseExpr(tokenList)
		if tokenList.peek().isSym("=") {
			tokenList.pop()
			lv, ok := e.(lvalueIntf)
			if !ok {
				t.raise("assign to non-lvalue")
			}
			e = parseExpr(tokenList)
			tokenList.popSym(";")
			stmtList.add(assignStmt{
				lv: lv,
				e:  e,
			})
		} else {
			tokenList.popSym(";")
			stmtList.add(exprStmt{
				e: e,
			})
		}
	}

	return
}
