package ClientHolder

import (
	"fmt"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) JoinAGame() {
	var (
		key string
		URL string = ch.GetURL(Enum.JoinAGame)
	)

	if ch.ClientData.IsConnect {
		fmt.Print("Want Join A Game?(y/n)")
		for {
			fmt.Scanln(&key)
			if key != "y" && key != "n" {
				fmt.Print("Wrong Enter Please Renter:")
				key = ""
			} else {
				break
			}
		}
	}
	if key == "y" {
		GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
		fmt.Println(ch.ClientData.Name + " Is Join A Game.")
	} else if key == "n" {
		ch.ConnectToServer()
	}
}
