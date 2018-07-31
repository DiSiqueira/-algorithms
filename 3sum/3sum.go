package sum

import (
	"fmt"
	"sort"

	"github.com/disiqueira/algorithms/binarysearch"
)

func ThreeSum(data []int) int {
	sort.Ints(data)

	total := 0
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			complement := -(data[i] + data[j])

			if binarysearch.Search(complement, data) &&
				data[i] < data[j] &&
				data[j] < complement{

				fmt.Printf("%d %d %d\n", data[i], data[j], complement)

				total++
			}
		}
	}

	return total
}
