#include "go_asm.h"
#include "textflag.h"

// void call(di *frame);

GLOBL ·callABI0(SB), NOPTR|RODATA, $8
DATA ·callABI0(SB)/8, $·call(SB)
TEXT ·call(SB), NOSPLIT, $0
	PUSHQ BP
	MOVQ  SP, BP
	SUBQ  $16, SP

	MOVQ DI, BX

	MOVQ  frame_DI(BX), DI
	MOVQ  frame_SI(BX), SI
	MOVQ  frame_DX(BX), DX
	MOVQ  frame_CX(BX), CX
	MOVQ  frame_R8(BX), R8
	MOVQ  frame_R9(BX), R9
	MOVQ  frame_AX(BX), AX
	MOVSD frame_X0(BX), X0
	MOVSD frame_X1(BX), X1
	MOVSD frame_X2(BX), X2
	MOVSD frame_X3(BX), X3
	MOVSD frame_X4(BX), X4
	MOVSD frame_X5(BX), X5
	MOVSD frame_X6(BX), X6
	MOVSD frame_X7(BX), X7

	CALL frame_FuncPC(BX)

	MOVQ AX, frame_AX(BX)
	MOVQ DX, frame_DX(BX)

	MOVQ BP, SP
	POPQ BP
	RET
