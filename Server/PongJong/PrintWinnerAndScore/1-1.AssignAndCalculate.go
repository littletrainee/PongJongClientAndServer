package PrintWinnerAndScore

import (
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Player"
	"github.com/littletrainee/slices"
)

// assign value to PrintWin Property
func (ps *PrintWinnerAndScore) AssignAndCalculate(thisplayer, otherplayer *Player.Player, IsTsumo Enum.Action) {
	ps.Name = thisplayer.Name.Get()
	ps.Meld = thisplayer.Meld.Get()
	if IsTsumo == Enum.IsTsumo {
		ps.Istsumo = "Tsumo"
		ps.Hand = thisplayer.Hand.Get()[:len(thisplayer.Hand.Get())-1]
		ps.Lastone = thisplayer.Hand.Get()[len(thisplayer.Hand.Get())-1]
	} else if IsTsumo == Enum.IsRon {
		ps.Istsumo = "Ron"
		ps.Hand = thisplayer.Hand.Get()
		ps.Lastone = otherplayer.River.Get()[len(otherplayer.River.Get())-1]
	}

	// Calculate Score And Point
	temp := ps.Hand
	if len(ps.Meld) != 0 {
		for _, meld := range ps.Meld {
			temp = append(temp, meld[:3]...)
		}
	}
	if len(temp) != 9 {
		temp = append(temp, ps.Lastone)
	}
	temp = slices.RemoveDuplicate(temp)

	if len(temp) == 1 {
		ps.Color = 1
		ps.Kind = 1
	} else if len(temp) == 9 {
		ps.Color = 3
		ps.Kind = 9
	} else {
		ps.Color, ps.Kind = splitColorAndKind(temp)
	}
}
