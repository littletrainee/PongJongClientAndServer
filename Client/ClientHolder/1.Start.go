package ClientHolder

import (
	"sync"

	"github.com/google/uuid"
)

func (ch *ClientHolder) Start(name string) {
	ch.Address.Set("0.0.0.0")
	ch.Port.Set("8080")
	ch.Pattern = []string{
		"CheckConnect",
		"JoinAGame",
		"GetRoomUUID",
		"GetGameData",
		"MakeADecisionToDiscard",
		"DeclareRiiChi",
		"MakeRon",
		"MakePong",
	}
	ch.ContentType = "application/json"
	ch.wg = &sync.WaitGroup{}
	ch.ClientData.Name = name
	ch.ClientData.ThisPlayerUUID = uuid.NewString()
}
