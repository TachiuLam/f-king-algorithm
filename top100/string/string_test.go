package string

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_isValid(t *testing.T) {
	var isValid func(s string) bool
	isValid = func(s string) bool {
		if len(s)%2 != 0 {
			return false
		}
		validMap := map[byte]byte{
			'}': '{', ']': '[', ')': '(',
		}
		stack := make([]byte, 0)
		for _num := range s {
			if _, ok := validMap[s[_num]]; !ok {
				stack = append(stack, s[_num])
			} else {
				if len(stack) == 0 || stack[len(stack)-1] != validMap[s[_num]] {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
		return len(stack) == 0
	}
	assert.Equal(t, isValid("()"), true)
	assert.Equal(t, isValid("([])"), true)
	assert.Equal(t, isValid("([)"), false)
	assert.Equal(t, isValid("([)]"), false)
}

func Test_longestPalindrome(t *testing.T) {
	var longestPalindrome func(s string) string
	longestPalindrome = func(s string) string {
		res := ""
		for i := 0; i < len(s)-1; i++ {
			s1 := findP(s, i, i)
			s2 := findP(s, i, i+1)
			if len(s1) > len(res) {
				res = s1
			}
			if len(s2) > len(res) {
				res = s2
			}
		}
		return res
	}
	assert.Equal(t, longestPalindrome("abc"), "a")
	assert.Equal(t, longestPalindrome("abcba"), "abcba")
	assert.Equal(t, longestPalindrome("abadefef"), "aba")
	assert.Equal(t, longestPalindrome("sdeaaf"), "aa")
}

func findP(s string, l, r int) string {
	if len(s) == 1 {
		return s
	}
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
