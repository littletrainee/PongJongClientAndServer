package ClientHolder

import (
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) GetGameData() {
	var URL string = ch.GetURL(Enum.GetGameData)
	GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
}
