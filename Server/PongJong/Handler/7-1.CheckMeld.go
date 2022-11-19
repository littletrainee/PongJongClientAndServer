package Handler

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

func (h *Handler) CheckMeld() {
	h.wg.Add(int(h.GameState.Maxplayer.Get()) - 1)
	switch h.GameState.GameTurn.Get() {
	case 0:
		if !h.Player1.IsRiiChi.Get() {
			checkpongbygoroutine(&h.Player1, &h.Player2, h)
		} else {
			h.wg.Done()
		}
	case 1:
		if !h.Player2.IsRiiChi.Get() {
			checkpongbygoroutine(&h.Player2, &h.Player1, h)
		} else {
			h.wg.Done()
		}
	}
	h.wg.Wait()
	if h.Player1.HasPongMeld.Get() || h.Player2.HasPongMeld.Get() {
		h.Action.Set(Enum.CanPong)
	} else {
		h.Action.Set(Enum.DrawCard)
	}
}
