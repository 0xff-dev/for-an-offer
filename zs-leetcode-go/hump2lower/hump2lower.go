// 驼峰命名转换成小写加_
// UserName --> user_name
package hump2lower

import (
	"regexp"
)

func hump2lower(str string) string {
	match := regexp.MustCompile(`([a-z0-9])([A-Z][a-z])`)
	s := match.ReplaceAllString(str, "${1}_${2}")
	return s
}
