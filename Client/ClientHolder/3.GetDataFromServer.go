package ClientHolder

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/littletrainee/PongJong_Client_And_Server/Client/ClientData"
)

// Send ClientData to Server and Get response to ClientData
func GetDataFromServer(cd *ClientData.ClientData, URL, contentType string) {
	requestCD := new(bytes.Buffer)
	if err := json.NewEncoder(requestCD).Encode(cd); err != nil {
		log.Println(err)
	}
	responseCD, err := http.Post(URL, contentType, requestCD)
	if err != nil {
		log.Println(err)
	}

	defer responseCD.Body.Close()
	if err := json.NewDecoder(responseCD.Body).Decode(cd); err != nil {
		log.Println(err)
	}
}
