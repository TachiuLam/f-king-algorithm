package 数组

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词
*/
func reverseWords(s string) string {
	// 双指针
	fast, slow := 0, 0
	b := []byte(s)
	// 1、去冗余空格
	// 去除字符串头部空格
	for len(b) > 0 && fast < len(b) && b[fast] == ' ' {
		fast++
	}
	// 去除单词间冗余空格
	for ; fast < len(b); fast++ {
		// 跳过连续空格
		if fast-1 > 0 && b[fast-1] == b[fast] && b[fast] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow++
	}
	// 删除尾部空格
	if slow-1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	// 2、反转字符串
	reverse(&b, 0, len(b)-1)

	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
