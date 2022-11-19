package ClientData

import "github.com/littletrainee/PongJong_Client_And_Server/Client/GameData"

type ClientData struct {
	Name           string            `json:"name"`
	ThisPlayerUUID string            `json:"uuid"`
	IsConnect      bool              `json:"isconnect"`
	IsJoinAGame    bool              `json:"isjoinagame"`
	RoomUUID       string            `json:"roomuuid"`
	GameData       GameData.GameData `json:"gamedata"`
}
