package main

import (
	"fmt"
)

type objIntf interface {
	asBool() bool
	asStr() string
	getType() string

	unaryOp(op string) objIntf
	binocularOp(op string, other objIntf) objIntf
	getAttr(attrName string) objIntf
	call(argList []objIntf) objIntf
	setItem(key objIntf, value objIntf)
}

type ObjBaseType struct {
	typeName string
}

func newObjBase(typeName string) ObjBaseType {
	return ObjBaseType{
		typeName: typeName,
	}
}

func (o ObjBaseType) asBool() bool {
	return true
}

func (o ObjBaseType) asStr() string {
	return fmt.Sprintf("<%s obj>", o.typeName)
}

func (o ObjBaseType) getType() string {
	return o.typeName
}

func (o ObjBaseType) unaryOp(op string) objIntf {
	exitWithErrMsg("unsupported op '%s' for type '%s'", op, o.getType())
	panic("bug")
}

func (o ObjBaseType) binocularOp(op string, other objIntf) objIntf {
	exitWithErrMsg("unsupported op '%s' for type '%s' and '%s'", op, o.getType(), other.getType())
	panic("bug")
}

func (o ObjBaseType) getAttr(attrName string) objIntf {
	exitWithErrMsg("type '%s' has no attr '%s'", o.getType(), attrName)
	panic("bug")
}

func (o ObjBaseType) call(argList []objIntf) objIntf {
	exitWithErrMsg("unsupported call op for type '%s'", o.getType())
	panic("bug")
}

func (o ObjBaseType) setItem(key objIntf, value objIntf) {
	exitWithErrMsg("unsupported setItem op for type '%s'", o.getType())
	panic("bug")
}
