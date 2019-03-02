REM Tennis Game

REM Initialize game state, and store in system RAM banks
LET BL = 36
LET BR = 36
PRINT "Bats at locations", BL, BR, "\n"
LET BX = 32
LET BY = 36
PRINT "Ball at", BX, BY,"\n"

PRINT "Store to RAM", BL, BR, BX, BY, "\n"
POKE 100, BL
POKE 101, BR
POKE 102, BX
POKE 103, BY

PRINT "Load from RAM", "\n"
LET AL = PEEK 100
LET AR = PEEK 101
LET AX = PEEK 102
LET AY = PEEK 103
LET ABATS = PEEK 200
PRINT "Got", AL, AR, AX, AY, "\n"
PRINT "Bats", ABATS, "\n"
DEBUG

DIM a(2,2)
LET a[1,1] = "Muh String"
POKE 400, a
LET NEWA = PEEK 400
PRINT "New thing is\n", NEWA[1,1], "\n"
END

.INTR KEYDOWN
PRINT "KEY DOWN", KEY
END

.INTR KEYUP
PRINT "KEY UP", KEY
END

.INTR VSYNC
PRINT "VSYNC", KEY
END
