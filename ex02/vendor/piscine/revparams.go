package piscine

import "ft"

func PrintStr (s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func ArgsCount(str []string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func RevParams(args []string) {
	for i := ArgsCount(args) - 1; i > 0; i-- {
		PrintStr(args[i])
		ft.PrintRune('\n')
	}
}
