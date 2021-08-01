package stack

import "testing"

func TestIntStack(t *testing.T) {
	sut := intStack{}
	sut.Push(1)
	sut.Push(2)
	sut.Push(3)
	sut.Push(4)

	for i := 4; i > 0; i-- {
		got := sut.Pop()
		want := i
		if got != want {
			t.Errorf("pop()= %v, want: %v", got, want)
		}
	}
}
