package convert

func ArrAtoi(s []string) (arr []int, err string, count int) {
	for _, letter := range s {
		num, err := atoi(letter)
		if err != "" {
			count++
			continue
		}
		arr = append(arr, num)
	}
	if count > 0 {
		return arr, "error", count
	}
	return arr, "", 0
}

func atoi(s string) (n int, err string) {
	if s == "0" {
		return 0, ""
	}
	if s < "0" || s > "9" {
		return 0, "error"
	}
	for _, letter := range s {
		n = n*10 + int(letter-48)
	}
	return n, ""
}

// uint32 is the set of all unsigned 32-bit integers.
// Range: 0 through 4294967295.

// uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.

func ArrByteToArrInt(b []byte) (n []int) {
	numS := ""
	for _, vol := range b {
		if vol == ' ' || vol == '\n' {
			num, _ := atoi(numS)
			n = append(n, num)
			numS = ""
		} else {
			numS = numS + string(vol)
		}
	}
	if numS != "" {
		num, _ := atoi(numS)
		n = append(n, num)
	}
	return n
}

func ArrByteToString(b []byte) (s string) {
	numS := ""
	for _, vol := range b {
		if vol == ' ' || vol == '\n' {
			numS += "x"
		} else {
			numS = numS + string(vol)
		}
	}
	return s
}

func Itoa(n int) (s string) {
	if n == 0 {
		return "0"
	}
	for n != 0 {
		tail := n % 10
		s = string(rune(tail+'0')) + s
	}
	return s
}
