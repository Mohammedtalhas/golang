package palindrome

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
func VerifyPalindrome(str string) bool {
	if str == Reverse(str) {
		return true
	}
	return false
}
