// Copyright 2019 The age Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package armordetect contains the single shared implementation of the
// "is this an armored age file?" probe used by the age command line's
// decryption and identity-file loading entry points and by the metadata
// inspector.
//
// Keeping the peek budget and the prefix check in one place ensures that
// every entry point recognizes exactly the same armored inputs, including
// the tolerance for up to [armorformat.MaxWhitespace] bytes of leading
// whitespace.
package armordetect

import (
	"bufio"

	"filippo.io/age/internal/armorformat"
)

// MaxWhitespace is the maximum amount of leading whitespace tolerated
// before an armored age header. It matches the leading-whitespace budget
// enforced by armor.NewReader.
const MaxWhitespace = armorformat.MaxWhitespace

// IsArmored reports whether the beginning of r is an armored age file.
//
// It peeks at the first MaxWhitespace bytes plus the length of the armor
// header, ignoring leading whitespace, and reports true only if the age
// armor header follows. It never consumes bytes from r, so the caller can
// still read the full input regardless of the result.
func IsArmored(r *bufio.Reader) bool {
	start, _ := r.Peek(MaxWhitespace + len(armorformat.Header))
	return armorformat.HasHeaderPrefix(start)
}
