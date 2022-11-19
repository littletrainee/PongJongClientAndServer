package ServerHolder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Server/ClientData"
	"github.com/littletrainee/slices"
)

func (sh *ServerHolder) CheckConnect(writer http.ResponseWriter, reader *http.Request) {
	// declare
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			if !CD.IsConnect {
				fmt.Println(CD.Name + " Is Connected.")
				if !slices.ContainsElement(sh.ConnectList, CD.Name) {
					sh.ConnectList = append(sh.ConnectList, CD.Name)
				}
				CD.IsConnect = true
			} else if CD.IsConnect {
				fmt.Println(CD.Name + " Is DisConnected.")
				if slices.ContainsElement(sh.ConnectList, CD.Name) {
					index := slices.FindIndexOfElement(sh.ConnectList, CD.Name)
					sh.ConnectList = append(sh.ConnectList[:index], sh.ConnectList[index+1:]...)
				}
				CD.IsConnect = false
			}
			if err := json.NewEncoder(writer).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
