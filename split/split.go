package split

import (
	ex "ascii-art-allinone/exit"
	st "ascii-art-allinone/structure"
)

// SplitArgs checks for illegal symbols in arguments and split by new line
func SplitArgs(abc *st.Art) {
	if abc.Args[0] == "" {
		ex.Exit(abc, 1, true)
	} else if abc.Args[0] == "\\n" {
		ex.Exit(abc, 2, true)
	}

	makeNewLine := false
	for _, symbol := range abc.Args[0] {
		if symbol != ' ' {
			// if has any symbols exept space
			abc.Text.Symbols = true
		}
		if symbol == 'n' {
			if makeNewLine {
				// make split if was "\\" and now "n"
				abc.Text.Rune = append(abc.Text.Rune, '\n')
			} else {
				// don't make split if only catched "n", just add it
				abc.Text.Rune = append(abc.Text.Rune, 'n')
			}
			// split done or found "n" without "\\"
			makeNewLine = false
			continue
		}
		if makeNewLine {
			// if loop before found "\\" but not founded "n", just append "\\"
			abc.Text.Rune = append(abc.Text.Rune, '\\')
		}
		// stop seek to make newline
		makeNewLine = false
		if symbol == '\\' {
			// if found "\\" then next loop catch "n"
			makeNewLine = true
		} else {
			// if others symbols, just add them
			abc.Text.Rune = append(abc.Text.Rune, symbol)
			makeNewLine = false
		}
	}
	// add last symbol is "\\" (tail), just append it
	if makeNewLine {
		abc.Text.Rune = append(abc.Text.Rune, '\\')
	}
}
