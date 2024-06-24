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

func RevParams() {
	for i := len(os.Args) - 1; i > 0; i-- {
		PrintStr(os.Args[i])
	}
}
