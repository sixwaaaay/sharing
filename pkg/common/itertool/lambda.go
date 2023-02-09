package itertool

// Reduce
//
// collection: the values to be used in reduce function
//
// accumulator: the handle function. three parameters represent
// the aggregated result, the current item, and the current index
//
// initial: the initial value for reduce function
//
//	Reduce([]int{1, 2, 3, 4}, func(acc float64, num int, _ int) float64 {
//		return acc + float64(num)
//	}, float64(0)) // 10
func Reduce[T any, R any](collection []T, accumulator func(R, T, int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}
	return initial
}

// ForEach
//
// collection: the values to be used in foreach function
//
// consumer: the handle function. two parameters represent
// the current item, and the current index
//
//	var items []string
//	ForEach([]string{"for", "each"}, func(item string, i int) {
//		items = append(items, item)
//	})// items => []string{"for", "each"}
func ForEach[T any](collection []T, consumer func(T, int)) {
	Reduce(collection, func(agg struct{}, item T, i int) struct{} {
		consumer(item, i)
		return agg
	}, struct{}{})
}
