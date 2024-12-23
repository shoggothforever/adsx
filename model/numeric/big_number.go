package numeric

import (
	"fmt"
	"strings"
)

// 加法
func BigAdd(a, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = fmt.Sprintf("%d", sum%10) + result
		carry = sum / 10
	}
	return result
}

// 减法
func BigSub(a, b string) string {
	flag := false
	if BigCmp(b, a) == 1 {
		a, b = b, a
		flag = true
	}
	i, j := len(a)-1, len(b)-1
	borrow := 0
	result := ""

	for i >= 0 || j >= 0 {
		diff := borrow
		if i >= 0 {
			diff += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			diff -= int(b[j] - '0')
			j--
		}
		if diff < 0 {
			diff += 10
			borrow = -1
		} else {
			borrow = 0
		}
		result = fmt.Sprintf("%d", diff) + result
	}
	// 移除前导零
	result = strings.TrimLeft(result, "0")
	if flag {
		result = "-" + result
	}
	if result == "" {
		return "0"
	}
	return result
}

// 乘法
func BigMul(a, b string) string {
	m, n := len(a), len(b)
	result := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(a[i]-'0') * int(b[j]-'0')
			sum := result[i+j+1] + mul
			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	// 将结果转为字符串
	var w strings.Builder

	for _, v := range result {
		if !(w.Len() == 0 && v == 0) { // 去掉前导零
			w.WriteByte(byte(v + '0'))
		}
	}
	if w.Len() == 0 {
		return "0"
	}
	return w.String()
}

// 除法
func BigDiv(a, b string) (string, string) {
	if b == "0" {
		panic("division by zero")
	}
	if a == "0" {
		return "0", "0"
	}
	// 去掉前导零
	a = strings.TrimLeft(a, "0")
	b = strings.TrimLeft(b, "0")

	// 如果被除数小于除数，直接返回商为 0，余数为 a
	if BigCmp(a, b) < 0 {
		return "0", a
	}

	// 商和余数初始化
	remainder := a
	r := "1"
	bond := (len(a) - len(b) + 1) * 4
	strs := make([]string, bond)
	for i := 0; i < bond; i++ {
		strs[bond-1-i] = r
		r = BigMul(r, "2")
	}
	result := "0"
	// 二进制模拟除法：从高位开始
	for i := 0; i < bond; i++ {
		// 将除数左移 i 位
		shiftedB := BigMul(b, strs[i])
		// 判断余数是否大于等于左移后的除数
		if BigCmp(remainder, shiftedB) >= 0 {
			//fmt.Println("bigger!!!")
			remainder = BigSub(remainder, shiftedB)
			result = BigAdd(result, strs[i])
		}
	}

	return result, remainder // 返回商和余数
}

// 比较 a 和 b 的大小
// 返回：1 表示 a > b，0 表示 a == b，-1 表示 a < b
func BigCmp(a, b string) int {
	a = strings.TrimLeft(a, "0")
	b = strings.TrimLeft(b, "0")

	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				return 1
			} else if a[i] < b[i] {
				return -1
			}
		}
	}
	return 0
}
