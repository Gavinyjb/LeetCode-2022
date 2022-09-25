package chouxiang

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var p Person
	p = NewPerson("yvjinbo", 23)
	fmt.Println(p)
}
