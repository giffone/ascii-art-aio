package alphabet

import (
	cnv "ascii-art-allinone/convert"
	ex "ascii-art-allinone/exit"
	sc "ascii-art-allinone/scan"
	st "ascii-art-allinone/structure"
)

// MakeAlphabetBigData makes map for reverse flag
func MakeAlphabetBigData(abc *st.Art) {
	s := &st.Scan{}
	abc.Alphabet.BigData = make(map[int]map[int]byte)
	abc.Alphabet.Fonts = append(abc.Alphabet.Fonts, "standard", "shadow", "thinkertoy")

	for _, font := range abc.Alphabet.Fonts {
		sc.ScanText(abc, "fonts/"+font+".txt")
		// read scanned bufer
		scanKey(abc, s, false)
		abc.Alphabet.Sourse = nil
		lenght := len(abc.Alphabet.RuneLen)
		if lenght != 95 {
			ex.Exit(abc, 0, true, "Alphabet: there is no 95 letters imported in "+font+", just: "+cnv.Itoa(lenght))
		}
		abc.Alphabet.RuneLen = nil
	}
}

// scanKey make map with key - uniq key
func scanKey(abc *st.Art, s *st.Scan, start bool) {
	abc.Alphabet.RuneLen = make(map[rune]int)
	// first symbol -32 and width of each symbol - 9
	s.Code, s.Height, s.Lines = 32, 9, 0
	keys, keysIndex := 0, 0
	for _, line := range abc.Alphabet.Sourse {
		if start {
			s.SymbolR = rune(s.Code)
			s.Slice = []byte(line)
			if s.Height != s.Lines {
				s.LLine = len(s.Slice)
				// save lenght of each symbol, if not exist, checks for errors
				lenScanRune(abc, s)
				// make key
				keys += Key(s.Slice, keysIndex+1)
				keysIndex++
			} else {
				// save map ang go to next creating
				add(abc, s, keys)
				s.Code++
				s.Height += 9
				keys, keysIndex, s.LLine = 0, 0, 0
			}
		}
		s.Lines++
		start = true
	}
	// save last symbol
	if keys > 0 && s.Code > 0 {
		add(abc, s, keys)
	}
}

// Key is generating uniq Key for map Key
func Key(slice []byte, line int) (value int) {
	for i, el := range slice {
		if el != ' ' {
			value += (value << 4) + int(el) + i + line
		} else {
			value += (value >> 4)
		}
	}
	return value
}

// add adding to map
func add(abc *st.Art, s *st.Scan, keys int) {
	if _, ok := abc.Alphabet.BigData[s.LLine]; !ok {
		// if there is no map with key of width of the art symbol, make it
		abc.Alphabet.BigData[s.LLine] = make(map[int]byte)
	}
	if symbol, ok := abc.Alphabet.BigData[s.LLine][keys]; ok {
		if symbol != ' ' {
			ex.Exit(abc, 0, true, "Alphabet: trying rewrite symbol: "+cnv.Itoa(keys)+" --> "+string(rune(s.Code)))
		} else {
			// remember width of space
			abc.Alphabet.Space = s.LLine
		}
	}
	abc.Alphabet.BigData[s.LLine][keys] = byte(s.SymbolR)
}
