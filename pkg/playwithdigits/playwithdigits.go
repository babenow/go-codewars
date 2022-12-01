// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package playwithdigits

import (
	"fmt"
	"math"
	"strconv"
)

const NegativeAnswer int = -1

// DigPow функция должна найти положительное целое число k, если оно
// существует, такое, что сумма цифр n, взятых в последовательных степенях p,
// равна k * n.
// 		Например:
// 		DigPow(89, 1) -> 1 // 8 ^ 1 + 9 ^ 2 = 89 = 89 * 1
// 		DigPow(695, 2) -> 2 // 6 ^ 2 + 9 ^ 3 + 5 ^ 4 = 1390 = 695 *2
func DigPow(n, p int) int {
	var summ = 0
	for i, ch := range fmt.Sprintf("%d", n) {
		k, _ := strconv.Atoi(string(ch))
		summ = summ + powInt(k, p+i)
	}

	kk := int(summ) % n

	if kk == 0 {
		return int(summ / n)
	}

	return NegativeAnswer
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
