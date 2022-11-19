package Wall

import (
	"testing"
)

func TestAppendTileToWall(t *testing.T) {
	var w Wall
	w.CreateTile()
	w.AppendTileToWall()
	if len(w.Wall.Get()) != 81 {
		t.Error("Wrong")
	} else {
		w.Print()
	}
}
