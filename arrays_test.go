package ut

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arr      []interface{}
		size     int
		expected [][]interface{}
	}{
		{
			name:     "normal case",
			arr:      []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			size:     3,
			expected: [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8}},
		},
		{
			name:     "empty array",
			arr:      []interface{}{},
			size:     3,
			expected: nil,
		},
		{
			name:     "negative size",
			arr:      []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			size:     -1,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Chunk(tc.arr, tc.size)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arrays   [][]interface{}
		expected []interface{}
	}{
		{
			name:     "normal case",
			arrays:   [][]interface{}{{1, 2}, {3, 4}},
			expected: []interface{}{1, 2, 3, 4},
		},
		{
			name:     "empty array",
			arrays:   [][]interface{}{{}, {}},
			expected: []interface{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Concat(tc.arrays...)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arr      []interface{}
		others   [][]interface{}
		expected []interface{}
	}{
		{
			name:     "normal case",
			arr:      []interface{}{1, 2, 3, 4, 5},
			others:   [][]interface{}{{2, 4}, {4, 6}},
			expected: []interface{}{1, 3, 5},
		},
		{
			name:     "empty array",
			arr:      []interface{}{},
			others:   [][]interface{}{{}, {}},
			expected: []interface{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Difference(tc.arr, tc.others...)
			sort.Slice(actual, func(i, j int) bool {
				return actual[i].(int) < actual[j].(int)
			})

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int, int, []int) bool
		expected  bool
	}{
		{
			name:  "all even numbers",
			input: []int{2, 4, 6, 8, 10},
			predicate: func(value, index int, arr []int) bool {
				return value%2 == 0
			},
			expected: true,
		},
		{
			name:  "some even numbers",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(value, index int, arr []int) bool {
				return value%2 == 0
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Every(testCase.input, testCase.predicate)
			if result != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestFill(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arr      []interface{}
		value    interface{}
		start    int
		end      int
		expected []interface{}
	}{
		{
			name:     "normal case",
			arr:      []interface{}{1, 2, 3, 4, 5},
			value:    "x",
			start:    1,
			end:      4,
			expected: []interface{}{1, "x", "x", "x", 5},
		},
		{
			name:     "empty array",
			arr:      []interface{}{},
			value:    "x",
			start:    0,
			end:      0,
			expected: []interface{}{},
		},
		{
			name:     "start and end out of range",
			arr:      []interface{}{1, 2, 3, 4, 5},
			value:    "x",
			start:    -1,
			end:      10,
			expected: []interface{}{"x", "x", "x", "x", "x"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Fill(tc.arr, tc.value, tc.start, tc.end)
			if !reflect.DeepEqual(tc.arr, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, tc.arr)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int, int, interface{}) bool
		arg       interface{}
		expected  []int
	}{
		{
			name:  "even numbers",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			predicate: func(i, value int, arg interface{}) bool {
				return value%2 == 0
			},
			arg:      nil,
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:  "numbers greater than threshold",
			input: []int{5, 10, 15, 20, 25},
			predicate: func(i, value int, arg interface{}) bool {
				threshold := arg.(int)
				return value > threshold
			},
			arg:      15,
			expected: []int{20, 25},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter(testCase.input, testCase.predicate, testCase.arg)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		arr         []interface{}
		predicate   func(interface{}) bool
		expectedIdx int
	}{
		{
			name: "normal case",
			arr:  []interface{}{1, 2, 3, 4, 5},
			predicate: func(val interface{}) bool {
				return val.(int) > 3
			},
			expectedIdx: 3,
		},
		{
			name: "no matching element",
			arr:  []interface{}{1, 2, 3, 4, 5},
			predicate: func(val interface{}) bool {
				return val.(int) > 10
			},
			expectedIdx: -1,
		},
		{
			name: "empty array",
			arr:  []interface{}{},
			predicate: func(val interface{}) bool {
				return val.(int) == 0
			},
			expectedIdx: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index := FindIndex(tc.arr, tc.predicate)
			if index != tc.expectedIdx {
				t.Errorf("expected index %d but got %d", tc.expectedIdx, index)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arr      []interface{}
		expected []interface{}
	}{
		{
			name:     "normal case",
			arr:      []interface{}{1, []interface{}{2, []interface{}{3, 4}}, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "flattened array is empty",
			arr:      []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "nested arrays with different types",
			arr:      []interface{}{1, []interface{}{2, "hello"}, true, []interface{}{3.14}},
			expected: []interface{}{1, 2, "hello", true, 3.14},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flattened := Flatten(tc.arr)
			if !reflect.DeepEqual(flattened, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, flattened)
			}
		})
	}
}

func TestIncludes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		value    int
		expected bool
	}{
		{
			name:     "value found",
			input:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: true,
		},
		{
			name:     "value not found",
			input:    []int{1, 2, 3, 4, 5},
			value:    6,
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Includes(testCase.input, testCase.value)
			if result != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		arrays   [][]interface{}
		expected []interface{}
	}{
		{
			name: "Test Intersection with Integers",
			arrays: [][]interface{}{
				{1, 2, 3, 4, 5},
				{4, 5, 6, 7},
				{2, 3, 5, 6},
			},
			expected: []interface{}{5},
		},
		{
			name: "Test Intersection with Strings",
			arrays: [][]interface{}{
				{"apple", "banana", "cherry"},
				{"cherry", "date", "elderberry"},
				{"banana", "cherry", "fig"},
			},
			expected: []interface{}{"cherry"},
		},
		{
			name: "Test Intersection with Empty Array",
			arrays: [][]interface{}{
				{1, 2, 3},
				{},
				{4, 5, 6},
			},
			expected: []interface{}{},
		},
		{
			name: "Test Intersection with No Common Elements",
			arrays: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []interface{}{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Intersection(testCase.arrays...)
			if !interfaceSliceEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v. %T %T", testCase.expected, result, testCase.expected, result)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []interface{}
		separator string
		expected  string
	}{
		{
			name:      "empty array",
			input:     []interface{}{},
			separator: "-",
			expected:  "",
		},
		{
			name:      "array with one element",
			input:     []interface{}{"hello"},
			separator: "-",
			expected:  "hello",
		},
		{
			name:      "array with multiple elements",
			input:     []interface{}{1, "two", true},
			separator: ",",
			expected:  "1,two,true",
		},
		{
			name:      "array with mixed types",
			input:     []interface{}{[]int{1, 2}, "three", map[string]string{"four": "five"}},
			separator: "-",
			expected:  "[1 2]-three-map[four:five]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Join(testCase.input, testCase.separator)
			if result != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		transform func(int, int, interface{}) string
		arg       interface{}
		expected  []string
	}{
		{
			name:  "empty slice",
			input: []int{},
			transform: func(index int, elem int, arg interface{}) string {
				return strconv.Itoa(index) + ":" + strconv.Itoa(elem) + ":" + arg.(string)
			},
			arg:      "optional",
			expected: []string{},
		},
		{
			name:  "slice with values",
			input: []int{1, 2, 3, 4, 5},
			transform: func(index int, elem int, arg interface{}) string {
				return strconv.Itoa(index) + ":" + strconv.Itoa(elem) + ":" + arg.(string)
			},
			arg:      "optional",
			expected: []string{"0:1:optional", "1:2:optional", "2:3:optional", "3:4:optional", "4:5:optional"},
		},
		{
			name:  "slice with negative values",
			input: []int{-2, -1, 0, 1, 2},
			transform: func(index int, elem int, arg interface{}) string {
				return strconv.Itoa(index*elem) + ":" + strconv.Itoa(elem) + ":" + arg.(string)
			},
			arg:      "extra",
			expected: []string{"0:-2:extra", "-1:-1:extra", "0:0:extra", "3:1:extra", "8:2:extra"},
		},
		{
			name:  "slice with large values",
			input: []int{100, 200, 300, 400, 500},
			transform: func(index int, elem int, arg interface{}) string {
				return strconv.Itoa(index+elem) + ":" + strconv.Itoa(elem) + ":" + arg.(string)
			},
			arg:      "extra",
			expected: []string{"100:100:extra", "201:200:extra", "302:300:extra", "403:400:extra", "504:500:extra"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Map(testCase.input, testCase.transform, testCase.arg)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestPull(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []interface{}
		values   []interface{}
		expected []interface{}
	}{
		{
			name:     "remove single value",
			input:    []interface{}{1, 2, 3, 4, 5},
			values:   []interface{}{2},
			expected: []interface{}{1, 3, 4, 5},
		},
		{
			name:     "remove multiple values",
			input:    []interface{}{1, 2, 3, 4, 5},
			values:   []interface{}{1, 3, 5},
			expected: []interface{}{2, 4},
		},
		{
			name:     "remove no values",
			input:    []interface{}{1, 2, 3, 4, 5},
			values:   []interface{}{},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Pull(testCase.input, testCase.values...)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		input        []int
		transform    func(int, int, int, interface{}) int
		initialValue int
		arg          interface{}
		expected     int
	}{
		{
			name:  "sum of numbers",
			input: []int{1, 2, 3, 4, 5},
			transform: func(accumulator, value, index int, arg interface{}) int {
				return accumulator + value
			},
			initialValue: 0,
			arg:          nil,
			expected:     15,
		},
		{
			name:  "product of numbers",
			input: []int{1, 2, 3, 4, 5},
			transform: func(accumulator, value, index int, arg interface{}) int {
				return accumulator * value
			},
			initialValue: 1,
			arg:          nil,
			expected:     120,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Reduce(testCase.input, testCase.transform, testCase.initialValue, testCase.arg)
			if result != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		input       []interface{}
		predicate   func(interface{}, int, []interface{}) bool
		expected    []interface{}
		removedSize int
	}{
		{
			name:  "remove even numbers",
			input: []interface{}{1, 2, 3, 4, 5},
			predicate: func(value interface{}, index int, arr []interface{}) bool {
				return value.(int)%2 == 0
			},
			expected:    []interface{}{1, 3, 5},
			removedSize: 2,
		},
		{
			name:  "remove strings starting with 'a'",
			input: []interface{}{"apple", "banana", "avocado", "orange"},
			predicate: func(value interface{}, index int, arr []interface{}) bool {
				return strings.HasPrefix(value.(string), "a")
			},
			expected:    []interface{}{"banana", "orange"},
			removedSize: 2,
		},
		{
			name:  "remove all elements",
			input: []interface{}{1, 2, 3, 4, 5},
			predicate: func(value interface{}, index int, arr []interface{}) bool {
				return true
			},
			expected:    []interface{}{},
			removedSize: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			removed := Remove(&testCase.input, testCase.predicate)

			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, testCase.input)
			}

			if len(removed) != testCase.removedSize {
				t.Errorf("expected removed size %d but got %d", testCase.removedSize, len(removed))
			}
		})
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "empty slice",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "slice with odd number of elements",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{5, 4, 3, 2, 1},
		},
		{
			name:     "slice with even number of elements",
			input:    []interface{}{"a", "b", "c", "d"},
			expected: []interface{}{"d", "c", "b", "a"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			Reverse(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, testCase.input)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "slice with unique values",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "slice with duplicate values",
			input: []int{1, 2, 3, 2, 4, 1, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input := make([]int, len(testCase.input))
			copy(input, testCase.input)

			Shuffle(input)

			// Ensure the shuffled array has the same elements as the input (but in a different order)
			if len(input) != len(testCase.input) {
				t.Fatalf("expected length %d, but got %d", len(testCase.input), len(input))
			}

			// Count the occurrence of each element in the input and shuffled arrays
			inputCount := make(map[int]int)
			shuffledCount := make(map[int]int)
			for _, elem := range testCase.input {
				inputCount[elem]++
			}
			for _, elem := range input {
				shuffledCount[elem]++
				if shuffledCount[elem] > inputCount[elem] {
					t.Errorf("duplicate element %d in shuffled array", elem)
				}
			}

			// Ensure that all elements in the input array are present in the shuffled array
			for elem, count := range inputCount {
				if shuffledCount[elem] != count {
					t.Errorf("element %d missing or duplicated in shuffled array", elem)
				}
			}
		})
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []interface{}
		start    int
		end      int
		expected []interface{}
	}{
		{
			name:     "empty slice",
			input:    []interface{}{},
			start:    0,
			end:      0,
			expected: []interface{}{},
		},
		{
			name:     "slice with default bounds",
			input:    []interface{}{1, 2, 3, 4, 5},
			start:    0,
			end:      len([]interface{}{1, 2, 3, 4, 5}),
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "slice with specified bounds",
			input:    []interface{}{"a", "b", "c", "d", "e"},
			start:    1,
			end:      4,
			expected: []interface{}{"b", "c", "d"},
		},
		{
			name:     "slice with out-of-bounds indices",
			input:    []interface{}{1, 2, 3, 4, 5},
			start:    -2,
			end:      10,
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "slice with invalid indices",
			input:    []interface{}{1, 2, 3, 4, 5},
			start:    3,
			end:      2,
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Slice(testCase.input, testCase.start, testCase.end)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestSome(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		predicate func(int, int, []int) bool
		expected  bool
	}{
		{
			name:  "some even numbers",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(value, index int, arr []int) bool {
				return value%2 == 0
			},
			expected: true,
		},
		{
			name:  "no even numbers",
			input: []int{1, 3, 5, 7, 9},
			predicate: func(value, index int, arr []int) bool {
				return value%2 == 0
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Some(testCase.input, testCase.predicate)
			if result != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    [][]interface{}
		expected []interface{}
	}{
		{
			name:     "empty slices",
			input:    [][]interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "slices with unique values",
			input:    [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "slices with duplicate values",
			input:    [][]interface{}{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var inputSlices [][]interface{}
			inputSlices = append(inputSlices, testCase.input...)

			result := Union(inputSlices...)
			sortSlice(result)            // Sort the result for comparison
			sortSlice(testCase.expected) // Sort the expected slice for comparison

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "empty slice",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "slice with unique values",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "slice with duplicate values",
			input:    []interface{}{1, 2, 3, 2, 4, 1, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Uniq(testCase.input)
			if !interfaceSliceEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestZip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		slice1   []interface{}
		slice2   []interface{}
		expected [][]interface{}
	}{
		{
			name:     "empty slices",
			slice1:   []interface{}{},
			slice2:   []interface{}{},
			expected: [][]interface{}{},
		},
		{
			name:     "slices of equal length",
			slice1:   []interface{}{1, 2, 3},
			slice2:   []interface{}{"a", "b", "c"},
			expected: [][]interface{}{{1, "a"}, {2, "b"}, {3, "c"}},
		},
		{
			name:     "slice1 is longer than slice2",
			slice1:   []interface{}{1, 2, 3, 4},
			slice2:   []interface{}{"a", "b", "c"},
			expected: [][]interface{}{{1, "a"}, {2, "b"}, {3, "c"}},
		},
		{
			name:     "slice2 is longer than slice1",
			slice1:   []interface{}{1, 2, 3},
			slice2:   []interface{}{"a", "b", "c", "d"},
			expected: [][]interface{}{{1, "a"}, {2, "b"}, {3, "c"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Zip(testCase.slice1, testCase.slice2)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}
