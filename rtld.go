package dyld

const (
	RTLD_LAZY   = 0x1
	RTLD_NOW    = 0x2
	RTLD_LOCAL  = 0x4
	RTLD_GLOBAL = 0x8

	RTLD_NOLOAD   = 0x10
	RTLD_NODELETE = 0x80
	RTLD_FIRST    = 0x100
)

// Special handle arguments for dlsym().
const (
	RTLD_NEXT      = ^uintptr(0) // Search subsequent objects.
	RTLD_DEFAULT   = ^uintptr(1) // Use default search algorithm.
	RTLD_SELF      = ^uintptr(2) // Search this and subsequent objects.
	RTLD_MAIN_ONLY = ^uintptr(4) // Search main executable only.

)
