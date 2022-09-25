package Problem_0069_SqrtX

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestName(t *testing.T) {
	for i := 0; i < 30; i++ {
		x := rand.Intn(100)
		right := int(math.Sqrt(float64(x)))
		if mySqrt(x) != right {
			fmt.Println("x,right,mine", x, right, mySqrt(x))
		}
	}
}
