package Wall

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
)

func (w *Wall) Print() {
	fmt.Printf("%s", w.Name.Get())
	for i, v := range w.Wall.Get() {
		CV.PrintRGB(v)
		if i != len(w.Wall.Get())-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n\n")
}
