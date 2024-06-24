package piscine

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintProgramName() {
	argv := os.Args[0]
	argv = argv[2:]

	PrintStr(argv)
}
