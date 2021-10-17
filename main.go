// Ascii-art all in one (default, fs, color, output, justify, reverse)
package main

import (
	al "ascii-art-allinone/align"
	alp "ascii-art-allinone/alphabet"
	chk "ascii-art-allinone/check"
	cl "ascii-art-allinone/colored"
	ex "ascii-art-allinone/exit"
	fl "ascii-art-allinone/flag"
	fn "ascii-art-allinone/fonts"
	in "ascii-art-allinone/input"
	ou "ascii-art-allinone/output"
	rv "ascii-art-allinone/reverse"
	sz "ascii-art-allinone/size"
	sp "ascii-art-allinone/split"
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
	"os"
)

func main() {
	abc := &st.Art{}
	// os.Args = append(os.Args, " ", "--align=right")
	abc.Args = os.Args[1:]
	// fmt.Printf("%#v \n", abc.Args)
	abc.LenArgs = len(abc.Args)
	if abc.LenArgs == 0 {
		ex.Exit(abc, 1, true)
	} else if abc.LenArgs == 1 {
		if abc.Args[0] == "" {
			ex.Exit(abc, 1, true)
		}
	}
	chk.AsciiSymbols(abc)
	// calculate terminal window size
	// abc.T.Width = 300
	sz.TSize(abc)
	if abc.T.Width == 0 {
		ex.Exit(abc, 0, true, "Terminal size = 0.")
	}
	// fmt.Println(abc.T.Width)
	// find parameters of coloring if exist
	cl.FindColorParam(abc)
	// find and sort flags
	fl.SortFlag(abc)
	// add founded flags to structure Flags
	fl.ParseFlag(abc)
	// find wrong flags and make them empty-args
	fl.FindWrongFlag(abc)
	// find font in arguments
	if abc.Flag.Reverse == "" {
		if !abc.Font.Founded {
			if !fn.FontChoose(abc) {
				fn.FontDefault(abc)
			} else {
				abc.Font.Founded = true
			}
		}
		// read font file, check for errors
		alp.MakeAlphabet(abc)
		// checks for missing runes
		chk.Maps(abc)
		// remember width of rune " "
		alp.LenSpace(abc)
		// find empty-args and delete them
		chk.EmptyArgs(abc)
		// find unwanted name of file
		chk.FileNameOutput(abc)
		// lenght arguments without fonts
		if abc.LenArgs != 1 {
			wr.TooManyArgsWarning(abc)
			ex.Exit(abc, 1, true)
		} else {
			// check symbols, "\n" and append into slice
			sp.SplitArgs(abc)
			if abc.Flag.Align == "justify" || abc.Flag.Align == "right" {
				// if has any symbols exept space
				if abc.Text.Symbols {
					// remove spaces left/right from newline
					al.JustifyRemoveSpacesStart(abc)
					al.JustifyRemoveSpacesEnd(abc)
				}
			}
			al.Space(abc) // add thin space as '\t'
			// resize text by terminal size, if need add newline
			sz.ResizeByT(abc, true)
			if abc.Flag.Align != "" {
				al.Align(abc)
			}
			in.MakeArt(abc)
		}
	} else {
		// read font file, check for errors
		alp.MakeAlphabetBigData(abc)
		// remember width of rune " "
		alp.LenSpace(abc)
		rv.ScanFile(abc)
		rv.ScanArt(abc)
	}
	// save to file or print in terminal
	ou.HowToOutput(abc)
	ex.Exit(abc, 0, true)
}
