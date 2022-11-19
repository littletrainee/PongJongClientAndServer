package Player

import (
	"fmt"
	"testing"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Wall"
)

func TestRiichiCheck(t *testing.T) {
	var (
		p Player
		w Wall.Wall
	)
	w.CreateTile()
	p.Hand.Set([]string{"rb", "rb", "rb", "gb", "gb", "gb", "bb", "bb", "bp"})
	v := p.RiiChiCheck()
	if len(v) == 0 {
		t.Error("Error")
	} else {
		fmt.Println(v)
	}
}
