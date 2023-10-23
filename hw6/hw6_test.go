package hw6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len)
		require.Nil(t, l.Head)
		require.Nil(t, l.Last)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len)

		middle := 2      // 20
		l.Remove(middle) // [10, 30]
		require.Equal(t, 2, l.Len)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len)
		require.Equal(t, 80, l.Head.Value)
		require.Equal(t, 70, l.Last.Value)

		l.MoveToFront(1) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(7) // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len)
		for i := l.Head; i != nil; i = i.Next {
			elems = append(elems, i.Value)
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

/*package hw6_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PushBack(t *testing.T) {
	type Testsnode = struct {
		Value int
		Next  *Testsnode
		Prev  *Testsnode
	}
	tests := []struct {
		i       int

	}{
		i:      1,
		Len:    0,
		Head:   nil,
		Last:   nil,
		ResLen: 1,
		ResHead: &Testsnode{
			Value: 1,
		},
		ResLast: &Testsnode{
			Value: 1,
		},
	}
	for _, test := range tests {
		a := List{}
		result := a.PushBack(test.i)
		assert.Equal(t, test.ResLast, result)
	}
}*/
