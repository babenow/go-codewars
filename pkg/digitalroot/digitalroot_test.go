// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package digitalroot_test

import (
	"testing"

	"github.com/babenow/go-codewars/pkg/digitalroot"
	"github.com/stretchr/testify/assert"
)

func TestDigitalRoot(t *testing.T) {
	tcs := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "Number 16",
			input:  16,
			output: 7,
		},
		{
			name:   "Number 195",
			input:  195,
			output: 6,
		},
		{
			name:   "Number 992",
			input:  992,
			output: 2,
		},
		{
			name:   "Number 167346",
			input:  167346,
			output: 9,
		},
		{
			name:   "Number 0",
			input:  0,
			output: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, digitalroot.DigitalRoot(tc.input), tc.output)
		})
	}
}
