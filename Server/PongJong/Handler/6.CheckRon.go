package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

func (h *Handler) CheckRon() {
	h.wg.Add(1)
	switch h.GameState.GameTurn.Get() {
	case 0:
		h.Player1.RonCheck(&h.Player2, h.wg)
		h.wg.Wait()
		if h.Player1.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.Action.Set(Enum.CanRon)
		} else if h.Player1.IsRiiChi.Get() {
			h.Action.Set(Enum.DrawCard)
		} else {
			h.Action.Set(Enum.CheckMeld)
		}
	case 1:
		h.Player2.RonCheck(&h.Player1, h.wg)
		h.wg.Wait()
		if h.Player2.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.Action.Set(Enum.CanRon)
		} else if h.Player2.IsRiiChi.Get() {
			h.Action.Set(Enum.DrawCard)
		} else {
			h.Action.Set(Enum.CheckMeld)
		}
	}
}
