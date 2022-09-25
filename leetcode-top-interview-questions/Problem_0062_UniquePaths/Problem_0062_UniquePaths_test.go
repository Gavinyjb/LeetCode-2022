package Problem_0062_UniquePaths

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGcd(t *testing.T) {
	for i := 0; i < 10; i++ {
		m := rand.Intn(20)
		n := rand.Intn(20)
		m, n = max(m, n), min(m, n)
		fmt.Println("m,n", m, n)
		fmt.Println(gcd(m, n) == getGreatestCommonDivsor(m, n))

	}
}
