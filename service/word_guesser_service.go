package service

func IsWordGuessed(arr []rune) bool {
	var result = true
	for _, r := range arr {
		if r == '_' {
			result = false
			break
		}
	}

	return result
}
