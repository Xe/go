// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package testenv

import (
	"os"
	"runtime"
)

func hasSymlink() (ok bool, reason string) {
	switch runtime.GOOS {
	case "android", "plan9":
		return false, ""
	case "wasip1":
		if os.Getenv("RUNTIME") == "wasmtime" {
			return false, ""
		}
	}
	return true, ""
}
