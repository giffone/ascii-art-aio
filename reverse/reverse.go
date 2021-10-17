package reverse

import (
	alp "ascii-art-allinone/alphabet"
	cnv "ascii-art-allinone/convert"
	ex "ascii-art-allinone/exit"
	sc "ascii-art-allinone/scan"
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
)

func ScanFile(abc *st.Art) {
	s := &st.Scan{}

	sc.ScanText(abc, abc.Flag.Reverse)

	count, start := 0, false
	for _, line := range abc.Alphabet.Sourse {
		s.Slice = []byte(line)
		lSlice := len(s.Slice)
		// if slice = "\n",
		if lSlice == 0 {
			if count != 8 && start {
				ex.Exit(abc, 0, true, "Reverse: Art not 8 lines of hight. Only "+cnv.Itoa(count)+"lines.")
			}
			abc.Read.Width = append(abc.Read.Width, s.LLine, 0)
			s.LLine, count, start = 0, 0, false
			continue
		}
		abc.Read.File = append(abc.Read.File, s.Slice...)
		abc.Read.File = append(abc.Read.File, '\n')
		// remember
		if s.LLine == 0 {
			s.LLine = lSlice
		} else {
			if s.LLine != lSlice {
				ex.Exit(abc, 0, true, "Reverse: Art letter has different width. Can not read.")
			}
		}
		count++
		start = true
	}
}

// ScanArt make template
func ScanArt(abc *st.Art) {
	for _, lenght := range abc.Read.Width {
		if lenght == 0 {
			abc.Read.Answer = append(abc.Read.Answer, '\n')
			continue
		} else {
			width := (lenght * 8) + 7 // +7 newlines
			cutLetter(abc, width)
			makeMatrix(abc, lenght)
			readMatrix(abc, lenght)
			width++ // +1 newline at the end between letters
		}
	}
}

// cutLetter make slice of one letter
func cutLetter(abc *st.Art, width int) {
	lFile := len(abc.Read.File)
	if width <= lFile {
		abc.Read.Tmp = append([]byte(nil), abc.Read.File[:width]...)
		abc.Read.TmpMatrix = append([]byte(nil), abc.Read.Tmp...)
		abc.Read.File = append([]byte(nil), abc.Read.File[width+1:]...)
	} else {
		ex.Exit(abc, 0, true, "Reverse: text did not formed because lines less than 8.")
	}
}

// makeMatrix make template, delete art excepting borders
func makeMatrix(abc *st.Art, lenght int) {
	color := 0
	for index, b := range abc.Read.Tmp {
		if b == 0 {
			continue
		} else if b == 27 { // color is begining
			color++
		} else {
			if color == 1 {
				if b == 91 { // color is ready
					color++
					continue
				}
			} else if color == 2 {
				if b == 'm' { // color stop
					color = 0
				}
				continue
			} else if color == 0 { // if not color
				if b != '\n' && b != ' ' {
					clearLeft(abc, index, lenght+1, 1) // +1 is '\n'
					clearRight(abc, index, lenght+1, 1)
				}
			}
		}
	}
}

// clearLeft delete symbol and other symbols by vertical index down
func clearLeft(abc *st.Art, index, step int, value byte) {
	for i := index; i < len(abc.Read.TmpMatrix); i += step {
		abc.Read.TmpMatrix[i] = value
	}
}

// clearRight delete symbol and other symbols by vertical index up
func clearRight(abc *st.Art, index, step int, value byte) {
	for i := index; i >= 0; i -= step {
		abc.Read.TmpMatrix[i] = value
	}
}

// readMatrix read art into border
func readMatrix(abc *st.Art, lenght int) {
	s := &st.Scan{Start: true}
	color := 0

	for i, symbol := range abc.Read.TmpMatrix {
		if symbol == '\n' {
			break
		} else {
			if s.Start {
				if symbol == 1 {
					s.IndexStart, s.IndexEnd = i, i
					color = 0
				}
			}
			if symbol == 27 {
				color++
				continue
			}
			if color == 1 {
				if symbol == 91 {
					color++
					abc.Flag.Color.Case = append(abc.Flag.Color.Case, 27, 91)
					continue
				}
			} else if color == 2 {
				if symbol == 'm' {
					abc.Read.Answer = append(abc.Read.Answer, abc.Flag.Color.Case...)
					abc.Read.Answer = append(abc.Read.Answer, 'm')
					abc.Flag.Color.Case = nil
					color = 0
					continue
				}
				abc.Flag.Color.Case = append(abc.Flag.Color.Case, symbol)
				continue
			}
			if symbol == ' ' && s.SymbolB == 1 {
				// if current symbol is ' ' and previous is art
				s.IndexEnd = i
				findLetter(abc, s, lenght+1)
				s.Start = true
				s.SymbolB = symbol
				s.CountSpace = 0
				continue
			} else if symbol == ' ' && (s.SymbolB == ' ' || s.SymbolB == 0) {
				// if current symbol is ' ' and previous is ' '
				s.CountSpace++
			}
			if symbol == 1 {
				s.Start = false
			}
		}
		// remember current symbol
		s.SymbolB = symbol
		if s.CountSpace == abc.Alphabet.Space {
			// if spaces is equal art space, just add it
			abc.Read.Answer = append(abc.Read.Answer, ' ')
			s.CountSpace = 0
			s.Start = true
		}
	}
}

func findLetter(abc *st.Art, s *st.Scan, step int) {
	keys, keysIndex := 0, 0
	width := s.IndexEnd - s.IndexStart + 1 // +1 correct (zero - is also need to count)
	filter := abc.Alphabet.BigData[width]

	for i := s.IndexStart; i < len(abc.Read.Tmp); i += step {
		slice := abc.Read.Tmp[i : i+width]
		keys += alp.Key(slice, keysIndex+1)
		keysIndex++
	}
	if letter, ok := filter[keys]; ok {
		abc.Read.Answer = append(abc.Read.Answer, letter)
	} else {
		// if not founded map, add [...] and save warning
		wr.Reverse(abc)
		abc.Read.Answer = append(abc.Read.Answer, '[', '.', '.', '.', ']')
	}
}
