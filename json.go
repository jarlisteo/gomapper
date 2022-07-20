package gomapper

func jsonIsArray(json string) bool {
	if json[0:1] == "[" {
		return true
	} else {
		return false
	}
}
