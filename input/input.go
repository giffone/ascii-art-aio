package input

import (
	// al "ascii-art-allinone/align"
	chk "ascii-art-allinone/check"
	cl "ascii-art-allinone/colored"
	st "ascii-art-allinone/structure"
)

// MakeArt beagin to make art
func MakeArt(abc *st.Art) {
	// choose color start / end (reset)
	cl.ChooseColor(abc)
	if abc.Flag.Color.MethodBy == "bySymbol" {
		if abc.Flag.Color.MethodColoring == "and" {
			// coloring in range between letter1 and letter2
			abc.Flag.Color.BySymbol.Range1, abc.Flag.Color.BySymbol.Range2 = findAndSymbol(abc)
		}
	} else if abc.Flag.Color.MethodBy == "byIndex" {
		// add max number for range if parameter was [5:]
		if abc.Flag.Color.ByIndex.MaxIndex {
			lRune := len(abc.Text.Rune)
			abc.Flag.Color.ByIndex.Range = append(abc.Flag.Color.ByIndex.Range, lRune)
			abc.Flag.Color.MethodColoring = "and"
		}
		if abc.Flag.Color.MethodColoring == "and" {
			// coloring in range between letter1 and letter2
			abc.Flag.Color.ByIndex.Range1, abc.Flag.Color.ByIndex.Range2 = findAndIndex(abc)
		}
	}
	// read letters from argument
	letters(abc)
	// add last newline if not exist
	// LastNewLine(abc)
}

// letters finds letters of the word in the standard.txt banner and prints it
func letters(abc *st.Art) {
	// count text blocks
	l := chk.Lines(abc)
	abc.Output.Final = make([]string, l)
	abc.Output.Index = 8
	// if method coloring all words, add color to the begining of the line
	if abc.Flag.Color.MethodColoring == "all" {
		cl.AddStartColor(abc)
	}

	index := 0
	for _, symbol := range abc.Text.Rune {
		if symbol == '\n' {
			if abc.Alphabet.Letter != 0 && abc.Alphabet.Letter != '\n' {
				abc.Output.Index += 9 // 8 lines text + 1 newline
			} else {
				abc.Output.Index += 1
			}
			abc.Alphabet.Letter = symbol
			continue
		}
		// remember current letter
		abc.Alphabet.Letter = symbol
		// by default coloring is false
		abc.Alphabet.Coloring = false
		// coloring for letter by letter
		if symbol != ' ' && symbol != '\t' {
			if abc.Flag.Color.MethodColoring == "or" {
				if abc.Flag.Color.MethodBy == "bySymbol" {
					if findOrSymbol(abc, symbol) {
						// coloring in range between one letter
						abc.Flag.Color.BySymbol.Range1, abc.Flag.Color.BySymbol.Range2 = symbol, symbol
					} else {
						// clear range for the next one
						abc.Flag.Color.BySymbol.Range1, abc.Flag.Color.BySymbol.Range2 = 0, 0
					}
				} else if abc.Flag.Color.MethodBy == "byIndex" {
					if findOrIndex(abc, index) {
						// coloring in range between one letter
						abc.Flag.Color.ByIndex.Range1, abc.Flag.Color.ByIndex.Range2 = index, index
					} else {
						// clear range for the next one
						abc.Flag.Color.ByIndex.Range1, abc.Flag.Color.ByIndex.Range2 = -1, -1
					}
				}
			}
			if abc.Flag.Color.MethodBy == "bySymbol" {
				// if choosen letter in range - make color
				if symbol >= abc.Flag.Color.BySymbol.Range1 && symbol <= abc.Flag.Color.BySymbol.Range2 {
					abc.Alphabet.Coloring = true
				}
			} else if abc.Flag.Color.MethodBy == "byIndex" {
				// if choosen letter in range - make color
				if index >= abc.Flag.Color.ByIndex.Range1 && index <= abc.Flag.Color.ByIndex.Range2 {
					abc.Alphabet.Coloring = true
				}
			}
			index++
		}
		makeOutputByMethod(abc)
	}
}

// makeOutputByMethod make output by method
func makeOutputByMethod(abc *st.Art) {
	if abc.Flag.Color.MethodColoring == "none" || abc.Flag.Color.MethodColoring == "all" {
		// if method "all" - it is no need to add color for each letter. only at start and the end line
		makeOutput(abc, false)
	} else {
		makeOutput(abc, abc.Alphabet.Coloring)
	}
}

// makeOutput append to array founded letter
func makeOutput(abc *st.Art, coloring bool) {
	index := 0
	for i := abc.Output.Index - 8; i < abc.Output.Index; i++ {
		line := abc.Alphabet.Rune[abc.Alphabet.Letter][index]
		if coloring {
			abc.Output.Final[i] += abc.Flag.Color.Case1 + line + abc.Flag.Color.Case2
		} else {
			abc.Output.Final[i] += line
		}
		index++
	}
}

// findAndSymbol find "a" and "f" letters in range (a-f)
func findAndSymbol(abc *st.Art) (symbol1 rune, symbol2 rune) {
	for _, word := range abc.Flag.Color.BySymbol.Range {
		for _, symbol := range word {
			if symbol1 == 0 {
				symbol1 = symbol
			}
			symbol2 = symbol
		}
	}
	// if wrong range, fix it
	if symbol1 > symbol2 {
		symbol1, symbol2 = symbol2, symbol1
	}
	return symbol1, symbol2
}

// findAndIndex find "0" and "8" indexes in range [0-8]
func findAndIndex(abc *st.Art) (index1 int, index2 int) {
	sourse := abc.Flag.Color.ByIndex.Range
	lSourse := len(sourse)

	index1 = sourse[0]
	index2 = sourse[lSourse-1]
	// if wrong range, fix it
	if index1 > index2 {
		index1, index2 = index2, index1
	}
	return index1, index2
}

// findOrSymbol find "a" or "f" letters in range (a,f)
func findOrSymbol(abc *st.Art, symbol rune) bool {
	for _, word := range abc.Flag.Color.BySymbol.Range {
		for _, letterF := range word {
			if symbol == letterF {
				return true
			}
		}
	}
	return false
}

// findOrIndex find "0" or "8" indexes in range (0,8)
func findOrIndex(abc *st.Art, index int) bool {
	for _, indexF := range abc.Flag.Color.ByIndex.Range {
		if index == indexF {
			return true
		}
	}
	return false
}
