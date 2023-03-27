package test

import (
	"context"
	"testing"
	"time"
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

type MyContext struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	context.Context
}

func TestContext(t *testing.T) {
	childCancel := true

	parentCtx, parentFunc := context.WithCancel(context.Background())
	mctx := MyContext{parentCtx}

	childCtx, childFun := context.WithCancel(mctx)

	if childCancel {
		childFun()
	} else {
		parentFunc()
	}

	t.Log(parentCtx)
	t.Log(mctx)
	t.Log(childCtx)

	// 防止主协程退出太快，子协程来不及打印
	time.Sleep(10 * time.Second)
}
