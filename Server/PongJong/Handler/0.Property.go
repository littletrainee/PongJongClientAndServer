package Handler

import (
	"sync"

	"github.com/littletrainee/GetSet"
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/GameState"
	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Player"
	PS "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/PrintWinnerAndScore"
	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Wall"
)

type Handler struct {
	Room                     GetSet.Type[[]string]
	Wall                     Wall.Wall
	Player1, Player2, RongBy Player.Player
	GameState                GameState.GameState
	wg                       *sync.WaitGroup
	PlayerDecision           GetSet.Type[string]
	Winner                   GetSet.Type[string]
	Action                   GetSet.Type[Enum.Action]
	PrintWin                 PS.PrintWinnerAndScore
}
