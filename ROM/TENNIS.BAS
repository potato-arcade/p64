REM Tennis Game

REM Initialize game state, and store in system RAM banks
10 POKE 1, RND 64
20 POKE 2, 8 + RND 56
30 POKE 3, 1
40 POKE 4, 1

REM Bat locations
50 POKE 5, 8 + RND 56
60 POKE 6, 8 + RND 56

REM Bat Vectors
70 POKE 7, 0
71 POKE 8, 0

REM Scores
72 POKE 9, 0
73 POKE 10, 0
END

.INTR KEYDOWN
81 LET LDY = PEEK 7
82 LET RDY = PEEK 8
83 LET K = KEY
84 IF K = "D" THEN LDY = -1
85 IF K = "F" THEN LDY = 1
86 IF K = "J" THEN RDY = 1
87 IF K = "K" THEN RDY = -1
88 POKE 7, LDY
89 POKE 8, RDY
END

.INTR KEYUP
POKE 7, 0
POKE 8, 0
END

.INTR VSYNC
100 CLEAR
110 X = PEEK 1
120 Y = PEEK 2
130 DX = PEEK 3
140 DY = PEEK 4

REM Bats
150 LY = PEEK 5
160 RY = PEEK 6
170 LDY = PEEK 7
180 RDY = PEEK 8
190 LY = LY + LDY
200 RY = RY + RDY
210 IF LY < 9 THEN LY = 8
220 IF LY > 53 THEN LY = 53
230 IF RY < 9 THEN RY = 8
240 IF RY > 53 THEN RY = 53
250 POKE 5, LY
260 POKE 6, RY

REM Move the ball
300 X = X + DX
310 Y = Y + DY

REM Top and lower Line

320 FOR I = 0 TO 63
330  SET I,7,1
340  SET I,63,1
350 NEXT I

REM Net 
440 FOR I = 1 TO 28
450  SET 32,I*2 + 8,1
460 NEXT I

REM Left and Right Bat
500 FOR I = 1 TO 8
510   SET 0,LY + I,1
520   SET 63,RY + I,1
530 NEXT I

900 GOSUB 1000
910 GOSUB 1100
920 GOSUB 1200
930 GOSUB 1300
940 GOTO 2000

1000 IF X > 0 THEN RETURN
1010 X = 1
1020 DX = 1
1030 RETURN

1100 IF X < 64 THEN RETURN
1110 X = 63
1120 DX = -1
1130 RETURN

1200 IF Y > 8 THEN RETURN
1210 Y = 8
1220 DY = 1
1230 RETURN

1300 IF Y < 64 THEN RETURN
1310 Y = 62
1320 DY = -1 * DY
1330 RETURN

2000 PRIN	e.RegisterBuiltin("LINE", 4, func(env interface{}, args []object.Object) object.Object {
		x := int(args[0].(*object.NumberObject).Value)
		y := int(args[1].(*object.NumberObject).Value)
		x2 := int(args[2].(*object.NumberObject).Value)
		y2 := int(args[3].(*object.NumberObject).Value)
		p.frameBuffer.Line(x, y, x2, y2)
		return blankObj
	})T
2020 POKE 1, X
2030 POKE 2, Y
2040 POKE 3, DX
2050 POKE 4, DY
2060 SET X,Y,1
END
