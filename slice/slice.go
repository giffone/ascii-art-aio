package slice

import st "ascii-art-allinone/structure"

// AddToRune
func AddToRune(abc *st.Art, slice []rune, index int) {
	abc.Text.Rune = append(abc.Text.Rune[:index], append(slice, abc.Text.Rune[index:]...)...)
}

// RemoveRune removing element from slice
func RemoveRune(slice []rune, i int) []rune {
	return append(slice[:i], slice[i+1:]...)
}

// AddToFinal
func AddToFinal(abc *st.Art, slice []string, index int) {
	abc.Output.Final = append(abc.Output.Final[:index], append(slice, abc.Output.Final[index:]...)...)
}
