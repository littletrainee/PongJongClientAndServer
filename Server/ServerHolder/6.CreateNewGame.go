package ServerHolder

import "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Handler"

func (sh *ServerHolder) CreateNewGame(Room []string) {
	sh.GameHandlerList = append(sh.GameHandlerList, Handler.Handler{})

	sh.GameHandlerList[0].Start(Room)
	sh.GameHandlerList[0].Update()
}
