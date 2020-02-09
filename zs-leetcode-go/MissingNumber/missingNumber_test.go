package MissingNumber
import  (
    "fmt"
    "testing"
)

func TestMissingNumber(t *testing.T) {
    for _, testCase := range [][]int{
        {3, 0, 1},
        {9, 6, 4,2,3,5,7,0,1},
    } {
        fmt.Println(missingNumber(testCase))
    }
}
