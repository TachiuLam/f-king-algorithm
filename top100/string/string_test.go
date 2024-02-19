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
