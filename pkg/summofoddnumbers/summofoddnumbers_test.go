// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package summofoddnumbers_test

import (
	"fmt"
	"testing"

	"github.com/babenow/go-codewars/pkg/summofoddnumbers"
	"github.com/stretchr/testify/assert"
)

func TestRowSummOddnumbers(t *testing.T) {
	tcs := []struct {
		level  int
		result int
	}{
		{level: 1, result: 1},
		{level: 2, result: 8},
		{level: 13, result: 2197},
		{level: 19, result: 6859},
		{level: 41, result: 68921},
		{level: 42, result: 74088},
		{level: 74, result: 405224},
		{level: 86, result: 636056},
		{level: 93, result: 804357},
		{level: 101, result: 1030301},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Level %d", tc.level), func(t *testing.T) {
			assert.Equal(t, summofoddnumbers.RowSummOddNumbers(tc.level), tc.result)
		})
	}
}
