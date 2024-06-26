package utils

func IsStringInArray(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}