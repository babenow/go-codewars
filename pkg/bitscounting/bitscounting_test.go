// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package bitscounting_test

import (
	"testing"

	"github.com/babenow/go-codewars/pkg/bitscounting"
	"github.com/stretchr/testify/assert"
)

func TestBitsCounting(t *testing.T) {
	tcs := []struct {
		name   string
		input  uint
		output int
	}{
		{
			name:   "Number 0",
			input:  0,
			output: 0,
		},
		{
			name:   "Number 1",
			input:  1,
			output: 1,
		},
		{
			name:   "Number 7",
			input:  7,
			output: 3,
		},
		{
			name:   "Number 9",
			input:  9,
			output: 2,
		},
		{
			name:   "Number 10",
			input:  10,
			output: 2,
		},
		{
			name:   "Number 1234",
			input:  1234,
			output: 5,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bitscounting.BitsCount(tc.input), tc.output)
		})
	}
}
