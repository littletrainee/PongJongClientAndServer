package Handler

import (
	"fmt"
	"sync"

	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Wall"
)

func (h *Handler) Start(Room []string) {
	h.Room.Set(Room)
	// Initialization Wall Random Seed
	Wall.Init()
	// Set GameState
	h.GameState.GameOn.Set(true) // loop switch
	h.GameState.GameTurn.Set(0)  // start with player 1
	h.GameState.GameRound.Set(0) // Start from 0 round
	h.GameState.Maxplayer.Set(2)
	h.wg = &sync.WaitGroup{}
	fmt.Println(h.PlayerDecision.Get())

	// Set Name
	h.Wall.Name.Set("Wall")
	h.Player1.Name.Set(h.Room.Get()[1])
	h.Player2.Name.Set(h.Room.Get()[2])
	// Set Player Code Number
	h.Player1.CodeNumber.Set(1)
	h.Player2.CodeNumber.Set(2)
	// Set Player 1 is bookmaker
	h.Player1.Bookmaker.Set(true)
	// Combine Tile_Color and Tile_Kind to Tile
	h.Wall.CreateTile()
	// Append Tile to Wall
	h.Wall.AppendTileToWall()
	// Shuffle Wall
	h.Wall.Shuffle()
	// Each Player Draw Card 2 Tile 4 time (8 Tile)
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			h.Player1.DrawCard(&h.Wall)
		}
		for j := 0; j < 2; j++ {
			h.Player2.DrawCard(&h.Wall)
		}
	}
	// Player 1  Draw the 9th tile
	h.Player1.DrawCard(&h.Wall)
	// Sort Hand
	h.Player1.SortHand()
	h.Player2.SortHand()
	h.Action.Set(Enum.CheckTsumo)

}
