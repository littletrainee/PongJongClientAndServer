package PrintWinnerAndScore

import (
	"github.com/littletrainee/slices"
)

func splitColorAndKind(temp []string) (uint8, uint8) {
	var color, kind []rune

	for _, v := range temp {
		if !slices.ContainsElement(color, rune(v[0])) {
			color = append(color, rune(v[0]))
		}
		if !slices.ContainsElement(kind, rune(v[1])) {
			kind = append(kind, rune(v[1]))
		}
	}
	return uint8(len(color)), uint8(len(kind))
}
