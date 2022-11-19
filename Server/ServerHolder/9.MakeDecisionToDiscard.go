package ServerHolder

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/ClientData"
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
)

func (sh *ServerHolder) MakeADecisionToDiscard(write http.ResponseWriter, reader *http.Request) {
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			// find target room to set discard and next player
			for i := range sh.GameHandlerList {
				if sh.GameHandlerList[i].Room.Get()[0] == CD.RoomUUID {
					if !CD.GameData.ThisPlayerIsRiiChi {
						// set target handler playerdecision to thisplayerdecision
						sh.GameHandlerList[i].PlayerDecision.Set(CD.GameData.ThisPlayerDecision)
						// set CD GameData ThisPlayerDecision to empty
						CD.GameData.ThisPlayerDecision = ""
						// set CheckTsumo to true
					}
					sh.GameHandlerList[i].Action.Set(Enum.IsDiscard)
					break
				}
			}
			if err := json.NewEncoder(write).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
