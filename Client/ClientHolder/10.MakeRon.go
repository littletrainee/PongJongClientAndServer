package ClientHolder

import (
	"fmt"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) MakeRon() {
	var (
		URL string = ch.GetURL(Enum.MakeRon)
		key string = ""
	)
	fmt.Printf("You Can Declare Ron, Do You Want Ron?(y/n)")
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
		ch.ClientData.GameData.Action = Enum.IsRon
	} else if key == "n" {
		ch.ClientData.GameData.Action = Enum.NotRon
	}
	GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
}
