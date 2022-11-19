package ServerHolder

import (
	"net/http"
	"sync"

	"github.com/littletrainee/GetSet"
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
	"github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Handler"
)

type ServerHolder struct {
	Address         GetSet.Type[string]
	Port            GetSet.Type[[]string]
	PatternMap      map[Enum.Pattern]func(http.ResponseWriter, *http.Request)
	ConnectList     []string
	WaitForGameList []string
	GameHandlerList []Handler.Handler
	wg              *sync.WaitGroup
}
