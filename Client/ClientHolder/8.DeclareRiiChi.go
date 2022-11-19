package ClientHolder

import (
	"fmt"
	"strconv"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"
)

func (ch *ClientHolder) DeclareRiiChi() {
	var (
		URL       string = ch.GetURL(Enum.DeclareRiiChi)
		keystring string
	)
	fmt.Print("You Can Declare RiiChi, Do You Want Declare Riichi?(y/n)")
	for {
		fmt.Scanln(&keystring)
		if keystring != "y" && keystring != "n" {
			fmt.Print("Wrong Enter Please Renter:")
			keystring = ""
		} else {
			break
		}
	}

	if keystring == "y" {
		fmt.Printf("Please select whitch one do you want to discard from index 1-%d:", len(ch.ClientData.GameData.ThisPlayerHand))
		for {
			fmt.Scanln(&keystring)
			keyint, _ := strconv.Atoi(keystring)
			if keyint > len(ch.ClientData.GameData.ThisPlayerHand) || keyint < 1 {
				fmt.Print("Wrong Enter Please Renter:")
				keystring = ""
			} else {
				break
			}
		}
		ch.ClientData.GameData.ThisPlayerDecision = keystring
		GetDataFromServer(&ch.ClientData, URL, ch.ContentType)
	} else if keystring == "n" {
		ch.ClientData.GameData.Action = Enum.DisCard
	}
}
