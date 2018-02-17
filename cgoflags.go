package ffi

// usage: PKG_CONFIG_PATH=/androidsys/lib/pkgconfig

// #cgo pkg-config: libffi
// #cgo CFLAGS: -D_GNU_SOURCE
// #cgo LDFLAGS: -lffi
// #include "ffi.h"
import "C"

// EOF
