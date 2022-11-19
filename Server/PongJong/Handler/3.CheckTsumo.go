package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

// Check Current Player is Tsumo
func (h *Handler) CheckTsumo() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		h.Player1.TsumoCheck()
	case 1:
		h.Player2.TsumoCheck()
	}
	if h.Player2.Iswin.Get() || h.Player1.Iswin.Get() { // if player2 or player1 is win(tsumo)
		h.Action.Set(Enum.IsTsumo) // set action to istsumo
	} else {
		h.Action.Set(Enum.CheckCanDeclareRiiChi) // set action to check can declare riichi
	}
}
