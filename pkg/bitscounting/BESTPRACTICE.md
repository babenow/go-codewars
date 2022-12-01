
```go
import "math/bits"

func CountBits(n uint) int {
  return bits.OnesCount(n)
}
```

```go
func CountBits(n uint) int {
  var res int = 0
  for (n>0) {
    if (n & 1 == 1) {
      res = res + 1
    }
    n = n >> 1
  }
  return res
}
```