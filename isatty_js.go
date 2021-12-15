//go:build js
// +build js

package isatty

import (
	"golang.org/x/term"
)

// IsTerminal returns true if the file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
	return term.IsTerminal(int(fd))
}

// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}
