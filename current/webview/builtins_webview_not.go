//go:build !b_webview
// +build !b_webview

package webview

import (
	"github.com/refaktor/rye/env"
)

var Builtins_webview = map[string]*env.Builtin{}
