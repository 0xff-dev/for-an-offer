package ExcelSheetColumnTitle

import (
	"bytes"
	"fmt"
)

func convertToTitle(n int) string {
	const ch string = "ZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	resCh := make([]byte, 0)
	for n > 26 {
		now := n%26
		resCh = append(resCh, ch[now])
		n = n/26
		if now == 0 {
			n--
		}
	}
	resCh = append(resCh, ch[n])
	buf := bytes.NewBufferString("")
	for index := len(resCh)-1; index > -1; index-- {
		buf.WriteString(fmt.Sprintf("%c", resCh[index]))
	}
	return buf.String()
}