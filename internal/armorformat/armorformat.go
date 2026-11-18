// Copyright 2019 The age Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package armorformat holds the wire-level constants of the age armor
// format and the pure byte-prefix predicate that recognizes an armored
// header. It is a leaf package with no dependencies on the armor package,
// so both armor itself and the higher-level armordetect package can share
// a single definition of these values.
package armorformat

import "bytes"

const (
	// Header is the first line of an armored age file.
	Header = "-----BEGIN AGE ENCRYPTED FILE-----"
	// Footer is the last line of an armored age file.
	Footer = "-----END AGE ENCRYPTED FILE-----"
)

// MaxWhitespace is the maximum amount of leading whitespace tolerated
// before an armored age header, both by the detection probe and by the
// streaming armor reader.
const MaxWhitespace = 1024

// HasHeaderPrefix reports whether b, typically a buffered peek of the
// beginning of a file, starts with the armored age header after ignoring
// leading whitespace.
func HasHeaderPrefix(b []byte) bool {
	return bytes.HasPrefix(bytes.TrimSpace(b), []byte(Header))
}
