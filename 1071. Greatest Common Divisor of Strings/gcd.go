package gcd

func gcdOfStrings(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	if len1 < len2 {
		len1, len2 = len2, len1
		str1, str2 = str2, str1
	}

	if str1[:len2] != str2 {
		return ""
	}

	gcd := gcdOfNumbers(len1, len2)

	for i := 0; i < len1; i++ {
		if str1[i] != str1[i%gcd] {
			return ""
		}
	}

	return str1[:gcd]
}

func gcdOfNumbers(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
