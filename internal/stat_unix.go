//go:build !windows && !js
// +build !windows,!js

package internal

import (
	"golang.org/x/sys/unix"
)

// Stat_t is the Linux Stat_t.
type Stat_t = unix.Stat_t
