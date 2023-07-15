package fusion

import (
	"fmt"
	"math/rand"
	"time"
)

// Chunk splits a slice into chunks of a specified size
func Chunk[T any](arr []T, size int) [][]T {
	length := len(arr)
	if size <= 0 || length == 0 {
		return nil
	}

	chunks := make([][]T, 0)
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		chunks = append(chunks, arr[i:end])
	}

	return chunks
}

// Concat concatenates two or more slices into a single slice
func Concat[T any](arrays ...[]T) []T {
	length := 0
	for _, arr := range arrays {
		length += len(arr)
	}

	concatenated := make([]T, 0, length)
	for _, arr := range arrays {
		concatenated = append(concatenated, arr...)
	}

	return concatenated
}

// Difference creates a slice of values that are not in the other given slices
func Difference[T comparable](arr []T, others ...[]T) []T {
	diffSet := make(map[T]struct{})
	for _, val := range arr {
		diffSet[val] = struct{}{}
	}

	for _, other := range others {
		for _, val := range other {
			delete(diffSet, val)
		}
	}

	difference := make([]T, 0, len(diffSet))
	for val := range diffSet {
		difference = append(difference, val)
	}

	return difference
}

// Every checks if all elements in the slice satisfy the given predicate.
// The predicate is invoked with three arguments: (value, index, array).
// It returns true if the predicate returns true for all elements, otherwise false.
func Every[T any](arr []T, predicate func(T, int, []T) bool) bool {
	for i, value := range arr {
		if !predicate(value, i, arr) {
			return false
		}
	}
	return true
}

// Fill modifies a slice by filling it with a specified value from a start index to an end index
func Fill(arr []interface{}, value interface{}, start int, end int) {
	length := len(arr)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	for i := start; i < end; i++ {
		arr[i] = value
	}
}

// Filter filters the elements of a slice based on a predicate function.
// The predicate function is invoked with three arguments: (value, index, arg),
// and should return a boolean indicating whether the value should be included in the result.
func Filter[T any](arr []T, predicate func(int, T, interface{}) bool, arg interface{}) []T {
	var result []T

	for i, value := range arr {
		if predicate(i, value, arg) {
			result = append(result, value)
		}
	}

	return result
}

// FindIndex returns the index of the first element in a slice that satisfies the provided testing function
func FindIndex(arr []interface{}, predicate func(interface{}) bool) int {
	for i, val := range arr {
		if predicate(val) {
			return i
		}
	}
	return -1
}

// Flatten flattens a slice of slices into a single slice
func Flatten(arr []interface{}) []interface{} {
	flattened := make([]interface{}, 0)
	for _, val := range arr {
		switch v := val.(type) {
		case []interface{}:
			flattened = append(flattened, Flatten(v)...)
		default:
			flattened = append(flattened, v)
		}
	}
	return flattened
}

// Includes checks if a given value is present in the slice.
// It returns true if the value is found, otherwise false.
func Includes[T comparable](arr []T, value T) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

// Intersection returns an array containing the unique values that are present in all of the input arrays.
func Intersection(arrays ...[]interface{}) []interface{} {
	// Count occurrences of each element
	counts := make(map[interface{}]int)
	for _, arr := range arrays {
		seen := make(map[interface{}]bool)
		for _, elem := range arr {
			if !seen[elem] {
				counts[elem]++
				seen[elem] = true
			}
		}
	}

	// Filter elements that appear in all arrays
	var result []interface{}
	for elem, count := range counts {
		if count == len(arrays) {
			result = append(result, elem)
		}
	}

	return result
}

// Join concatenates all elements of an array into a single string using the provided separator.
func Join(arr []interface{}, separator string) string {
	if len(arr) == 0 {
		return ""
	}

	var result string
	for i, elem := range arr {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", elem)
	}

	return result
}

// Map applies a transformation function to each element of the input array/slice
// and returns a new array/slice with the transformed values.
func Map[T any, U any](arr []T, transformFunc func(int, T, interface{}) U, arg interface{}) []U {
	result := make([]U, len(arr))

	for i, element := range arr {
		transformedValue := transformFunc(i, element, arg)
		result[i] = transformedValue
	}

	return result
}

// Pull removes all occurrences of the specified values from a slice.
func Pull[T comparable](arr []T, values ...T) []T {
	var result []T
	excluded := make(map[T]struct{})

	for _, value := range values {
		excluded[value] = struct{}{}
	}

	for _, item := range arr {
		if _, excluded := excluded[item]; !excluded {
			result = append(result, item)
		}
	}

	return result
}

// Reduce applies a function against an accumulator and each element in the slice (from left to right)
// to reduce it to a single value.
// The transform function is invoked with four arguments: (accumulator, value, index, arg).
// The initial value of the accumulator is provided as the initialValue parameter.
func Reduce[T any, R any](arr []T, transform func(R, T, int, interface{}) R, initialValue R, arg interface{}) R {
	accumulator := initialValue

	for i, value := range arr {
		accumulator = transform(accumulator, value, i, arg)
	}

	return accumulator
}

// Remove removes all elements from the slice for which the predicate returns truthy,
// and returns an array of the removed elements.
// The predicate is invoked with three arguments: (value, index, array).
func Remove[T any](arr *[]T, predicate func(T, int, []T) bool) []T {
	var removed []T
	remaining := (*arr)[:0]

	for index, value := range *arr {
		if predicate(value, index, *arr) {
			removed = append(removed, value)
		} else {
			remaining = append(remaining, value)
		}
	}

	*arr = remaining
	return removed
}

// Reverse reverses the elements of a slice in place
func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Shuffle shuffles the elements of the slice using the Fisher-Yates algorithm. It modifies the input slice in place.
func Shuffle[T any](arr []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Slice returns a portion of a slice from a start index to an end index
func Slice[T any](arr []T, start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}
	if start > end {
		return nil
	}
	return arr[start:end]
}

// Some checks if at least one element in the slice satisfies the given predicate.
// The predicate is invoked with three arguments: (value, index, array).
// It returns true if the predicate returns true for any element, otherwise false.
func Some[T any](arr []T, predicate func(T, int, []T) bool) bool {
	for i, value := range arr {
		if predicate(value, i, arr) {
			return true
		}
	}
	return false
}

// Union returns a new slice that contains the unique elements from all input slices.
// The order of elements in the resulting slice is undefined.
func Union[T comparable](slices ...[]T) []T {
	// Use a map to track unique elements
	uniqueElements := make(map[T]struct{})

	for _, slice := range slices {
		for _, elem := range slice {
			uniqueElements[elem] = struct{}{}
		}
	}

	// Build the resulting slice from the unique elements
	result := make([]T, 0, len(uniqueElements))
	for elem := range uniqueElements {
		result = append(result, elem)
	}

	return result
}

// Uniq creates a new slice of unique values in the order of their first occurrence in the original slice
func Uniq[T comparable](arr []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, value := range arr {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// Zip merges multiple slices into a single slice of tuples, where each tuple contains
// the corresponding elements from each of the input slices.
// The length of the resulting slice is determined by the shortest input slice.
func Zip[T1, T2 interface{}](slice1 []T1, slice2 []T2) [][]interface{} {
	length := len(slice1)
	if len(slice2) < length {
		length = len(slice2)
	}

	zipped := make([][]interface{}, length)
	for i := 0; i < length; i++ {
		zipped[i] = []interface{}{slice1[i], slice2[i]}
	}

	return zipped
}
