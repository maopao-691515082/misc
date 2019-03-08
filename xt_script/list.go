package main

import (
	"strings"
)

type listObjType struct {
	ObjBaseType

	ol []objIntf
}

func newList(objList []objIntf) *listObjType {
	return &listObjType{
		ObjBaseType: newObjBase("list"),

		ol: objList,
	}
}

func (this *listObjType) asBool() bool {
	return len(this.ol) != 0
}

func (this *listObjType) asStr() string {
	sl := make([]string, 0, len(this.ol) * 2 + 2)
	sl = append(sl, "[")
	for i, o := range this.ol {
		if i != 0 {
			sl = append(sl, ", ")
		}
		sl = append(sl, o.asStr())
	}
	sl = append(sl, "]")
	return strings.Join(sl, "")
}

func (this *listObjType) binocularOp(op string, otherObj objIntf) objIntf {
	switch other := otherObj.(type) {
	case intObjType:
		switch op {
		case "[]":
			idx := other.v
			if idx < 0 {
				idx += int64(len(this.ol))
			}
			if idx < 0 || idx >= int64(len(this.ol)) {
				exitWithErrMsg("list index out of range")
			}
			return this.ol[idx]
		}
	}
	return this.ObjBaseType.binocularOp(op, otherObj)
}

func (this *listObjType) getAttr(attrName string) objIntf {
	switch attrName {
	case "len":
		return newFunc(func (al []objIntf) objIntf {
			if len(al) > 0 {
				exitWithErrMsg("list.len need 0 args")
			}
			return newInt(int64(len(this.ol)))
		})
	case "add":
		return newFunc(func (al []objIntf) objIntf {
			this.ol = append(this.ol, al...)
			return this
		})
	case "pop":
		return newFunc(func (al []objIntf) objIntf {
			if len(al) > 0 {
				exitWithErrMsg("list.pop need 0 args")
			}
			if len(this.ol) == 0 {
				exitWithErrMsg("list.pop on empty list")
			}
			o := this.ol[len(this.ol) - 1]
			this.ol = this.ol[: len(this.ol) - 1]
			return o
		})
	}
	return this.ObjBaseType.getAttr(attrName)
}

func (this *listObjType) setItem(keyObj objIntf, value objIntf) {
	switch key := keyObj.(type) {
	case intObjType:
		this.ol[key.v] = value
	default:
		exitWithErrMsg("list index must be int")
	}
}
