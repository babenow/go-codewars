// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package bitscounting

import (
	"fmt"
)

// BitsCount Функция принимает на вход неотрицательное число n
// и возвращает количество единиц, которое содержатся в
// полученномм из него двоичном числе
//
//		Например:
//		BitsCount(1) -> 1 // 1
//		BitsCount(7) -> 3 // 111
//	 	BitsCount(9) -> 2 // 1001
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
