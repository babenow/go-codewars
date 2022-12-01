// Copyright at 2022 Maks Babenko <babenoff.code@outlook.com>.
// All rights reserved
// Licensed under the MIT license that can be found in the LICENSE file.
package digitalroot

import (
	"fmt"
	"strconv"
)

// DigitalRoot Функция получения числового корня - рекурсивной
// суммы всех цифр числа. Учитывая n необходимо взять сумму
// всех цифр n. Если полуученная суммма > 9 (имеет более одной цифры)
// то его нужно уменьшить этой же функцией пока не будет достигнут
// необходимый результат
//
//			Например
//			16 -> 1 + 6 = 7 // Все верно
//	 	942 -> 9 + 4 + 2 = 15 -> 1 + 5 = 6 // Верно в 2 этапа
func DigitalRoot(n int) int {
	var vec []int
	var res int
	for _, ch := range fmt.Sprintf("%d", n) {
		i, _ := strconv.Atoi(string(ch))
		vec = append(vec, i)
	}

	for _, v := range vec {
		res += v
	}

	if res > 9 {
		return DigitalRoot(res)
	}

	return res
}
