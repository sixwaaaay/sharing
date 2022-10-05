package itertool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReduce(t *testing.T) {
	assertions := require.New(t)
	result := Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	assertions.Equal(result, 10)

	result = Reduce([]int{1, 2, 3, 4}, func(agg int, _ int, i int) int {
		return agg + i
	}, 0)
	assertions.Equal(result, 6)

	res := Reduce([]int{1, 2, 3, 4}, func(agg float64, item int, _ int) float64 {
		return agg + float64(item)
	}, float64(0))
	assertions.Equal(res, 10.0)

}

func TestForEach(t *testing.T) {
	var items []string
	var index []int
	assertions := require.New(t)
	ForEach([]string{"hello", "world"}, func(item string, i int) {
		items = append(items, item)
		index = append(index, i)
	})
	assertions.ElementsMatch([]string{"hello", "world"}, items)
	assertions.ElementsMatch([]int{0, 1}, index)
}
