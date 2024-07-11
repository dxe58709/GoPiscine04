package piscine

import "ft"

func PrintStr (s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func PrintParams(args []string) {
	args = args[1:]
	for _, arg := range args {
		PrintStr(arg)
		ft.PrintRune('\n')
	}
}