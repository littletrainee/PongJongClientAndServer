package ClientHolder

import (
	"sync"

	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/PongJong_Client_And_Server/Client/ClientData"
)

type ClientHolder struct {
	Address     GetSet.Type[string]
	Port        GetSet.Type[string]
	Pattern     []string
	ContentType string
	wg          *sync.WaitGroup
	ClientData  ClientData.ClientData
}
