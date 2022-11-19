package ClientHolder

import (
	"time"

	CC "github.com/littletrainee/ClearConsole"
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) Update() {
	ch.ConnectToServer()         // connect to Server
	if ch.ClientData.IsConnect { // check is connect
		ch.JoinAGame() // then join a game
	}
	if ch.ClientData.IsJoinAGame { // check isjoin a game
		ch.wg.Add(1)        // wait group add one time
		go ch.GetRoomUUID() // try to get room uuid by another thread
	}
	ch.wg.Wait()                      // wait for another thread end
	if ch.ClientData.RoomUUID != "" { // check room uuid isn't empty
		ch.GetGameData()
		for { //gameloop
			if ch.ClientData.GameData.Action == Enum.PrintWinnerAndScore {
				ch.GetGameData()
				CC.ClearConsole()
				ch.PrintWinner()
				break
			} else {
				if ch.ClientData.GameData.ThisPlayerCodeNumber ==
					ch.ClientData.GameData.WhoToDiscard { // check is turn to this player
					ch.GetGameData() // try get game data from server
					CC.ClearConsole()
					ch.ClientData.GameData.Print(ch.ClientData.Name)
					switch ch.ClientData.GameData.Action { // switch action
					case Enum.IsTsumo, Enum.IsRon: // action is istsumo
						// get win yaku
					case Enum.CanDeclareRiiChi:
						ch.DeclareRiiChi()
					case Enum.DisCard:
						ch.MakeADecisionToDiscard()
					case Enum.CanRon:
						ch.MakeRon()
					case Enum.CanPong:
						ch.MakePong()
					}
				} else {
					ch.GetGameData() // try get game data from server
					CC.ClearConsole()
					ch.ClientData.GameData.Print(ch.ClientData.Name)
				}
			}
			time.Sleep(time.Second / 10)
		}
	}
}
