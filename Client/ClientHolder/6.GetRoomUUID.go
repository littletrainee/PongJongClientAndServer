package ClientHolder

import (
	"fmt"
	"time"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

// Get RoomUUID From Server if RoomUUID IS Empty Then Sleep 3 Second For Repeat Request From Server
func (ch *ClientHolder) GetRoomUUID() {
	defer ch.wg.Done()
	var URL string = ch.GetURL(Enum.GetRoomUUID)
	if ch.ClientData.IsJoinAGame {
		for {
			GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
			fmt.Println(ch.ClientData.RoomUUID)
			if ch.ClientData.RoomUUID != "" {
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}
