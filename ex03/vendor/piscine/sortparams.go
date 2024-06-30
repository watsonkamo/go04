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

func SortParams() {
	args := os.Args
	argCount := 0
	for range args {
		argCount++
	}

	for i := 1; i < argCount; i++ {
		for j := i + 1; j < argCount; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i := 1; i < argCount; i++ {
		PrintStr(args[i])
	}
}
