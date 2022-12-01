// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package bitscounting

import (
	"fmt"
)

func BitsCount(n uint) int {
	var res int = 0
	stmt := fmt.Sprintf("%b", n)
	for _, ch := range stmt {
		if ch == '1' {
			res++
		}
	}
	return res
}
