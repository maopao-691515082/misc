package main

var (

	unaryOpConvMap = map[string]string{
		"!": "!",
		"-": "neg",
		"+": "pos",
	}

	unaryOpSet = listToSet([]string{"~", "!", "neg", "pos"})

	binocularOpSet = listToSet([]string{"!=", "==", "&&", "||", "%", "*", "-", "+", "<", ">", "/"})

	opPriorityList = [][]string{
		{"||"},
		{"&&"},
		{"==", "!="},
		{"<", ">"},
		{"+", "-"},
		{"*", "/", "%"},
		{"!", "neg", "pos"},
	}
	opPriorityMap map[string]int

)

func init() {
	opPriorityMap = make(map[string]int)
	for i, opList := range opPriorityList {
		for _, op := range opList {
			opPriorityMap[op] = i
		}
	}
}

type exprIntf interface {
	eval() objIntf
}

type lvalueIntf interface {
	set(e exprIntf)
}

type parseStkType struct {
	startToken tokenType
	exprStk    []exprIntf
	opStk      []string
}

func (stk *parseStkType) pushOp(op string) {
	for len(stk.opStk) > 0 {
		if opPriorityMap[stk.opStk[len(stk.opStk) - 1]] > opPriorityMap[op] {
			stk.popTopOp()
		} else if opPriorityMap[stk.opStk[len(stk.opStk) - 1]] < opPriorityMap[op] {
			break
		} else {
			_, ok := unaryOpSet[op]
			if ok {
				break
			}
			stk.popTopOp()
		}
	}
	stk.opStk = append(stk.opStk, op)
}

func (stk *parseStkType) popTopOp() {
	op := stk.opStk[len(stk.opStk) - 1]
	stk.opStk = stk.opStk[: len(stk.opStk) - 1]
	_, ok := unaryOpSet[op]
	if ok {
		if len(stk.exprStk) < 1 {
			stk.startToken.raise("invalid expr")
		}
		e := stk.popExpr()
		stk.pushExpr(opExpr{
			op:   op,
			args: []exprIntf{e},
		})
		return
	}
	_, ok = binocularOpSet[op]
	if ok {
		if len(stk.exprStk) < 2 {
			stk.startToken.raise("invalid expr")
		}
		eb := stk.popExpr()
		ea := stk.popExpr()
		stk.pushExpr(opExpr{
			op:   op,
			args: []exprIntf{ea, eb},
		})
		return
	}
	panic("bug")
}

func (stk *parseStkType) pushExpr(e exprIntf) {
	stk.exprStk = append(stk.exprStk, e)
}

func (stk *parseStkType) popExpr() exprIntf {
	e := stk.exprStk[len(stk.exprStk) - 1]
	stk.exprStk = stk.exprStk[: len(stk.exprStk) - 1]
	return e
}

func (stk *parseStkType) finish() exprIntf {
	for len(stk.opStk) > 0 {
		stk.popTopOp()
	}
	if len(stk.exprStk) != 1 {
		stk.startToken.raise("invalid expr")
	}
	return stk.exprStk[0]
}

func parseExpr(tl *tokenListType) exprIntf {
	stk := &parseStkType{
		startToken: tl.peek(),
		exprStk:    []exprIntf{},
		opStk:      []string{},
	}
	for {
		t := tl.pop()

		if t.isSym(nil) {
			sym, ok := unaryOpConvMap[t.v.(string)]
			if ok {
				stk.pushOp(sym)
				continue
			}
		}

		if t.isSym("(") {
			stk.pushExpr(parseExpr(tl))
			tl.popSym(")")
		} else if t.isName() {
			stk.pushExpr(loadVarExpr{
				name: t.v.(string),
			})
		} else if t.isInt() {
			stk.pushExpr(intLiteralExpr{
				v: t.v.(int64),
			})
		} else if t.isStr() {
			stk.pushExpr(strLiteralExpr{
				v: t.v.(string),
			})
		} else if t.isSym("[") {
			el := parseExprList(tl)
			tl.popSym("]")
			stk.pushExpr(buildListExpr{
				el: el,
			})
		} else {
			t.raise("invalid expr")
		}

		for tl.len() > 0 {
			t = tl.pop()
			if t.isSym("[") {
				eb := parseExpr(tl)
				tl.popSym("]")
				ea := stk.popExpr()
				stk.pushExpr(opExpr{
					op:   "[]",
					args: []exprIntf{ea, eb},
				})
			} else if t.isSym("(") {
				el := parseExprList(tl)
				tl.popSym(")")
				e := stk.popExpr()
				stk.pushExpr(opExpr{
					op:   "()",
					args: append([]exprIntf{e}, el...),
				})
			} else if t.isSym(".") {
				attrName := tl.popName()
				e := stk.popExpr()
				stk.pushExpr(getAttrExpr{
					e:        e,
					attrName: attrName,
				})
			} else {
				tl.revert()
				break
			}
		}

		if isExprEnd(tl.peek()) {
			break
		}

		t = tl.pop()
		if t.isSym(nil) {
			sym := t.v.(string)
			_, ok := binocularOpSet[sym]
			if !ok {
				t.raise("expect binocular operator")
			}
			stk.pushOp(sym)
		} else {
			t.raise("expect binocular operator")
		}
	}

	return stk.finish()
}

func isExprEnd(t tokenType) (ok bool) {
	if !t.isSym(nil) {
		return
	}
	switch t.v.(string) {
	case ")", "]", ",", ";", "=":
		ok = true
	}
	return
}

func parseExprList(tl *tokenListType) []exprIntf {
	el := []exprIntf{}
	for {
		t := tl.peek()
		if t.isSym(")") || t.isSym("]") {
			return el
		}

		el = append(el, parseExpr(tl))

		t = tl.peek()
		if t.isSym(",") {
			tl.popSym(",")
		} else {
			return el
		}
	}
}
