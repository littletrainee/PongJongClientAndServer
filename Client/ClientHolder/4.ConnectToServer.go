package ClientHolder

import (
	"fmt"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) ConnectToServer() {
	var (
		key string
		URL string = ch.GetURL(Enum.CheckConnect)
	)
	if !ch.ClientData.IsConnect {
		fmt.Print(ch.ClientData.Name + " Want Connect To Server?(y/n)")
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
			GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
			fmt.Println(ch.ClientData.Name + " Is Connect To Server")
		} else if key == "n" {
			fmt.Println(ch.ClientData.Name + " Isn't Connect To Server")
		}
	} else {
		GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
		fmt.Println(ch.ClientData.Name + " Is Disconnect From Server")
	}
}
