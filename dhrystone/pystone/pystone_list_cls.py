import sys, time

LOOPS = int(sys.argv[1]) if len(sys.argv) == 2 else 500000

Ident1 = 1
Ident2 = 2
Ident3 = 3
Ident4 = 4
Ident5 = 5

IntGlob = 0
BoolGlob = False
Char1Glob = 0
Char2Glob = 0
Array1Glob = [0] * 51
Array2Glob = [[0] * 51 for i in xrange(51)]

PtrGlb = None
PtrGlbNext = None

def Func3(EnumParIn):
    EnumLoc = EnumParIn
    if EnumLoc == Ident3:
        return True
    return False

def Func2(StrParI1, StrParI2):
    IntLoc = 1
    CharLoc = 0
    while IntLoc <= 1:
        if Func1(StrParI1[IntLoc], StrParI2[IntLoc + 1]) == Ident1:
            CharLoc = 65
            IntLoc += 1
    if CharLoc >= 87 and CharLoc <= 90:
        IntLoc = 7
    if CharLoc == 88:
        return True
    else:
        if StrParI1 > StrParI2:
            IntLoc += 7
            return True
        else:
            return False

def Func1(CharPar1, CharPar2):
    CharLoc1 = CharPar1
    CharLoc2 = CharLoc1
    if CharLoc2 != CharPar2:
        return Ident1
    else:
        return Ident2

def Proc8(Array1Par, Array2Par, IntParI1, IntParI2):
    global IntGlob
    IntLoc = IntParI1 + 5
    Array1Par[IntLoc] = IntParI2
    Array1Par[IntLoc + 1] = Array1Par[IntLoc]
    Array1Par[IntLoc + 30] = IntLoc
    for IntIndex in (IntLoc, IntLoc + 1):
        Array2Par[IntLoc][IntIndex] = IntLoc
    Array2Par[IntLoc][IntLoc - 1] += 1
    Array2Par[IntLoc + 20][IntLoc] = Array1Par[IntLoc]
    IntGlob = 5

def Proc7(IntParI1, IntParI2):
    IntLoc = IntParI1 + 2
    IntParOut = IntParI2 + IntLoc
    return IntParOut

def Proc6(EnumParIn):
    EnumParOut = EnumParIn
    if not Func3(EnumParIn):
        EnumParOut = Ident4
    if not (EnumParIn == Ident1):
        EnumParOut = Ident1
    elif EnumParIn == Ident2:
        if IntGlob > 100:
            EnumParOut = Ident1
        else:
            EnumParOut = Ident4
    elif EnumParIn == Ident3:
        EnumParOut = Ident2
    elif EnumParIn == Ident4:
        pass
    elif EnumParIn == Ident5:
        EnumParOut = Ident3
    return EnumParOut

def Proc5():
    global BoolGlob, Char1Glob
    Char1Glob = 65
    BoolGlob = False

def Proc4():
    global Char2Glob
    BoolLoc = Char1Glob == 65
    BoolLoc = BoolLoc or BoolGlob
    Char2Glob = 66

def Proc3(PtrParOut):
    global IntGlob
    if PtrGlb is not None:
        PtrParOut = PtrGlb[0]
    else:
        IntGlob = 100
    PtrGlb[3] = Proc7(10, IntGlob)
    return PtrParOut

def Proc2(IntParIO):
    IntLoc = IntParIO + 10
    EnumLoc = 0
    while True:
        if Char1Glob == 65:
            IntLoc -= 1
            IntParIO = IntLoc - IntGlob
            EnumLoc = Ident1
        if EnumLoc == Ident1:
            break
    return IntParIO

def Proc1(PtrParIn):
    NextRecord = PtrParIn[0]
    NextRecord[:] = PtrGlb
    PtrParIn[3] = 5
    NextRecord[3] = PtrParIn[3]
    NextRecord[0] = PtrParIn[0]
    NextRecord[0] = Proc3(NextRecord[0])
    if NextRecord[1] == Ident1:
        NextRecord[3] = 6
        NextRecord[2] = Proc6(PtrParIn[2])
        NextRecord[0] = PtrGlb[0]
        NextRecord[3] = Proc7(NextRecord[3], 10)
    else:
        PtrParIn[:] = NextRecord
    return PtrParIn

def Proc0():
    global PtrGlbNext, PtrGlb, BoolGlob
    PtrGlbNext = [None, 0, 0, 0, None]
    PtrGlb = [None, 0, 0, 0, None]
    PtrGlb[0] = PtrGlbNext
    PtrGlb[1] = Ident1
    PtrGlb[2] = Ident3
    PtrGlb[3] = 40
    PtrGlb[4] = "DHRYSTONE PROGRAM, SOME STRING"
    String1Loc = "DHRYSTONE PROGRAM, 1'ST STRING"
    Array2Glob[8][7] = 10

    for i in xrange(LOOPS):
        Proc5()
        Proc4()
        IntLoc1 = 2
        IntLoc2 = 3
        String2Loc = "DHRYSTONE PROGRAM, 2'ND STRING"
        EnumLoc = Ident2
        BoolGlob = not Func2(String1Loc, String2Loc)
        IntLoc3 = 0
        while IntLoc1 < IntLoc2:
            IntLoc3 = 5 * IntLoc1 - IntLoc2
            IntLoc3 = Proc7(IntLoc1, IntLoc2)
            IntLoc1 += 1
        Proc8(Array1Glob, Array2Glob, IntLoc1, IntLoc3)
        PtrGlb = Proc1(PtrGlb)
        CharIndex = 65
        while CharIndex <= Char2Glob:
            if EnumLoc == Func1(CharIndex, 67):
                EnumLoc = Proc6(Ident1)
            CharIndex += 1
        IntLoc3 = IntLoc2 * IntLoc1
        IntLoc2 = IntLoc3 / IntLoc1
        IntLoc2 = 7 * (IntLoc3 - IntLoc2) - IntLoc1
        IntLoc1 = Proc2(IntLoc1)

def main():
    ts = time.time()
    Proc0()
    tm = time.time() - ts
    print("Time used: %s sec" % tm)
    print("This machine benchmarks at %s PyStones/second" % (LOOPS / tm))

if __name__ == "__main__":
    main()
