package main

type blockStmt struct {
	sl *stmtListType
}

func (s blockStmt) exec() int {
	return s.sl.exec()
}

type loopStmt struct {
	sl *stmtListType
}

func (s loopStmt) exec() int {
	for {
		r := s.sl.exec()
		if r == execResult_BREAK {
			return execResult_NORMAL
		}
	}
}

type bocStmt struct {
	boc string
}

func (s bocStmt) exec() int {
	switch s.boc {
	case "break":
		return execResult_BREAK
	case "continue":
		return execResult_CONTINUE
	}
	panic("bug")
}

type ifStmtIfPartType struct {
	e  exprIntf
	sl *stmtListType
}

type ifStmt struct {
	ifPartList []ifStmtIfPartType
	elsePart   *stmtListType
}

func (s ifStmt) exec() int {
	for _, ifPart := range s.ifPartList {
		if ifPart.e.eval().asBool() {
			return ifPart.sl.exec()
		}
	}
	if s.elsePart != nil {
		return s.elsePart.exec()
	}
	return execResult_NORMAL
}

type assignStmt struct {
	lv lvalueIntf
	e  exprIntf
}

func (s assignStmt) exec() int {
	s.lv.set(s.e)
	return execResult_NORMAL
}

type exprStmt struct {
	e exprIntf
}

func (s exprStmt) exec() int {
	s.e.eval()
	return execResult_NORMAL
}
