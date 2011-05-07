package tools

func Reverse(s string) string {
	l := []int(s)
	n, h := len(s), len(s)/2
	for i := 0; i < h; i++ {
		l[i], l[n-i-1] = l[n-i-1], l[i]
	}
	return string(l)
}
