package Wall

import (
	"math/rand"
	"time"

	. "github.com/littletrainee/GetSet"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

type Wall struct {
	Name Type[string]
	Wall Type[[]string]
}
