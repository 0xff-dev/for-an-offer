package CountandSay

import "fmt"

func countAndSay(n int) string {
	constDatas := []string{"", "1", "11", "21", "1211", "111221"}
	if n <= 5 {
		return constDatas[n]
	}
	nowString := constDatas[5]
	for i := 6; i <= n; i++ {
		tmpString := ""
		start, end := 0, 0
		cnt := 0
		data := nowString
		for end != len(data) {
			for end < len(data) && data[end] == data[start] {
				end ++
				cnt ++
			}
			tmpString += fmt.Sprintf("%d%c", cnt, data[start])
			start = end
			cnt = 0
		}
		nowString = tmpString
	}
	return nowString
}
