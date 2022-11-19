package ClientHolder

import (
	"fmt"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) MakePong() {
	var (
		URL string = ch.GetURL(Enum.MakePong)
		key string = ""
	)

	fmt.Print("You Can Pong, Do You Want Pong?(y/n)")
	for {
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Print("Wrong Enter Please Renter:")
			key = ""
		} else {
			break
		}
	}
	if key == "y" {
		ch.ClientData.GameData.Action = Enum.IsPong
	} else if key == "n" {
		ch.ClientData.GameData.Action = Enum.NotPong
	}
	GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
}
