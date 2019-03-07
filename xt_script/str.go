package main

type strObjType struct {
	ObjBaseType

	v string
}

func newStr(v string) strObjType {
	return strObjType{
		ObjBaseType: newObjBase("str"),

		v: v,
	}
}

func (this strObjType) asBool() bool {
	return this.v != ""
}

func (this strObjType) asStr() string {
	return this.v
}

func (this strObjType) binocularOp(op string, otherObj objIntf) objIntf {
	switch other := otherObj.(type) {
	case strObjType:
		switch op {
		case "==", "!=", "<", ">":
			var r bool
			switch op {
			case "==":
				r = this.v == other.v
			case "!=":
				r = this.v != other.v
			case "<":
				r = this.v < other.v
			case ">":
				r = this.v > other.v
			default:
				panic("bug")
			}
			if r {
				return newInt(1)
			} else {
				return newInt(0)
			}
		case "+":
			return newStr(this.v + other.v)
		}
	}
	return this.ObjBaseType.binocularOp(op, otherObj)
}
