#include "textflag.h"

TEXT ·libc_dlopen_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dlopen(SB)
TEXT ·libc_dlopen_preflight_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dlopen_preflight(SB)
TEXT ·libc_dlerror_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dlerror(SB)
TEXT ·libc_dlclose_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dlclose(SB)
TEXT ·libc_dlsym_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dlsym(SB)
TEXT ·libc_dladdr_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_dladdr(SB)
