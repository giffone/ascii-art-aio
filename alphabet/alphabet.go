package alphabet

import (
	cnv "ascii-art-allinone/convert"
	ex "ascii-art-allinone/exit"
	sc "ascii-art-allinone/scan"
	st "ascii-art-allinone/structure"
)

// MakeAlphabet makes map of ascii-art
func MakeAlphabet(abc *st.Art) {
	// first symbol -32 and width of each symbol - 9
	s := &st.Scan{Code: 32, Height: 9}

	abc.Alphabet.RuneLen = make(map[rune]int)
	abc.Alphabet.Rune = make(map[rune][]string)

	sc.ScanText(abc, abc.Font.Value)

	// read scanned bufer
	scanRune(abc, s, false)
	abc.Alphabet.Sourse = nil
	// if alphabet does not consist 95 symbols (32-126)
	lenght := len(abc.Alphabet.Rune)
	if lenght != 95 {
		ex.Exit(abc, 0, true, "Alphabet: there is no 95 letters imported, just: "+cnv.Itoa(lenght))
	}
}

// scanRune make map with key - rune
func scanRune(abc *st.Art, s *st.Scan, start bool) {
	for _, line := range abc.Alphabet.Sourse {
		s.LLine = len(line)
		if start {
			s.SymbolR = rune(s.Code)
			// save rune exept last line ""
			abc.Alphabet.Rune[s.SymbolR] = append(abc.Alphabet.Rune[s.SymbolR], line)
			if s.Height != s.Lines {
				// save lenght of each symbol, if not exist, checks for errors
				lenScanRune(abc, s)
				// go to next map
			} else {
				s.Code++
				s.Height += 9
			}
		}
		s.Lines++
		start = true
	}
}

// LenSpace remember width of space of current alphabet
func LenSpace(abc *st.Art) {
	if space, ok := abc.Alphabet.Rune[' ']; ok {
		for _, line := range space {
			lLine := len(line)
			if lLine > abc.Alphabet.Space {
				abc.Alphabet.Space = lLine
			}
		}
	} else {
		abc.Alphabet.Space = 6
	}
}

// lenScanRune saves lenght of symbol, if not exist, and compare for each line, if compares is different - error
func lenScanRune(abc *st.Art, s *st.Scan) {
	if _, ok := abc.Alphabet.RuneLen[s.SymbolR]; !ok {
		abc.Alphabet.RuneLen[s.SymbolR] = s.LLine
	} else {
		if s.LLine != abc.Alphabet.RuneLen[s.SymbolR] {
			ex.Exit(abc, 0, true, "Alphabet: length of symbol \""+string(s.SymbolR)+"\" is different.")
		}
	}
}
