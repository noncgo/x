//go:build darwin
// +build darwin

package corefoundation

// This file provides CFDicionary APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfdictionary-rum

import (
	"github.com/noncgo/x/darwin/internal/types"
)

// Dictionary is an opaque reference to a CFDictionary type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfdictionary
type Dictionary types.CFDictionary
