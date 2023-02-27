package main

type S struct {
	M *int
}

func ref(y *int) (z S) {
	z.M = y
	return z
}

func main() {
	//var x S
	var i int
	ref(&i)
}
