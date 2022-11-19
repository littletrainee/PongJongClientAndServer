package ServerHolder

import (
	"net/http"
	"sync"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
)

func (sh *ServerHolder) Start() {
	sh.wg = &sync.WaitGroup{}
	sh.Address.Set("0.0.0.0")
	sh.Port.Set([]string{"8080"})
	sh.PatternMap = make(map[Enum.Pattern]func(http.ResponseWriter, *http.Request))
	sh.PatternMap[Enum.CheckConnect] = sh.CheckConnect
	sh.PatternMap[Enum.JoinAGame] = sh.JoinAGame
	sh.PatternMap[Enum.GetRoomUUID] = sh.GetRoomUUID
	sh.PatternMap[Enum.GetGameData] = sh.GetGameData
	sh.PatternMap[Enum.DeclareRiiChi] = sh.DeclareRiiChi
	sh.PatternMap[Enum.MakeADecisionToDiscard] = sh.MakeADecisionToDiscard
	sh.PatternMap[Enum.MakeRon] = sh.MakeRon
	sh.PatternMap[Enum.MakePong] = sh.MakePong
}
