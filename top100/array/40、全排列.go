package array

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
// https://leetcode.cn/problems/permutations/solutions/481119/46-quan-pai-lie-hui-su-suan-fa-jing-dian-ti-mu-xia/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
func permute(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, len(nums))
	state := make([]bool, len(nums))

	var dfs func(nums []int, cur int)
	dfs = func(nums []int, cur int) {
		if cur == len(nums) { // 退出
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !state[i] { // 该位置未被排列
				path = append(path, nums[i])
				state[i] = true
				dfs(nums, cur+1)
				state[i] = false          // 下个循环重新使用
				path = path[:len(path)-1] // 重置末尾的数字，尝试其它值的排序
			}
		}
	}

	dfs(nums, 0)
	return res
}
