package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

// Check current Player Can Declare Riichi
func (h *Handler) CheckCanDeclareRiiChi() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		if h.Player1.IsRiiChi.Get() {
			h.Action.Set(Enum.DisCard)
		} else {
			if p := h.Player1.RiiChiCheck(); len(p) != 0 {
				h.Action.Set(Enum.CanDeclareRiiChi)
			} else {
				h.Action.Set(Enum.DisCard)
			}
		}
	case 1:
		if h.Player2.IsRiiChi.Get() {
			h.Action.Set(Enum.DisCard)
		} else {
			if p := h.Player2.RiiChiCheck(); len(p) != 0 {
				h.Action.Set(Enum.CanDeclareRiiChi)
			} else {
				h.Action.Set(Enum.DisCard)
			}
		}
	}
}
