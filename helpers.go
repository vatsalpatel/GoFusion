package ut

import (
	"fmt"
	"sort"
)

// Check interface{} slices for equality
func interfaceSliceEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to sort a slice of interface{} elements
func sortSlice(slice []interface{}) {
	sort.Slice(slice, func(i, j int) bool {
		return fmt.Sprintf("%v", slice[i]) < fmt.Sprintf("%v", slice[j])
	})
}
