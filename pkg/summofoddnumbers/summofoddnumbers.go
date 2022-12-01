// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package summofoddnumbers

import "math"

// RowSummOddNumbers Функция вычисляет сумму чисел в нечетном треугольнике
func RowSummOddNumbers(n int) int {
	return int(math.Pow(float64(n), 3))
}
