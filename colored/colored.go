// package colored contains all functions related to coloring
package colored

import (
	chk "ascii-art-allinone/check"
	cnv "ascii-art-allinone/convert"
	ex "ascii-art-allinone/exit"
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
	"strings"
)

// ColorLib consists color libruary
func ColorLib(name string) string {
	switch name {
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "purple":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "orange":
		return "\033[38;2;255;186;0m"
	default:
		return "\033[0m"
	}
}

func FindColorParam(abc *st.Art) {
	// by default method for coloring - "none"
	abc.Flag.Color.MethodColoring = "none"
	for i := 0; i < abc.LenArgs; i++ {
		arg := abc.Args[i]
		if arg == "--color" || arg == "--color=" {
			ex.Exit(abc, 0, true, "Color: not founded color parameters. Try -h for help.\n")
		} else if strings.HasPrefix(arg, "--color=") {
			// if color has parameters [...]
			if i+1 < abc.LenArgs {
				argNext := &abc.Args[i+1]
				if strings.HasPrefix(*argNext, "{") {
					// if founded parameters for coloring
					*argNext = colorParam(abc, *argNext, "{}")
					break
				} else if strings.HasPrefix(*argNext, "[") {
					// if founded parameters for coloring
					*argNext = colorParam(abc, *argNext, "[]")
					break
				}
			}
			// if no parameters - color all words
			abc.Flag.Color.MethodColoring = "all"
		}
	}
}

func colorParam(abc *st.Art, value string, bracket string) string {
	lValue := len(value)
	if chk.Brackets(value, lValue, "{}", "[]") {
		arrParam, param, addLast := []string{}, "", false
		if lValue == 2 {
			// if parameter like [],{} , color whole text
			abc.Flag.Color.MethodColoring = "all"
			value = ""
		} else if lValue == 3 && chk.ByteConsist(value[1], ':', '+', '-', '.', ',', ';') {
			// if parameter like (:) or [+] or (-)
			abc.Flag.Color.MethodColoring = "all"
			value = ""
		} else {
			if chk.ByteConsist(value[1], ':', '+', '-') {
				// if parameter like (:d...) or [:5...]
				if bracket == "{}" {
					// add first symbol of accepted ascii
					arrParam = append(arrParam, " ")
					abc.Flag.Color.MethodColoring = "and"
				} else {
					// add first "0" index
					arrParam = append(arrParam, "0")
					abc.Flag.Color.MethodColoring = "and"
				}
			} else if chk.ByteConsist(value[lValue-2], ':', '+', '-') {
				// if parameter like(...d:) or [...5:]
				addLast = true
			}
			for _, elem := range value {
				if elem == '}' || elem == ']' {
					if param != "" {
						arrParam = append(arrParam, param)
					}
					break
				} else if elem == '{' || elem == '[' {
					continue
				}
				if elem == ':' || elem == '-' || elem == '+' || elem == '.' || elem == ',' || elem == ';' {
					if param != "" {
						arrParam = append(arrParam, param)
					}
					param = ""
					if elem == ':' || elem == '-' || elem == '+' {
						abc.Flag.Color.MethodColoring = "and"
						continue
					}
					abc.Flag.Color.MethodColoring = "or"
					continue
				}
				if elem != 0 {
					param += string(elem)
				}
			}
			if addLast {
				if bracket == "{}" {
					// add last symbol of accepted ascii
					arrParam = append(arrParam, "~")
				} else {
					// count all letters in all arguments to make max possible range index to add Color.ByIndex.Range2
					abc.Flag.Color.ByIndex.MaxIndex = true
				}
			}
			// save range of letters and delete from args
			if bracket == "{}" {
				abc.Flag.Color.BySymbol.Range = arrParam
				abc.Flag.Color.MethodBy = "bySymbol"
			} else {
				err, count := "", 0
				abc.Flag.Color.ByIndex.Range, err, count = cnv.ArrAtoi(arrParam)
				if err != "" {
					wr.ColorIndexesWarning(abc, count)
				}
				abc.Flag.Color.MethodBy = "byIndex"
			}
			value = ""
			// correct method to "or" if parameters less than 2
			if len(arrParam) < 2 {
				abc.Flag.Color.MethodColoring = "or"
			}
		}
	} else {
		wr.ColorWarning(abc, value)
	}
	return value
}

func ChooseColor(abc *st.Art) {
	// start color
	abc.Flag.Color.Case1 = ColorLib(strings.ToLower(abc.Flag.Color.Value))
	// stop color
	abc.Flag.Color.Case2 = ColorLib("")
}

func AddStartColor(abc *st.Art) {
	for i := 0; i < len(abc.Output.Final); i++ {
		abc.Output.Final[i] += abc.Flag.Color.Case1
	}
}

func AddEndColor(abc *st.Art) {
	for i := 0; i < len(abc.Output.Final); i++ {
		abc.Output.Final[i] += abc.Flag.Color.Case2
	}
}
