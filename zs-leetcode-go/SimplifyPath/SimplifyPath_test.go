package SimplifyPath

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	for _, path := range []string{"/home", "/home/", "/../", "/home//goo/",
		"/a/./b/../../c/", "/a/../../b/../c//.//", "/a//b////c/d//././/..", "/", "/////", "/."} {
		fmt.Println(simplifyPath(path))
	}
}
