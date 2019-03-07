package main

import (
	"fmt"
)

type intObjType struct {
	ObjBaseType

	v int64
}

func newInt(v int64) intObjType {
	return intObjType{
		ObjBaseType: newObjBase("int"),

		v: v,
	}
}

func (this intObjType) asBool() bool {
	return this.v != 0
}

func (this intObjType) asStr() string {
	return fmt.Sprintf("%d", this.v)
}

func (this intObjType) unaryOp(op string) objIntf {
	switch op {
	case "!":
		if this.asBool() {
			return newInt(0)
		} else {
			return newInt(1)
		}
	case "neg":
		return newInt(-this.v)
	case "pos":
		return this
	default:
	}
	return this.ObjBaseType.unaryOp(op)
}

func (this intObjType) binocularOp(op string, otherObj objIntf) objIntf {
	switch other := otherObj.(type) {
	case intObjType:
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
		case "%":
			return newInt(this.v % other.v)
		case "*":
			return newInt(this.v * other.v)
		case "-":
			return newInt(this.v - other.v)
		case "+":
			return newInt(this.v + other.v)
		case "/":
			return newInt(this.v / other.v)
		}
	}
	return this.ObjBaseType.binocularOp(op, otherObj)
}
