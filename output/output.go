package output

import (
	ex "ascii-art-allinone/exit"
	st "ascii-art-allinone/structure"
	"os"
)

// final output to terminal
func printOutput(abc *st.Art) {
	if abc.Read.Answer != nil {
		os.Stdout.Write(abc.Read.Answer)
	} else {
		for _, line := range abc.Output.Final {
			os.Stdout.WriteString(line + "\n")
		}
	}
}

func fileOutput(abc *st.Art, file *os.File) {
	for _, line := range abc.Output.Final {
		if _, err := file.WriteString(line + "\n"); err != nil {
			ex.Exit(abc, 0, true, "Output: Can not write text to "+abc.Flag.Output)
		}
	}
}

func HowToOutput(abc *st.Art) {
	// fmt.Println(abc.Text.Rune)
	if abc.Flag.Output != "" {
		file, err := os.OpenFile(abc.Flag.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ex.Exit(abc, 0, true, "Output: Not correct file name to save.")
		}
		defer file.Close()
		fileOutput(abc, file)
	} else {
		printOutput(abc)
	}
}
