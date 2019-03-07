package main

type opExpr struct {
	op   string
	args []exprIntf
}

func (e opExpr) eval() objIntf {
	_, ok := unaryOpSet[e.op]
	if ok {
		return e.args[0].eval().unaryOp(e.op)
	}
	_, ok = binocularOpSet[e.op]
	if ok || e.op == "[]" {
		return e.args[0].eval().binocularOp(e.op, e.args[1].eval())
	}
	if e.op == "()" {
		al := make([]objIntf, 0, len(e.args))
		for _, ae := range e.args[1 :] {
			al = append(al, ae.eval())
		}
		return e.args[0].eval().call(al)
	}
	panic("bug")
}

type loadVarExpr struct {
	name string
}

func (e loadVarExpr) eval() objIntf {
	obj, ok := varMap[e.name]
	if !ok {
		exitWithErrMsg("undefined var '%s'", e.name)
	}
	return obj
}

func (e loadVarExpr) set(ve exprIntf) {
	varMap[e.name] = ve.eval()
}

type intLiteralExpr struct {
	v int64
}

func (e intLiteralExpr) eval() objIntf {
	return newInt(e.v)
}

type strLiteralExpr struct {
	v string
}

func (e strLiteralExpr) eval() objIntf {
	return newStr(e.v)
}

type buildListExpr struct {
	el []exprIntf
}

func (e buildListExpr) eval() objIntf {
	//todo
	return nil
}

type getAttrExpr struct {
	e        exprIntf
	attrName string
}

func (e getAttrExpr) eval() objIntf {
	//todo
	return nil
}
