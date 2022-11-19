package ServerHolder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/ClientData"
	"github.com/littletrainee/slices"
)

func (sh *ServerHolder) JoinAGame(writer http.ResponseWriter, reader *http.Request) {
	// declare
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			if !CD.IsJoinAGame {
				fmt.Println(CD.Name + " Is Join WaitLine For A Game.")
				if !slices.ContainsElement(sh.WaitForGameList, CD.Name) {
					sh.WaitForGameList = append(sh.WaitForGameList, CD.Name)
				}
				CD.IsJoinAGame = true
			}
			if err := json.NewEncoder(writer).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
