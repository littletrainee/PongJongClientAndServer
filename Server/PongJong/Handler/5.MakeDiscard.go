package Handler

import (
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
)

func (h *Handler) MakeDiscard() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		if h.Player1.IsRiiChi.Get() {
			h.Player1.AutoDiscard()
		} else {
			h.Player1.ManualDiscard(h.PlayerDecision)
			h.Player1.SortHand()
		}
	case 1:
		if h.Player2.IsRiiChi.Get() {
			h.Player2.AutoDiscard()
		} else {
			h.Player2.ManualDiscard(h.PlayerDecision)
			h.Player2.SortHand()
		}
	}
	h.Action.Set(Enum.CheckRon)
	h.GameState.TurnNext()
}
