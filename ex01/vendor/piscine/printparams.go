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

func PrintParams() {
	for _, c := range os.Args[1:] {
		PrintStr(c)
	}
}
