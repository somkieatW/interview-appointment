package utils

// IsEmpty ..
func IsEmpty(value string) bool {
	return value == ""
}

// IsNotEmpty ..
func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}
