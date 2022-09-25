package Problem_0043_Multiply

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	for i := 0; i < 80; i++ {
		nums1 := rand.Intn(10000)
		nums2 := rand.Intn(10000)
		wants := nums1 + nums2
		ans := plus(strconv.Itoa(nums1), strconv.Itoa(nums2))
		want := strconv.Itoa(wants)
		if ans != want {
			t.Errorf("want:%v got:%v\n", want, ans)
		}
	}
}
func TestMultiply(t *testing.T) {
	for i := 0; i < 80; i++ {
		nums1 := rand.Intn(10000)
		nums2 := rand.Intn(10000)
		wants := nums1 * nums2
		ans := multiply(strconv.Itoa(nums1), strconv.Itoa(nums2))
		want := strconv.Itoa(wants)
		if ans != want {
			t.Errorf("want:%v got:%v\n", want, ans)
		}
	}
}

func TestTemp(t *testing.T) {

}
