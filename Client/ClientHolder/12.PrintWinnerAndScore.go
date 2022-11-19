package ClientHolder

import (
	"fmt"

	"github.com/littletrainee/PongJong_Client_And_Server/Client/GameData"
)

func (ch *ClientHolder) PrintWinner() {
	fmt.Printf("               %s  Is %s\n\n", ch.ClientData.GameData.PrintWinnerAndScore.Name, ch.ClientData.GameData.PrintWinnerAndScore.Istsumo)
	switch len(ch.ClientData.GameData.PrintWinnerAndScore.Meld) {
	case 1:
		fmt.Print("  ")
	case 0:
		fmt.Print("     ")
	}
	for i, v := range ch.ClientData.GameData.PrintWinnerAndScore.Hand {
		GameData.Printonebyone(v, i, len(ch.ClientData.GameData.PrintWinnerAndScore.Hand), "  ")
	}
	if len(ch.ClientData.GameData.PrintWinnerAndScore.Meld) != 0 {
		fmt.Print("     ")
		for _, meld := range ch.ClientData.GameData.PrintWinnerAndScore.Meld {
			fmt.Printf(" [ ")
			for i, v := range meld {
				GameData.Printonebyone(v, i, len(meld), " ]")
			}
		}
	}

	fmt.Print("       ")
	GameData.PrintRGB(ch.ClientData.GameData.PrintWinnerAndScore.Lastone)
	fmt.Printf("\n\n")

	fmt.Println("===========================================")
	fmt.Printf("役名                             %s\n", ch.ClientData.GameData.PrintWinnerAndScore.YakuType)
	fmt.Println("-------------------------------------------")
	fmt.Printf("親                                    %d\n", ch.ClientData.GameData.PrintWinnerAndScore.Score[0])
	fmt.Printf("子                                    %d\n", ch.ClientData.GameData.PrintWinnerAndScore.Score[1])
	fmt.Println("===========================================")
}
