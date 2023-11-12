package binary_tree

func numTrees(n int) int {
	tmp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		tmp[i] = make([]int, n+1)
	}
	return count(1, n, tmp)
}

func count(low, high int, tmp [][]int) int {
	if low > high {
		return 1 // 叶子节点为nil，是一种情况，返回1
	}
	if tmp[low][high] != 0 { // 已经计算过该情况
		return tmp[low][high]
	}
	res := 0
	for i := low; i <= high; i++ {
		left := count(low, i-1, tmp)
		right := count(i+1, high, tmp)
		res += left * right
	}
	tmp[low][high] = res
	return res
}
