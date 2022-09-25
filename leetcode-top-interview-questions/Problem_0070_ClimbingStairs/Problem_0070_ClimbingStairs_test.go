package Problem_0070_ClimbingStairs

import (
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		n := rand.Intn(10)
		right := help(n)
		//got := dp(n)
		//got := try(n, 2, 0)
		got := dpPlus(n, 2)
		if right != got {
			t.Errorf("dp()错了\n")
		}
	}
}
