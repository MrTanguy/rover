package utils

// FindIndex is a helper function to find the index of the current orientation in the directions slice
func FindIndex(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
