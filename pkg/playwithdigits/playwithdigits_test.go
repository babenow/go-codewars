// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package playwithdigits_test

import (
	"testing"

	"github.com/babenow/go-codewars/pkg/playwithdigits"
	"github.com/stretchr/testify/assert"
)

func TestDigPow(t *testing.T) {
	tcs := []struct {
		name   string
		input  func() (n int, p int)
		result int
	}{
		{
			name: "Number 89 and 1",
			input: func() (n int, p int) {
				return 89, 1
			},
			result: 1,
		},
		{
			name: "Number 695 and 2",
			input: func() (n int, p int) {
				return 695, 2
			},
			result: 2,
		},
		{
			name: "Number 92 and 1",
			input: func() (n int, p int) {
				return 92, 1
			},
			result: playwithdigits.NegativeAnswer,
		},
		{
			name: "Number 46288 and 3",
			input: func() (n int, p int) {
				return 46288, 3
			},
			result: 51,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			n, p := tc.input()
			assert.Equal(t, playwithdigits.DigPow(n, p), tc.result)
		})
	}
}
