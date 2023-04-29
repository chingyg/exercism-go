package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {

	if isValidIndex(slice, index) {
		return slice[index], true
	} else {
		return 0, false
	}

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if isValidIndex(slice, index) {
		slice[index] = value

	} else {
		slice = append(slice, value)
	}

	return slice

}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {

	if length < 1 {
		return []int{}
	} else {
		slice := []int{}

		for i := 0; i < length; i++ {
			slice = append(slice, value)

		}
		return slice
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isValidIndex(slice, index) {

		return append(slice[:index], slice[index+1:]...)

	}

	return slice
}

func isValidIndex(slice []int, index int) bool {

	return index > -1 && index < len(slice)
}
