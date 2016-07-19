package main

import (
    "time"
    "fmt"
    "runtime"
)

const LOOPS = 50000000

type Record struct {
    PtrComp *Record
    Discr int32
    EnumComp int32
    IntComp int32
    StringComp string
}

func copyRecord(nr *Record, r *Record) {
    nr.PtrComp = r.PtrComp
    nr.Discr = r.Discr
    nr.EnumComp = r.EnumComp
    nr.IntComp = r.IntComp
    nr.StringComp = r.StringComp
}

const (
    Ident1 = 1
    Ident2 = 2
    Ident3 = 3
    Ident4 = 4
    Ident5 = 5
)

type G struct {
    IntGlob int32
    BoolGlob bool
    Char1Glob uint8
    Char2Glob uint8
    Array1Glob *[51]int32
    Array2Glob *[51][51]int32
    PtrGlb *Record
    PtrGlbNext *Record
}

func NewG() *G {
    g := new(G)
    g.Array1Glob = new([51]int32)
    g.Array2Glob = new([51][51]int32)
    g.PtrGlb = new(Record)
    g.PtrGlbNext = new(Record)
    return g
}

func Func3(EnumParIn int32) bool {
    EnumLoc := EnumParIn
    if EnumLoc == Ident3 {
        return true
    }
    return false
}

func Func2(StrParI1 string, StrParI2 string) bool {
    var IntLoc int32
    IntLoc = 1
    var CharLoc uint8
    CharLoc = 0
    for IntLoc <= 1 {
        if Func1(StrParI1[IntLoc], StrParI2[IntLoc + 1]) == Ident1 {
            CharLoc = 'A'
            IntLoc ++
        }
    }
    if CharLoc >= 'W' && CharLoc <= 'Z' {
        IntLoc = 7
    }
    if CharLoc == 'X' {
        return true;
    } else {
        if StrParI1 > StrParI2 {
            IntLoc += 7
            return true
        } else {
            return false
        }
    }
}

func Func1(CharPar1 uint8, CharPar2 uint8) int32 {
    CharLoc1 := CharPar1
    CharLoc2 := CharLoc1
    if CharLoc2 != CharPar2 {
        return Ident1
    } else {
        return Ident2
    }
}

func Proc8(g *G, Array1Par *[51]int32, Array2Par *[51][51]int32, IntParI1 int32, IntParI2 int32) {
    var IntLoc int32
    IntLoc = IntParI1 + 5
    Array1Par[IntLoc] = IntParI2
    Array1Par[IntLoc+1] = Array1Par[IntLoc]
    Array1Par[IntLoc+30] = IntLoc
    for IntIndex := IntLoc; IntIndex <= IntLoc + 1; IntIndex ++ {
        Array2Par[IntLoc][IntIndex] = IntLoc
    }
    Array2Par[IntLoc][IntLoc - 1] ++
    Array2Par[IntLoc + 20][IntLoc] = Array1Par[IntLoc]
    g.IntGlob = 5
}

func Proc7(IntParI1 int32, IntParI2 int32) int32 {
    IntLoc := IntParI1 + 2
    IntParOut := IntParI2 + IntLoc
    return IntParOut
}

func Proc6(g *G, EnumParIn int32) int32 {
    var EnumParOut int32
    EnumParOut = EnumParIn
    if !Func3(EnumParIn) {
        EnumParOut = Ident4
    }
    if !(EnumParIn == Ident1) {
        EnumParOut = Ident1
    } else if EnumParIn == Ident2 {
        if g.IntGlob > 100 {
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

func Proc5(g *G) {
    g.Char1Glob = 'A'
    g.BoolGlob = false
}

func Proc4(g *G) {
    BoolLoc := g.Char1Glob == 'A'
    BoolLoc = BoolLoc || g.BoolGlob
    g.Char2Glob = 'B'
}

func Proc3(g *G, PtrParOut *Record) *Record {
    if g.PtrGlb != nil {
        PtrParOut = g.PtrGlb.PtrComp
    } else {
        g.IntGlob = 100
    }
    g.PtrGlb.IntComp = Proc7(10, g.IntGlob)
    return PtrParOut
}

func Proc2(g *G, IntParIO int32) int32 {
    IntLoc := IntParIO + 10
    var EnumLoc int32
    EnumLoc = 0
    for true {
        if g.Char1Glob == 'A' {
            IntLoc --
            IntParIO = IntLoc - g.IntGlob
            EnumLoc = Ident1
        }
        if EnumLoc == Ident1 {
            break
        }
    }
    return IntParIO
}

func Proc1(g *G, PtrParIn *Record) {
    NextRecord := PtrParIn.PtrComp
    copyRecord(NextRecord, g.PtrGlb)
    PtrParIn.IntComp = 5
    NextRecord.IntComp = PtrParIn.IntComp
    NextRecord.PtrComp = PtrParIn.PtrComp
    NextRecord.PtrComp = Proc3(g, NextRecord.PtrComp)
    if NextRecord.Discr == Ident1 {
        NextRecord.IntComp = 6
        NextRecord.EnumComp = Proc6(g, PtrParIn.EnumComp)
        NextRecord.PtrComp = g.PtrGlb.PtrComp
        NextRecord.IntComp = Proc7(NextRecord.IntComp, 10)
    } else {
        copyRecord(PtrParIn, NextRecord)
    }
}

func Proc0(g *G) {
    g.PtrGlb.PtrComp = g.PtrGlbNext
    g.PtrGlb.Discr = Ident1
    g.PtrGlb.EnumComp = Ident3
    g.PtrGlb.IntComp = 40
    g.PtrGlb.StringComp = "DHRYSTONE PROGRAM, SOME STRING"
    String1Loc := "DHRYSTONE PROGRAM, 1'ST STRING"
    g.Array2Glob[8][7] = 10

    for i := 0; i < LOOPS; i ++ {
        Proc5(g)
        Proc4(g)
        var IntLoc1 int32
        IntLoc1 = 2
        var IntLoc2 int32
        IntLoc2 = 3
        String2Loc := "DHRYSTONE PROGRAM, 2'ND STRING"
        EnumLoc := int32(Ident2)
        g.BoolGlob = !Func2(String1Loc, String2Loc)
        var IntLoc3 int32
        IntLoc3 = 0
        for IntLoc1 < IntLoc2 {
            IntLoc3 = 5 * IntLoc1 - IntLoc2;
            IntLoc3 = Proc7(IntLoc1, IntLoc2);
            IntLoc1 = IntLoc1 + 1;
        }
        Proc8(g, g.Array1Glob, g.Array2Glob, IntLoc1, IntLoc3);
        Proc1(g, g.PtrGlb);
        var CharIndex uint8;
        CharIndex = 'A';
        for CharIndex <= g.Char2Glob {
            if EnumLoc == Func1(CharIndex, 'C') {
                EnumLoc = Proc6(g, Ident1);
            }
            CharIndex ++;
        }
        IntLoc3 = IntLoc2 * IntLoc1;
        IntLoc2 = IntLoc3 / IntLoc1;
        IntLoc2 = 7 * (IntLoc3 - IntLoc2) - IntLoc1;
        IntLoc1 = Proc2(g, IntLoc1);
    }
}

var finish_chan chan int = make(chan int, 10)
var core_count int = 8

func main() {
    runtime.GOMAXPROCS(core_count)

    for i := 0; i < core_count; i ++ {
        go func (idx int) {
            ts := time.Now()
            Proc0(NewG())
            tm := time.Now().Sub(ts)
            float_tm := float64(tm) / 1e9
            fmt.Printf("[%d]Time Used %f\n", idx, float_tm)
            fmt.Printf("[%d]This machine benchmarks at %f GoStones/second\n", idx, LOOPS / float_tm);
            finish_chan <- 0
        }(i)
    }
    for i := 0; i < core_count; i ++ {
        <- finish_chan
    }
}
