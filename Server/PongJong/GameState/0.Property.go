package GameState

import "github.com/littletrainee/GetSet"

type GameState struct {
	GameOn    GetSet.Type[bool]
	GameTurn  GetSet.Type[uint8]
	GameRound GetSet.Type[uint8]
	Maxplayer GetSet.Type[uint8]
}
