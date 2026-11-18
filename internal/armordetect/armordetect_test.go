// Copyright 2019 The age Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package armordetect_test

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"filippo.io/age/internal/armordetect"
	"filippo.io/age/internal/armorformat"
)

func TestIsArmored(t *testing.T) {
	armored := armorformat.Header + "\n" + armorformat.Footer + "\n"
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"plain age file", "age-encryption.org/v1\n", false},
		{"generic PEM", "-----BEGIN OPENSSH PRIVATE KEY-----\n", false},
		{"armored", armored, true},
		{"armored with leading whitespace", "   \n\t \n" + armored, true},
		{"armored with maximum leading whitespace",
			strings.Repeat(" ", armorformat.MaxWhitespace) + armored, true},
		{"leading whitespace over budget hides the header",
			strings.Repeat(" ", armorformat.MaxWhitespace+1) + armored, false},
		{"empty input", "", false},
		{"only whitespace", strings.Repeat("\n", 10), false},
		{"header prefix only", armorformat.Header[:10], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := bufio.NewReader(strings.NewReader(tt.in))
			if got := armordetect.IsArmored(br); got != tt.want {
				t.Fatalf("IsArmored = %v, want %v", got, tt.want)
			}
			// IsArmored must not consume any bytes from the underlying input.
			rest, err := io.ReadAll(br)
			if err != nil {
				t.Fatalf("ReadAll: %v", err)
			}
			if string(rest) != tt.in {
				t.Errorf("input was consumed: got %q, want %q", rest, tt.in)
			}
		})
	}
}
