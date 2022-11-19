package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

func (h *Handler) MakeDeclareRiiChiAndDiscard() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		h.Player1.IsRiiChi.Set(true)
		h.Player1.ManualDiscard(h.PlayerDecision)
		h.Player1.SortHand()
	case 1:
		h.Player2.IsRiiChi.Set(true)
		h.Player2.ManualDiscard(h.PlayerDecision)
		h.Player2.SortHand()
	}
	h.Action.Set(Enum.CheckRon)
	h.GameState.TurnNext()
}
