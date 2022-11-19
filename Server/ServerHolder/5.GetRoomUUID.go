package ServerHolder

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/ClientData"
	"github.com/littletrainee/slices"
)

func (sh *ServerHolder) GetRoomUUID(writer http.ResponseWriter, reader *http.Request) {
	// declare
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			for _, v := range sh.GameHandlerList {
				if slices.ContainsElement(v.Room.Get(), CD.Name) {
					CD.RoomUUID = v.Room.Get()[0]
					break
				} else {
					CD.RoomUUID = ""
				}
			}
			if err := json.NewEncoder(writer).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
