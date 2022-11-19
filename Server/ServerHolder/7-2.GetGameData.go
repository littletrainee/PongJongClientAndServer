package ServerHolder

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/ClientData"
)

func (sh *ServerHolder) GetGameData(writer http.ResponseWriter, reader *http.Request) {
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			for _, v := range sh.GameHandlerList {
				if v.Room.Get()[0] == CD.RoomUUID {
					switch CD.Name {
					case v.Player1.Name.Get():
						CD.GameData = getdata(&v.Player1, &v.Player2, &v.GameState)
					case v.Player2.Name.Get():
						CD.GameData = getdata(&v.Player2, &v.Player1, &v.GameState)
					}
					CD.GameData.Action = v.Action.Get()
					CD.GameData.PrintWin = v.PrintWin
				}
			}
			if err := json.NewEncoder(writer).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
