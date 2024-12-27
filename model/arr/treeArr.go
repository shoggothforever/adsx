package arr

func LowBit(x int) int {
	return x & -x
}

type TreeArr struct {
	// 原始数组
	Nums []int
	// 用于计算区间和的树状数组
	Trees []int
}

func NewTreeArr(nums []int) *TreeArr {
	arr := &TreeArr{Nums: make([]int, len(nums)), Trees: make([]int, len(nums)+1)}
	for i := 0; i < len(nums); i++ {
		arr.update(i, nums[i])
	}
	return arr
}

func (t *TreeArr) update(index int, val int) {
	delta := val - t.Nums[index]
	t.Nums[index] = val
	for i := index + 1; i < len(t.Trees); i += LowBit(i) {
		t.Trees[i] += delta
	}
}
func (t *TreeArr) getSum(x int) int {
	sum := 0
	for ; x > 0; x &= x - 1 {
		sum += t.Trees[x]
	}
	return sum
}
func (t *TreeArr) SumRange(l, r int) int {
	return t.getSum(r+1) - t.getSum(l)
}
