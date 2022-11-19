package GameState

func (gs *GameState) TurnNext() {
	temp := gs.GameTurn.Get()
	if temp < gs.Maxplayer.Get()-1 {
		temp++
	} else {
		temp = 0
	}
	gs.GameTurn.Set(temp)
}
