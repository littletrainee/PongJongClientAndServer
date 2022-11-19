package Handler

import (
	"time"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
)

func (h *Handler) Update() {
	for h.GameState.GameOn.Get() { // GameOn loop
		if len(h.Wall.Wall.Get()) == 0 { // Check wall is empty
			h.Winner.Set("DRAW")
		}
		switch h.Action.Get() {
		case Enum.CheckTsumo:
			h.CheckTsumo() // check tsumo
		case Enum.CheckCanDeclareRiiChi:
			h.CheckCanDeclareRiiChi()
		case Enum.MakeDeclareRiiChiAndDiscard:
			h.MakeDeclareRiiChiAndDiscard()
		case Enum.IsDiscard:
			h.MakeDiscard()
		case Enum.CheckRon:
			h.CheckRon()
		case Enum.IsRon, Enum.IsTsumo:
			if h.Player1.Iswin.Get() {
				h.PrintWin.AssignAndCalculate(&h.Player1, &h.Player2, h.Action.Get())
			} else if h.Player2.Iswin.Get() {
				h.PrintWin.AssignAndCalculate(&h.Player2, &h.Player1, h.Action.Get())
			}
			h.PrintWin.SetYakuAndPoint()
			h.Action.Set(Enum.PrintWinnerAndScore)
			// h.GameState.GameOn.Set(false)
		case Enum.CheckMeld:
			h.CheckMeld()
		case Enum.IsPong:
			h.MakePong()
		case Enum.DrawCard:
			h.DrawCard()
		default:
		}
		time.Sleep(time.Second / 10)
	}
	// set who win

}
