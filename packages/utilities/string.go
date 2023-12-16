package utilities

func CollectDigitsFromString(s []byte) string {
	j := 0
	for _, b := range s {
		if ('0' <= b && b <= '9') ||
			// ('a' <= b && b <= 'z') ||
			// ('A' <= b && b <= 'Z') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func CollectCharacterFromString(s []byte, character byte) string {
	j := 0
	for _, b := range s {
		if b == character {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}
