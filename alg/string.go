package alg

import (
	"fmt"
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

type MyersNode struct {
	x, y int
	pre  *MyersNode
}

// diff算法原型，计算最小编辑距离
func Myers(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	res := max(m, n)
	//y=x-k (k轴坐标表示)
	type Kaxis int
	type Kmap map[int]*MyersNode
	type Depth int
	type DeepMap map[int]Kmap
	kmap := make(Kmap)
	kmap[1] = &MyersNode{
		x:   0,
		y:   -1,
		pre: nil,
	}
	for dep := 0; dep <= m+n; dep++ {
		for k := -dep; k <= dep; k += 2 {
			down := false
			var kprev = k - 1
			if k == -dep || (k != dep) && kmap[k+1].x > kmap[k-1].x {
				down = true
				kprev = k + 1
			}
			preNode := kmap[kprev]
			x := preNode.x + 1
			if down {
				x = preNode.x
			}
			y := x - k
			for x < m && y < n && s1[x] == s2[y] {
				x++
				y++
			}
			node := &MyersNode{x: x, y: y, pre: preNode}
			kmap[k] = node
			if x == m && y == n {
				fmt.Printf("endnode %+v\n", *node)
				getMovement(node, &s1, &s2)
				return dep
			}
		}
	}
	return res
}
func getMovement(node *MyersNode, s1, s2 *string) {
	if node == nil || node.pre == nil {
		return
	}
	getMovement(node.pre, s1, s2)
	x := node.x
	y := node.y
	if node.pre != nil {
		var str string
		for x > node.pre.x && y > node.pre.y {
			x--
			y--
			str = fmt.Sprintf("%c", (*s1)[x]) + str
		}
		if len(str) > 0 {
			fmt.Println(str)
		}
		if node.pre.x+1 == node.x {
			fmt.Printf("-%c\n", (*s1)[node.x])
		} else if node.pre.y+1 == node.y {
			fmt.Printf("+%c\n", (*s2)[node.y])
		}
	}
	fmt.Printf("(%d,%d)\n", node.x, node.y)
}
