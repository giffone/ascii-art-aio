package exit

import (
	st "ascii-art-allinone/structure"
	wr "ascii-art-allinone/warning"
	"os"
)

// exit from program if it's done
func Exit(abc *st.Art, n int, exit bool, messages ...string) {
	for n != 0 {
		os.Stdout.WriteString("\n")
		n--
	}
	wr.Warning(abc)
	if len(messages) > 0 {
		for _, message := range messages {
			os.Stdout.WriteString(message + "\n")
		}
	}
	if exit {
		os.Exit(0)
	}
}
