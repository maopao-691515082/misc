package main

type funcObjType struct {
	ObjBaseType

	f func (argList []objIntf) objIntf
}

func newFunc(f func (argList []objIntf) objIntf) funcObjType {
	return funcObjType{
		ObjBaseType: newObjBase("func"),

		f: f,
	}
}

func (fo funcObjType) call(argList []objIntf) objIntf {
	return fo.f(argList)
}
