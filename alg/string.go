package alg

import (
	"github.com/shoggothforever/adsx/model/binary_search"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

// 求两个字符串的最长公共子序列
func Lcs(s1, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)
	n := len(r1)
	m := len(r2)
	f := make([]int, m+1)
	tmp := 0
	pre := 0
	for i := 1; i <= n; i++ {
		pre = f[0]
		for j := 1; j <= m; j++ {
			tmp = f[j]
			f[j] = max(f[j-1], f[j])
			if r1[i-1] == r2[j-1] {
				f[j] = max(pre+1, f[j])
			}
			pre = tmp
		}
	}
	return f[m]
}

// 求最长上升子序列
func Lis(s string) int {
	r := []rune(s)
	seq := make([]int, 0)
	for i := 0; i < len(r); i++ {
		lb := binary_search.LowerBound(seq, int(r[i]))
		if lb >= len(seq) {
			seq = append(seq, int(r[i]))
		} else {
			seq[lb] = int(r[i])
		}
	}
	return len(seq)
}
