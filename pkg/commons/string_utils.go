package commons

func IsEmpty(str string) bool {
	if len(str) <= 0 {
		return true
	}
	return false
}
