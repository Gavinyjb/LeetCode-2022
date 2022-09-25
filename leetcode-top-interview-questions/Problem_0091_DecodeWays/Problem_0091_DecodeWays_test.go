package Problem_0091_DecodeWays

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := "2611055971756562"
	ans := numDecodings(s)
	fmt.Println(ans)
}
