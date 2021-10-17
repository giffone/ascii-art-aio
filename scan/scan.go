package scan

import (
	ex "ascii-art-allinone/exit"
	st "ascii-art-allinone/structure"
	"bufio"
	"os"
	"strconv"
)

func ScanText(abc *st.Art, name string) {
	// open file.txt selected in arguments (flag)
	file, err := os.Open(name)
	if err != nil {
		ex.Exit(abc, 0, true, "Scan: can not to open the file "+name)
	}
	defer file.Close()
	// scan file.txt for content
	sourse := bufio.NewScanner(file)

	// alphabet = nil
	if sourse.Err() != nil {
		ex.Exit(abc, 0, true, "Scan: can not to read/scan the file "+abc.Font.Value)
	}
	lines := 0
	for sourse.Scan() {
		abc.Alphabet.Sourse = append(abc.Alphabet.Sourse, sourse.Text())
		lines++
	}
	if abc.Flag.Reverse == "" {
		// if file does not consist 855 lines in fontfile
		linesText := strconv.Itoa(lines)
		if lines != 855 {
			ex.Exit(abc, 0, true, "Scan: something wrong with font file: "+linesText+" lines instead of 855. Good bye!")
		}
	}
}
