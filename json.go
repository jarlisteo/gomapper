package gomapper

func jsonIsArray(input string) bool {
	if input[0:1] == "[" {
		return true
	} else {
		return false
	}
}
