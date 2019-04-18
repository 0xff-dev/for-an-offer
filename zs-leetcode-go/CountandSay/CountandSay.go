package CountandSay

import "fmt"
import "bytes"

func countAndSay(n int) string {
	constDatas := []string{"", "1", "11", "21", "1211", "111221"}
	if n <= 5 {
		return constDatas[n]
	}
	nowString := constDatas[5]
	for i := 6; i <= n; i++ {
		buf := bytes.NewBufferString("")
		start, end := 0, 0
		cnt := 0
		data := nowString
		for end != len(data) {
			for end < len(data) && data[end] == data[start] {
				end ++
				cnt ++
			}
			buf.WriteString(fmt.Sprintf("%d%c", cnt, data[start]))
			start = end
			cnt = 0
		}
		nowString = buf.String()
	}
	return nowString
}
