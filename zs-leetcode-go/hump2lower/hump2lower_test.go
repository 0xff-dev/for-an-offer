package hump2lower

import (
	"fmt"
	"strings"
	"testing"
)

func TestHump2lower(t *testing.T) {
	names := []string{
		"userName", // user_name
		"UserName", // user_name
		"USERName", // username
	}
	for _, name := range names {
		fmt.Println(name, "---", strings.ToLower(hump2lower(name)))
	}
}
