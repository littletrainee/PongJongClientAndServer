package GameData

import "fmt"

func (gd *GameData) Print(name string) {
	// opposite player
	fmt.Printf(gd.OppositePlayerName + "'s Meld:")
	for _, meld := range gd.OppositePlayerMeld {
		fmt.Printf("[ ")
		for i, v := range meld {
			Printonebyone(v, i, len(meld), " ]")
		}
	}
	fmt.Printf("\n\n")
	fmt.Printf(gd.OppositePlayerName + "'s River: ")
	for i, v := range gd.OppositePlayerRiver {
		Printonebyone(v, i, len(gd.OppositePlayerRiver), "\n")
	}

	fmt.Printf("\n\n%s's River: ", name)

	// this player
	for i, v := range gd.ThisPlayerRiver {
		Printonebyone(v, i, len(gd.ThisPlayerRiver), "\n")
	}
	fmt.Println()
	fmt.Printf(name + "'s Meld: ")
	for _, meld := range gd.ThisPlayerMeld {
		fmt.Printf("[ ")
		for i, v := range meld {
			Printonebyone(v, i, len(meld), " ]")
		}
	}
	fmt.Println()
	fmt.Printf(name + "'s Hand:")
	for i, v := range gd.ThisPlayerHand {
		Printonebyone(v, i, len(gd.ThisPlayerHand), "\n")
	}

}
