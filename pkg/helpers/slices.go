package helpers

// ContainsString checks if a slice contains a provided value and returns true if it does, false if it doesn't
func ContainsString(slice []string, value string) bool {
	for _, i := range slice {
		if value == i {
			return true
		}
	}

	return false
}
