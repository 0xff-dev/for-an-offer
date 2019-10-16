package PowerofTwo

import (
	"fmt"
	"testing"
	"time"
)

func TestIsPowerOfTwo(t *testing.T) {
	now := time.Now()
	for _, item := range []int{1856237854, 5748577434, 23478275375, 758374873, 74837439, 48734737827387, 7763623623726,
		10000000, 489384938493, 3877373927392, 73873827377, 7483747343, 4773737837, 4737372329} {
		//fmt.Println(isPowerOfTwo(item))
		fmt.Println(isPowerOfTwo1(item))
	}
	fmt.Println(time.Since(now))
}
