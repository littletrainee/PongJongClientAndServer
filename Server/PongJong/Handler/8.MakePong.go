package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

func (h *Handler) MakePong() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		h.Player1.MakePongMeld(&h.Player2)
	case 1:
		h.Player2.MakePongMeld(&h.Player1)
	}
	h.Action.Set(Enum.CheckCanDeclareRiiChi)
}
