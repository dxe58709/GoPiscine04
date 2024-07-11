package main

import (
		"piscine"
		"ft"
	)

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func main () {
	PrintStr(piscine.PrintProgramName())
}
