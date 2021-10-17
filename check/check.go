// package check contains all functions related to data validation
package check

import (
	ex "ascii-art-allinone/exit"
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
	"strings"
)

// AsciiSymbols checks for illegal symbols in arguments
func AsciiSymbols(abc *st.Art) {
	for _, text := range abc.Args {
		for _, symbol := range text {
			if symbol == '\n' {
				continue
			}
			// searches symbol between " "(space) and "~"(tilda) in ASCII table, in another case exits
			if symbol < ' ' || symbol > '~' {
				ex.Exit(abc, 0, true, "Ascii: not correct symbols in arguments!")
			}
		}
	}
}

// EmptyArgs checks for "" arguments
func EmptyArgs(abc *st.Art) {
	argsTemp := []string{}
	for _, arg := range abc.Args {
		if len(arg) > 0 {
			argsTemp = append(argsTemp, arg)
		}
	}
	abc.Args = argsTemp
	abc.LenArgs = len(abc.Args)
}

// FileNameOutput checks the incoming name of the the specified file. Name has to suffix ".txt"
func FileNameOutput(abc *st.Art) {
	if abc.Flag.Output != "" {
		if !strings.HasSuffix(abc.Flag.Output, ".txt") {
			ex.Exit(abc, 0, true, "Output: can not create a file \""+abc.Flag.Output+"\". Try to change extention to .txt")
		}
	}
}

// Brackets checks if value has two bracket opened and closed [...] or {...}
func Brackets(value string, lValue int, brackets ...string) bool {
	count := 0
	if lValue > 1 {
		for _, bracket := range brackets {
			if bracket[0] == value[0] && bracket[1] == value[lValue-1] {
				count++
			}
		}
	}
	return count != 0
}

// ByteConsist checks if the value is consist in []byte
func ByteConsist(value byte, symbols ...byte) bool {
	count := 0
	for _, symbol := range symbols {
		if symbol == value {
			count++
		}
	}
	return count != 0
}

// TextInTSize checks text in terminal size if it big or not
func TextInTSize(abc *st.Art, width int, symbol rune) (add bool, lSymbol int) {
	lSymbol = widthSymbolArt(abc, symbol)
	if width+lSymbol > abc.T.Width {
		return false, 0
	}
	return true, lSymbol
}

// Maps checks for missing ascii symbols in created maps
func Maps(abc *st.Art) {
	alphabet := []rune{' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		+'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		+':', ';', '<', '=', '>', '?', '@',
		+'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		+'[', '\\', ']', '^', '_', '`',
		+'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		+'{', '|', '}', '~'}
	ok, stop := true, false
	for _, symbol := range alphabet {
		_, ok = abc.Alphabet.Rune[symbol]
		if !ok {
			// write all missing symbols
			wr.Alphabet(abc, symbol)
			stop = true
		}
	}
	if stop {
		ex.Exit(abc, 0, true)
	}
}

// Lines checks if line (text) empty or not
func Lines(abc *st.Art) (count int) {
	for _, line := range abc.Text.Width {
		if line == 0 {
			count++
		} else {
			count += 8
		}
	}
	return count
}

// widthSymbolArt checks width of art
func widthSymbolArt(abc *st.Art, key rune) (width int) {
	for _, line := range abc.Alphabet.Rune[key] {
		lLine := len(line)
		if lLine > width {
			width = lLine
		}
	}
	return width
}

func Separator(sep string) bool {
	switch sep {
	case ":":
		return true
	case ";":
		return true
	case "+":
		return true
	case "-":
		return true
	case ".":
		return true
	default:
		return false
	}
}
