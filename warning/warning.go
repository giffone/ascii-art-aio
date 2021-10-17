package warning

import (
	st "ascii-art-allinone/structure"
	"os"
	"strconv"
)

func Warning(abc *st.Art) {
	if abc.WarningMessage != "" {
		os.Stdout.WriteString(abc.WarningMessage)
	}
}

func ColorWarning(abc *st.Art, value string) {
	bracks := 2
	for _, letter := range value {
		if letter == '[' || letter == ']' {
			bracks--
		}
	}
	if bracks < 2 {
		abc.WarningMessage += "Color: \"" + value + "\" parameter is wrong. Try -h for help.\n"
	}
}

func FlagWarning(abc *st.Art, count int) {
	if count > 0 {
		abc.WarningMessage += "Flags: " + strconv.Itoa(count) + " wrong flag. Try -h for help.\n"
	}
}

func ColorIndexesWarning(abc *st.Art, count int) {
	if count > 0 {
		abc.WarningMessage += "Color: " + strconv.Itoa(count) + " wrong indexes skipped. Try -h for help.\n"
	}
}

func TooManyArgsWarning(abc *st.Art) {
	abc.WarningMessage += "Args: too many arguments founded:\n"
	for _, arg := range abc.Args {
		abc.WarningMessage += "- \"" + arg + "\"\n"
	}
	for _, arg := range abc.TrashArgs {
		abc.WarningMessage += "- \"" + arg + "\"\n"
	}
	abc.WarningMessage += "Try -h for help.\n"
}

func Alphabet(abc *st.Art, symbol rune) {
	abc.WarningMessage += "Alphabet: symbol \"" + string(symbol) + "\" is missing. Can not to continue.\n"
}

func Reverse(abc *st.Art) {
	abc.WarningMessage += "Reverse: some letters are unreadable. Marked [...].\n"
}
