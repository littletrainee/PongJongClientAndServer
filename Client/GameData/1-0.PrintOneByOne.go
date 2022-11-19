package GameData

import (
	"fmt"
)

func Printonebyone(value string, index int, length int, suffix string) {
	PrintRGB(value)
	if index != length-1 {
		fmt.Printf(", ")
	} else {
		fmt.Print(suffix)
	}
}
