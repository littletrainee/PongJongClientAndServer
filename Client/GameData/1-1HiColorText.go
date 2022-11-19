package GameData

import (
	"strings"

	"github.com/fatih/color"
)

var (
	HiBlueText  *color.Color = color.New(color.FgHiBlue)
	HiRedText   *color.Color = color.New(color.FgHiRed)
	HiGreenText *color.Color = color.New(color.FgHiGreen)
)

func PrintRGB(s string) {
	switch s[0] {
	case 'r':
		HiRedText.Printf("%v", strings.ToUpper(string(s[1])))
	case 'g':
		HiGreenText.Printf("%v", strings.ToUpper(string(s[1])))
	case 'b':
		HiBlueText.Printf("%v", strings.ToUpper(string(s[1])))
	}
}
