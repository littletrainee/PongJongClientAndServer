package DisconnectionTimer

import (
	"fmt"
	"sync"
	"time"
)

func (dt *DisconnectionTimer) Update(wg *sync.WaitGroup) {
	defer wg.Done()
	for ; dt.StartTime != 0; dt.StartTime-- {
		fmt.Println(dt.StartTime)
		time.Sleep(time.Second)
	}
}
