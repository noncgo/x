//go:build darwin
// +build darwin

#include "go_asm.h"
#include "textflag.h"
#include "../../src/runtime/cgo/abi_amd64.h"

GLOBL ·crosscallCallbackABI0(SB), NOPTR|RODATA, $8
DATA ·crosscallCallbackABI0(SB)/8, $·crosscallCallback(SB)

// crosscallCallback is the implementation of C FSEventStreamCallback
// that calls back into Go function passed through the user info.
//
// It is called from C on a system stack using System V calling convention.
//
// Registers:
//  DI stream
//  SI userInfo
//  DX numEvents
//  CX eventPaths
//  R8 eventFlags
//  R9 eventIDs
//
TEXT ·crosscallCallback(SB), NOSPLIT, $0-0
	// Transition from C ABI to Go ABI.
	PUSH_REGS_HOST_TO_ABI0()

	ADJSP $callbackArgs__size
	MOVQ  DI, callbackArgs_stream(SP)
	MOVQ  SI, callbackArgs_ctxt(SP)
	MOVQ  DX, callbackArgs_numEvents(SP)
	MOVQ  CX, callbackArgs_eventPaths(SP)
	MOVQ  R8, callbackArgs_eventFlags(SP)
	MOVQ  R9, callbackArgs_eventIDs(SP)
	LEAQ  (SP), SI

	ADJSP $3*8
	MOVQ  ·callbackABIInternal(SB), AX
	MOVQ  AX, (0*8)(SP)
	MOVQ  SI, (1*8)(SP)
	MOVQ  $0, (2*8)(SP)
	CALL  runtime·cgocallback(SB)

	ADJSP $-callbackArgs__size
	ADJSP $-3*8
	POP_REGS_HOST_TO_ABI0()
	RET
