package Wall

import "math/rand"

// shuffle wall
func (w *Wall) Shuffle() {
	temp := w.Wall.Get()
	rand.Shuffle(len(temp), func(i, j int) { temp[i], temp[j] = temp[j], temp[i] })
	w.Wall.Set(temp)
}
