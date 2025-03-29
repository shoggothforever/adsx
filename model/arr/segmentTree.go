package arr

const MAXN int = 10001

type SegMentTree struct {
	ini  int
	cap  int
	tree []int
	nums []int
}

func NewSegMentTree(cap int, initialNums []int) *SegMentTree {
	s := &SegMentTree{
		ini:  0,
		cap:  cap,
		tree: make([]int, cap<<2),
		nums: initialNums,
	}
	s.build(1, len(initialNums), 1)
	return s
}
func (s *SegMentTree) push_up(rt int) {
	s.tree[rt] = s.tree[rt<<1] + s.tree[rt<<1|1]
}
func (s *SegMentTree) update(pos, val, l, r, rt int) {
	if l == r {
		s.tree[rt] = val
		return
	}
	mid := (l + r) >> 1
	if pos > mid {
		s.update(pos, val, mid+1, r, rt<<1|1)
	} else {
		s.update(pos, val, l, mid, rt<<1)
	}
	s.push_up(rt)
}
func (s *SegMentTree) build(l, r, rt int) {
	if l == r {
		s.tree[rt] = s.nums[s.ini]
		s.ini++
		return
	}
	mid := (l + r) >> 1
	s.build(l, mid, rt<<1)
	s.build(mid+1, r, rt<<1|1)
	s.push_up(rt)
}
func (s *SegMentTree) query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return s.tree[rt]
	}
	mid := (l + r) >> 1
	ret := 0
	if L <= mid {
		ret += s.query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		ret += s.query(L, R, mid+1, r, rt<<1|1)
	}
	return ret
}
