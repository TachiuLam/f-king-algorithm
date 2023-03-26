package 查分数组

// Difference 差分数组工具类
type Difference struct {
	// 差分数组
	diff []int
}

func newDifference(length int) *Difference {
	return &Difference{
		diff: make([]int, length),
	}
}

// increase 查分数组增减数值
func (d *Difference) increase(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// result 将查分数组还原为原始数组
func (d *Difference) result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = d.diff[i] + res[i-1]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	difference := newDifference(1001)
	for _, trip := range trips {
		// 第 trip[2]站乘客已经下车
		u1, u2, u3 := trip[0], trip[1], trip[2]-1
		difference.increase(u2, u3, u1)
	}
	res := difference.result()
	// res的每一个元素都不应该超过capacity
	for i := 0; i < len(res); i++ {
		if res[i] > capacity {
			return false
		}
	}
	return true
}
