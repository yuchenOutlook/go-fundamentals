package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// sort in descending order
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	fmt.Println("Strings sorted using sort.Slice(), sorted in descending order:", strs)

	slices.SortFunc(strs, func (a, b string) int {
		return strings.Compare(a, b)
	})
	fmt.Println("Strings sorted using slices.SortFunc(), sort again in ascending order: ", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
}