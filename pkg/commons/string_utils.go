package commons

// IsEmpty checks for empty string
func IsEmpty(str string) bool {
	if len(str) <= 0 {
		return true
	}
	return false
}
