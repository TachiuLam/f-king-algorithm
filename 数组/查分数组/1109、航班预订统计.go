package 查分数组

type Diff struct {
	diff []int
}

func NewDiff(n int) *Diff {
	diff := make([]int, n)
	return &Diff{diff: diff}
}

func (d *Diff) Increase(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Diff) Result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := NewDiff(n)
	for _, each := range bookings {
		i, j, val := each[0]-1, each[1]-1, each[2]
		diff.Increase(i, j, val)
	}
	return diff.Result()
}
