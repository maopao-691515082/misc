package main

import (
    "time"
    "fmt"
)

const LOOPS = 5000000

type Record struct {
    PtrComp *Record
    Discr interface{}
    EnumComp interface{}
    IntComp interface{}
    StringComp string
}

func copyRecord(nr *Record, r *Record) {
    nr.PtrComp = r.PtrComp
    nr.Discr = r.Discr
    nr.EnumComp = r.EnumComp
    nr.IntComp = r.IntComp
    nr.StringComp = r.StringComp
}

var Ident1 interface{}
var Ident2 interface{}
var Ident3 interface{}
var Ident4 interface{}
var Ident5 interface{}

var IntGlob interface{}
var BoolGlob interface{}
var Char1Glob interface{}
var Char2Glob interface{}
var Array1Glob0 [51]interface{}
var Array2Glob0 [51][51]interface{}
var Array1Glob []interface{} = Array1Glob0[:]
var Array2Glob [][51]interface{} = Array2Glob0[:]

var PtrGlb *Record
var PtrGlbNext *Record

func Func3(EnumParIn interface{}) bool {
    EnumLoc := EnumParIn
    if EnumLoc == Ident3 {
        return true
    }
    return false
}

func Func2(StrParI1 string, StrParI2 string) bool {
    var IntLoc interface{}
    IntLoc = 1
    var CharLoc interface{}
    CharLoc = '\x00'
    for IntLoc.(int) <= 1 {
        if Func1(StrParI1[IntLoc.(int)], StrParI2[IntLoc.(int) + 1]) == Ident1 {
            CharLoc = 'A'
            IntLoc = IntLoc.(int) + 1
        }
    }
    if CharLoc.(int32) >= 'W' && CharLoc.(int32) <= 'Z' {
        IntLoc = 7
    }
    if CharLoc.(int32) == 'X' {
        return true;
    } else {
        if StrParI1 > StrParI2 {
            IntLoc = IntLoc.(int) + 7
            return true
        } else {
            return false
        }
    }
}

func Func1(CharPar1 interface{}, CharPar2 interface{}) interface{} {
    CharLoc1 := CharPar1
    CharLoc2 := CharLoc1
    if CharLoc2 != CharPar2 {
        return Ident1
    } else {
        return Ident2
    }
}

func Proc8(Array1Par []interface{}, Array2Par [][51]interface{}, IntParI1 interface{}, IntParI2 interface{}) {
    var IntLoc interface{}
    IntLoc = IntParI1.(int) + 5
    Array1Par[IntLoc.(int)] = IntParI2
    Array1Par[IntLoc.(int) + 1] = Array1Par[IntLoc.(int)]
    Array1Par[IntLoc.(int) + 30] = IntLoc
    for IntIndex := IntLoc; IntIndex.(int) <= IntLoc.(int) + 1; IntIndex = IntIndex.(int) + 1 {
        Array2Par[IntLoc.(int)][IntIndex.(int)] = IntLoc
    }
    Array2Par[IntLoc.(int)][IntLoc.(int) - 1] = Array2Par[IntLoc.(int)][IntLoc.(int) - 1].(int) + 1
    Array2Par[IntLoc.(int) + 20][IntLoc.(int)] = Array1Par[IntLoc.(int)]
    IntGlob = 5
}

func Proc7(IntParI1 interface{}, IntParI2 interface{}) interface{} {
    var IntLoc interface{} = IntParI1.(int) + 2
    var IntParOut interface{} = IntParI2.(int) + IntLoc.(int)
    return IntParOut
}

func Proc6(EnumParIn interface{}) interface{} {
    var EnumParOut interface{}
    EnumParOut = EnumParIn
    if !Func3(EnumParIn) {
        EnumParOut = Ident4
    }
    if !(EnumParIn == Ident1) {
        EnumParOut = Ident1
    } else if EnumParIn == Ident2 {
        if IntGlob.(int) > 100 {
            EnumParOut = Ident1
        } else {
            EnumParOut = Ident4
        }
    } else if EnumParIn == Ident3 {
        EnumParOut = Ident2
    } else if EnumParIn == Ident4 {
    } else if EnumParIn == Ident5 {
        EnumParOut = Ident3
    }
    return EnumParOut
}

func Proc5() {
    Char1Glob = 'A'
    BoolGlob = false
}

func Proc4() {
    var BoolLoc interface{} = Char1Glob == 'A'
    BoolLoc = BoolLoc.(bool) || BoolGlob.(bool)
    Char2Glob = 'B'
}

func Proc3(PtrParOut *Record) *Record {
    if PtrGlb != nil {
        PtrParOut = PtrGlb.PtrComp
    } else {
        IntGlob = 100
    }
    PtrGlb.IntComp = Proc7(10, IntGlob)
    return PtrParOut
}

func Proc2(IntParIO interface{}) interface{} {
    var IntLoc interface{} = IntParIO.(int) + 10
    var EnumLoc interface{}
    EnumLoc = 0
    for true {
        if Char1Glob == 'A' {
            IntLoc = IntLoc.(int) - 1
            IntParIO = IntLoc.(int) - IntGlob.(int)
            EnumLoc = Ident1
        }
        if EnumLoc == Ident1 {
            break
        }
    }
    return IntParIO
}

func Proc1(PtrParIn *Record) {
    NextRecord := PtrParIn.PtrComp
    copyRecord(NextRecord, PtrGlb)
    PtrParIn.IntComp = 5
    NextRecord.IntComp = PtrParIn.IntComp
    NextRecord.PtrComp = PtrParIn.PtrComp
    NextRecord.PtrComp = Proc3(NextRecord.PtrComp)
    if NextRecord.Discr == Ident1 {
        NextRecord.IntComp = 6
        NextRecord.EnumComp = Proc6(PtrParIn.EnumComp)
        NextRecord.PtrComp = PtrGlb.PtrComp
        NextRecord.IntComp = Proc7(NextRecord.IntComp, 10)
    } else {
        copyRecord(PtrParIn, NextRecord)
    }
}

func Proc0() {
    PtrGlbNext = new(Record)
    PtrGlb = new(Record)
    PtrGlb.PtrComp = PtrGlbNext
    PtrGlb.Discr = Ident1
    PtrGlb.EnumComp = Ident3
    PtrGlb.IntComp = 40
    PtrGlb.StringComp = "DHRYSTONE PROGRAM, SOME STRING"
    String1Loc := "DHRYSTONE PROGRAM, 1'ST STRING"
    Array2Glob[8][7] = 10

    for i := 0; i < LOOPS; i ++ {
        Proc5()
        Proc4()
        var IntLoc1 interface{}
        IntLoc1 = 2
        var IntLoc2 interface{}
        IntLoc2 = 3
        String2Loc := "DHRYSTONE PROGRAM, 2'ND STRING"
        EnumLoc := Ident2
        BoolGlob = !Func2(String1Loc, String2Loc)
        var IntLoc3 interface{}
        IntLoc3 = 0
        for IntLoc1.(int) < IntLoc2.(int) {
            IntLoc3 = 5 * IntLoc1.(int) - IntLoc2.(int);
            IntLoc3 = Proc7(IntLoc1, IntLoc2);
            IntLoc1 = IntLoc1.(int) + 1;
        }
        Proc8(Array1Glob, Array2Glob, IntLoc1, IntLoc3);
        Proc1(PtrGlb);
        var CharIndex interface{};
        CharIndex = 'A';
        for CharIndex.(int32) <= Char2Glob.(int32) {
            if EnumLoc == Func1(CharIndex, 'C') {
                EnumLoc = Proc6(Ident1);
            }
            CharIndex = CharIndex.(int32) + 1;
        }
        IntLoc3 = IntLoc2.(int) * IntLoc1.(int);
        IntLoc2 = IntLoc3.(int) / IntLoc1.(int);
        IntLoc2 = 7 * (IntLoc3.(int) - IntLoc2.(int)) - IntLoc1.(int);
        IntLoc1 = Proc2(IntLoc1);
    }
}

func main() {
    Ident1 = 1
    Ident2 = 2
    Ident3 = 3
    Ident4 = 4
    Ident5 = 5
    
    IntGlob = 0
    BoolGlob = false
    Char1Glob = 0
    Char2Glob = 0
    for i := 0; i < 51; i ++ {
        Array1Glob[i] = 0
        for j := 0; j < 51; j ++ {
            Array2Glob[i][j] = 0
        }
    }
    PtrGlb = nil
    PtrGlbNext = nil

    ts := time.Now()
    Proc0()
    tm := time.Now().Sub(ts)
    float_tm := float64(tm) / 1e9
    fmt.Printf("Time Used %f\n", float_tm)
    fmt.Printf("This machine benchmarks at %f GoStones/second\n", LOOPS / float_tm);
}
