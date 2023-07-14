package ut

// Get returns the value if found, otherwise returns the zero value of the type
func GetOrDefault[T comparable, V any](m map[T]V, key T, defaultValue V) V {
	value, ok := m[key]
	if !ok {
		return defaultValue
	}
	return value
}

// Keys returns a slice containing all the keys from the given map.
func Keys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// MapKeys applies the given function to each key-value pair in the map and returns a slice
// containing the results of the function applied to each key.
func MapKeys[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R {
	keys := make([]R, 0, len(m))
	for k, v := range m {
		keys = append(keys, fn(k, v))
	}
	return keys
}

// MapValues applies the given function to each key-value pair in the map and returns a slice
// containing the results of the function applied to each value.
func MapValues[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R {
	values := make([]R, 0, len(m))
	for k, v := range m {
		values = append(values, fn(k, v))
	}
	return values
}

// Merge merges multiple maps into a single map.
// If duplicate keys are encountered, the value from the last map in the input order is used.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Omit creates a new map without the specified keys and their corresponding values from the input map.
func Omit[K comparable, V any](input map[K]V, keys ...K) map[K]V {
	omitKeys := make(map[K]struct{})
	for _, key := range keys {
		omitKeys[key] = struct{}{}
	}

	result := make(map[K]V)
	for key, value := range input {
		if _, omit := omitKeys[key]; !omit {
			result[key] = value
		}
	}

	return result
}

// Pick creates a new map with the specified keys and their corresponding values from the input map.
func Pick[K comparable, V any](input map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for _, key := range keys {
		if value, ok := input[key]; ok {
			result[key] = value
		}
	}
	return result
}

// Values returns a slice containing all the values from the given map.
func Values[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
