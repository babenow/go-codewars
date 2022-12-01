Мой вариант решения оказался лучшим вариантом

```go
import "math"

// RowSummOddNumbers Функция вычисляет сумму чисел в нечетном треугольнике
func RowSummOddNumbers(n int) int {
	return int(math.Pow(float64(n), 3))
}
```