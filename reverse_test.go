package tools

import (
	"testing"
	"fmt"
)

var ts = []struct {
	in  string
	out string
}{
	{"t", "t"},
	{"hello", "olleh"},
	{"hell", "lleh"},
}

func Test_rev(t *testing.T) {
	for _, v := range ts {
		if s := Reverse(v.in); s != v.out {
			fmt.Printf("Reverse(%s)=%s;expect %s\n",
				v.in, s, v.out)
		}
	}
}
