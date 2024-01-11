package common

func IsStringInArray(array []string, item string) bool {
	for _, obj := range array {
		if obj == item {
			return true
		}
	}
	return false
}
