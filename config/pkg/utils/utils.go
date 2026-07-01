package utils

func SafeCompareString(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte = 0
	for i := range a {
		result |= a[i] ^ b[i]
	}

	return result == 0
}
