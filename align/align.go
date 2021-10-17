package align

import (
	ex "ascii-art-allinone/exit"
	sz "ascii-art-allinone/size"
	sli "ascii-art-allinone/slice"
	st "ascii-art-allinone/structure"
)

// JustifyRemoveSpacesStart delete spaces at start from newline
func JustifyRemoveSpacesStart(abc *st.Art) {
	lRune := len(abc.Text.Rune)
	start := false
	for i := 0; i < lRune; i++ {
		symbol := abc.Text.Rune[i]
		if symbol != ' ' {
			start = true
		}
		if !start {
			abc.Text.Rune = sli.RemoveRune(abc.Text.Rune, i)
			abc.Text.Count--
			lRune = len(abc.Text.Rune)
			i--
			continue
		}
		if symbol == '\n' {
			start = false
		}
	}
}

// JustifyRemoveSpacesEnd delete spaces at end of newline
func JustifyRemoveSpacesEnd(abc *st.Art) {
	lRune := len(abc.Text.Rune)
	start := false
	for i := lRune - 1; i >= 0; i-- {
		symbol := abc.Text.Rune[i]
		if symbol != ' ' {
			start = true
		}
		if !start {
			abc.Text.Rune = sli.RemoveRune(abc.Text.Rune, i)
			abc.Text.Count--
			continue
		}
		if symbol == '\n' {
			start = false
		}
	}
}

func Align(abc *st.Art) {
	line := 0 // lines with text
	for i, width := range abc.Text.Width {
		if width != 0 {
			spaces := abc.Text.Spaces[i]
			remain := abc.T.Width - width
			alignMethod(abc, remain, spaces, line)
			line++
		}
	}
}

// Align
func alignMethod(abc *st.Art, remain, spaces, line int) {
	if remain > 0 {
		switch abc.Flag.Align {
		case "left":
			return
		case "center":
			addLeftSpace(abc, remain/2, line)
		case "right":
			addLeftSpace(abc, remain, line)
		case "justify":
			if spaces > 0 {
				addMidSpace(abc, remain, spaces, line)
			}
		default:
			ex.Exit(abc, 0, true, "Align: can not recognize position. Try to use:\n- left\n- right\n- center\n- justify")
		}
		sz.ResizeByT(abc, false)
	}
}

func addLeftSpace(abc *st.Art, remain, index int) {
	pieces := remain / abc.Alphabet.Space
	slice := make([]rune, pieces)
	tail := remain - (pieces * abc.Alphabet.Space)
	count := 0

	makeSpiceRuneSlice(&slice)
	addTab(&slice, tail)

	for i, symbol := range abc.Text.Rune {
		if index == count {
			sli.AddToRune(abc, slice, i)
			break
		}
		if symbol == '\n' {
			count++
		}
	}
}

func addMidSpace(abc *st.Art, remain, spaces, line int) {
	count, plusOne := 0, false

	pieces := remain / spaces / abc.Alphabet.Space
	slice := make([]rune, pieces)
	sliceExtra := make([]rune, pieces)

	allTail := (remain - (pieces * spaces * abc.Alphabet.Space))
	tail := allTail / spaces
	tailExtra := allTail - (tail * spaces)

	makeSpiceRuneSlice(&slice)
	addTab(&slice, tail)

	if tailExtra > 0 {
		plusOne = true
		// copy(sliceExtra, *slice)
		makeSpiceRuneSlice(&sliceExtra)
		addTab(&sliceExtra, tail)
	}

	lRune := len(abc.Text.Rune)
	for i := 0; i < lRune; i++ {
		symbol := abc.Text.Rune[i]
		if symbol == '\n' {
			count++
		}
		if line == count {
			if symbol == ' ' {
				if plusOne {
					count := tailExtra
					for count != 0 {
						sliceExtra = append(sliceExtra, '\t')
						count--
					}
					plusOne = false
					sli.AddToRune(abc, sliceExtra, i)
					lRune = len(abc.Text.Rune)
					i += pieces + tail + tailExtra
					continue
				}
				sli.AddToRune(abc, slice, i)
				lRune = len(abc.Text.Rune)
				i += pieces + tail
			}
		} else if line < count {
			break
		}
	}
}

func makeSpiceRuneSlice(slice *[]rune) {
	for i := range *slice {
		(*slice)[i] = ' '
	}
}

func addTab(slice *[]rune, space int) {
	for space != 0 {
		*slice = append(*slice, '\t')
		space--
	}
}

func Space(abc *st.Art) {
	abc.Alphabet.Rune['\t'] = []string{" ", " ", " ", " ", " ", " ", " ", " ", ""}
}
