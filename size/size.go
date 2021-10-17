// package mathem consists of all calculation functions
package size

import (
	chk "ascii-art-allinone/check"
	cnv "ascii-art-allinone/convert"
	ex "ascii-art-allinone/exit"
	sli "ascii-art-allinone/slice"
	st "ascii-art-allinone/structure"
	"log"
	"os"
	"os/exec"
)

// TSize get terminal size
func TSize(abc *st.Art) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	abc.T.Size = cnv.ArrByteToArrInt(out)

	if len(abc.T.Size) == 2 {
		abc.T.Height = abc.T.Size[0]
		abc.T.Width = abc.T.Size[1] - 2 // correction -1
	} else {
		ex.Exit(abc, 0, true, "Terminal: size can not calculate, height and width is :"+cnv.ArrByteToString(out))
	}
}

// ResizeByT resize text if lines bigger than terminal
func ResizeByT(abc *st.Art, make bool) {
	split := &st.Resize{}
	abc.Text.Width, abc.Text.Spaces = []int{}, []int{}
	lRune := len(abc.Text.Rune)

	if abc.Text.Rune[lRune-1] != '\n' {
		abc.Text.Rune = append(abc.Text.Rune, '\n')
		lRune++
	}

	for i := 0; i < lRune; i++ {
		symbol := abc.Text.Rune[i]
		if symbol == '\n' {
			makeNewLine(abc, split, false)
			abc.Text.Count++
			continue
		}
		// if text is smaller than terminal size, add letter
		if split.Add, split.LSymbol = chk.TextInTSize(abc, split.Width, symbol); split.Add {
			// width of art symbol
			split.Width += split.LSymbol
			if symbol == ' ' {
				split.Spaces++
			}
		} else {
			if make {
				// if text is bigger than terminal, add a newline
				makeNewLine(abc, split, true)
				lRune++
			}
		}
		abc.Text.Count++
	}
}

func makeNewLine(abc *st.Art, spl *st.Resize, add bool) {
	index := abc.Text.Count
	if add {
		sli.AddToRune(abc, []rune{'\n'}, index)
	}
	if spl.Width != 0 {
		abc.Text.Width = append(abc.Text.Width, spl.Width, 0)
		abc.Text.Spaces = append(abc.Text.Spaces, spl.Spaces, 0)
	} else {
		abc.Text.Width = append(abc.Text.Width, 0)
		abc.Text.Spaces = append(abc.Text.Spaces, 0)
	}

	spl.Width, spl.Spaces = 0, 0
}
