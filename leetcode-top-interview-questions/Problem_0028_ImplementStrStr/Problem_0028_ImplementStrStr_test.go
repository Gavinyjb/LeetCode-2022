package Problem0028implementstrstr

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	needle := "bba"
	ms := []byte(needle)
	next := getNexts(ms)
	fmt.Println(next)
}
