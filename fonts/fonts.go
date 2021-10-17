package fonts

import (
	st "ascii-art-allinone/structure"
)

func FontChoose(abc *st.Art) bool {
	if abc.LenArgs > 1 {
		for i, arg := range abc.Args {
			if arg == "st" || arg == "standard" {
				abc.Font.Value = "fonts/standard.txt"
				abc.TrashArgs = append(abc.TrashArgs, arg)
				abc.Args[i] = ""
				return true
			} else if arg == "th" || arg == "thinkertoy" {
				abc.Font.Value = "fonts/thinkertoy.txt"
				abc.TrashArgs = append(abc.TrashArgs, arg)
				abc.Args[i] = ""
				return true
			} else if arg == "sh" || arg == "shadow" {
				abc.Font.Value = "fonts/shadow.txt"
				abc.TrashArgs = append(abc.TrashArgs, arg)
				abc.Args[i] = ""
				return true
			}
		}
		return false
	}
	return false
}

func FontDefault(abc *st.Art) {
	if abc.Font.Value == "" {
		abc.Font.Value = "fonts/standard.txt"
		abc.Font.Founded = true
	}
}

