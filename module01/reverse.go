package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var newword string
	for i := len(word) - 1; i >= 0; i-- {
		newword += string(word[i])
	}
	return newword
}
