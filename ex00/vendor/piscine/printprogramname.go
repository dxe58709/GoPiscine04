package piscine

import "os"

func PrintProgramName () string {
	return os.Args [0][2:]
}
