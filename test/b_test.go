package test

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	t.Log(Len, len(s))

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	t.Log(Cap, cap(s))

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	t.Log(count)
}

type Programmer struct {
	name     string
	age      int
	language string
}

func TestUnsafeSizeof(t *testing.T) {
	p := Programmer{"stefno", 18, "go"}
	t.Log(p)

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"

	t.Log(p)
}
