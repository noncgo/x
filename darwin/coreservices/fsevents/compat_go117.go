//go:build darwin && !go1.18
// +build darwin,!go1.18

package fsevents

// any is an alias for interface{} and is equivalent to interface{} in all
// ways.
type any = interface{}
