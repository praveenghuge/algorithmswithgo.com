package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if len(list) < 0 || list == nil {
		return false
	}
	for _, number := range list {
		if number == num {
			return true
		}
	}
	return false
}
