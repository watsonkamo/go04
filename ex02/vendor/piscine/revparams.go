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
	argCount := 0
	for range os.Args {
		argCount++
	}

	for i := argCount - 1; i > 0; i-- {
		PrintStr(os.Args[i])
	}
}

// func RevParams() {
// 	argCount := StrLen(os.Args)

// 	for i := argCount - 1; i > 0; i-- {
// 		PrintStr(os.Args[i])
// 	}
// }

// func RevParams() {
// 	args := os.Args
// 	for i := StrLen(string(args[])) - 1; i > 0; i-- {
// 		PrintStr(os.Args[i])
// 	}
// }
