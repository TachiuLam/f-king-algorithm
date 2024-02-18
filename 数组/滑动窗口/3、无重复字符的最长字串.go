package 滑动窗口

func lengthOfLongestSubstring(s string) int {
	var (
		exits  map[byte]struct{}
		res    [][2]int
		source = []byte(s)
		l      = 0
		r      = 1
	)
	for _, each := range source {
		if _, ok := exits[each]; !ok {
			exits[each] = struct{}{}
		} else {
			delete(exits, each)
			res = append(res, [2]int{l, r})
			l++
		}
		r++
	}

	var maxLength int
	for _, eachRes := range res {
		length := eachRes[1] - eachRes[0]
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
