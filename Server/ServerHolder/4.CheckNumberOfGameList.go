package ServerHolder

import "github.com/google/uuid"

// Check WaitingLine Player Number Can Make A Game
func (sh *ServerHolder) CheckNumberOfGameList() {
	for {
		if len(sh.WaitForGameList) >= 2 {
			uuid := uuid.NewString()                       // set room uuid
			room := []string{uuid}                         // Add uuid to Room
			room = append(room, sh.WaitForGameList[:2]...) // Add player to room
			// remove player from WaitForGameList
			for i := 0; i < 2; i++ {
				sh.WaitForGameList = sh.WaitForGameList[1:]
			}
			// sh.RoomsList = append(sh.RoomsList, room)
			go sh.CreateNewGame(room) // Add room to IsGameing list and Create a New Game
		}
	}
}
