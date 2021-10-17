package flag

import (
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
	"flag"
	"strings"
)

// move all flags to the left side for flag.parsing
func SortFlag(abc *st.Art) {
	lArgs := len(abc.Args)
	sorted := true
	// bubble sort
	for sorted {
		sorted = false
		for i := 0; i < lArgs-1; i++ {
			// if have flags - move to the left side for flag.parsing
			if !strings.HasPrefix(abc.Args[i], "--") && strings.HasPrefix(abc.Args[i+1], "--") {
				abc.Args[i], abc.Args[i+1] = abc.Args[i+1], abc.Args[i]
				sorted = true
			}
		}
	}
}

// save flags with parameters
func ParseFlag(abc *st.Art) {
	// if flag like "--" is empty, save to trash before parsing
	for _, letter := range abc.Args {
		if letter == "--" {
			abc.TrashArgs = append(abc.TrashArgs, letter)
		}
	}
	flag.StringVar(&abc.Flag.Output, "output", "", "only *txt format can be used, for example: --output=file1.txt\n")
	flag.StringVar(&abc.Flag.Color.Value, "color", "system color", "for example: --color=red\n"+
		"with parameters index: --color=red [] or [:10] or [5:9] or [2,15] or [4-6] or [2+9]\n"+
		"with parameters letters: --color=red {} or {:g} or {ac} or {a,b} or {a-e} or {a:f}\n"+
		"use colors: black, red, green, yellow, blue, purple, cyan, white\n")
	flag.StringVar(&abc.Flag.Reverse, "reverse", "", "only *txt format can be used, for example: --reverse=file1.txt\n")
	flag.StringVar(&abc.Flag.Align, "align", "", "for example: --align=center\n"+
		"use parameters: left, right, center, justify\n")
	// save flags to structure
	flag.Parse()
	// arguments without flag
	abc.Args = flag.Args()
	abc.LenArgs = len(abc.Args)
}

func FindWrongFlag(abc *st.Art) {
	count := 0
	for _, arg := range abc.Args {
		if strings.HasPrefix(arg, "-") {
			count++
		}
	}
	wr.FlagWarning(abc, count)
}
