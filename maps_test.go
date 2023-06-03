package ut

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "map with keys",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Keys(testCase.input)
			sort.Strings(result)
			sort.Strings(testCase.expected)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) string
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(k string, v int) string { return k },
			expected: []string{},
		},
		{
			name: "map with keys of different types",
			input: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			fn:       func(k string, v int) string { return k + "-key" },
			expected: []string{"one-key", "two-key", "three-key"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := MapKeys(testCase.input, testCase.fn)
			sort.Strings(result)            // Sort the result for comparison
			sort.Strings(testCase.expected) // Sort the expected slice for comparison
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) string
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(k string, v int) string { return strconv.Itoa(v) },
			expected: []string{},
		},
		{
			name: "map with values of different types",
			input: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			fn:       func(k string, v int) string { return strconv.Itoa(v) + "-value" },
			expected: []string{"1-value", "2-value", "3-value"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := MapValues(testCase.input, testCase.fn)
			sort.Strings(result)            // Sort the result for comparison
			sort.Strings(testCase.expected) // Sort the expected slice for comparison
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		maps     []map[string]int
		expected map[string]int
	}{
		{
			name:     "empty maps",
			maps:     []map[string]int{},
			expected: map[string]int{},
		},
		{
			name: "maps with non-overlapping keys",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"c": 3, "d": 4},
				{"e": 5, "f": 6},
			},
			expected: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
				"f": 6,
			},
		},
		{
			name: "maps with overlapping keys",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
				{"c": 5, "d": 6},
			},
			expected: map[string]int{
				"a": 1,
				"b": 3,
				"c": 5,
				"d": 6,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Merge(testCase.maps...)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestOmit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		keys     []string
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			keys:     []string{"a", "b"},
			expected: map[string]int{},
		},
		{
			name: "non-empty map with omitted keys",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			keys: []string{"a", "c"},
			expected: map[string]int{
				"b": 2,
			},
		},
		{
			name: "non-empty map with non-existing keys",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			keys: []string{"d", "e"},
			expected: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Omit(testCase.input, testCase.keys...)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestPick(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		keys     []string
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			keys:     []string{"a", "b"},
			expected: map[string]int{},
		},
		{
			name: "non-empty map with selected keys",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			keys: []string{"a", "c"},
			expected: map[string]int{
				"a": 1,
				"c": 3,
			},
		},
		{
			name: "non-empty map with non-existing keys",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			keys:     []string{"d", "e"},
			expected: map[string]int{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Pick(testCase.input, testCase.keys...)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestValues(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "map with values",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Values(testCase.input)
			sort.Ints(result)
			sort.Ints(testCase.expected)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected %v but got %v", testCase.expected, result)
			}
		})
	}
}
