package Handler

import "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Player"

func checkpongbygoroutine(p2, p1 *Player.Player, h *Handler) {
	var riverlastelement string = p1.River.Get()[len(p1.River.Get())-1]
	go func() {
		if !p2.IsRiiChi.Get() {
			p2.CheckHasPongMeld(riverlastelement, h.wg)
		} else {
			h.wg.Done()
		}
	}()
}
