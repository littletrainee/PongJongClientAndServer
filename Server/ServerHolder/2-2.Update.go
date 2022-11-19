package ServerHolder

import (
	"log"
	"net/http"
)

func (sh *ServerHolder) Update() {
	// goroutin each HandleFunc
	for _, value := range sh.PatternMap {
		temp := GetHandleFunctionName(value)
		go http.HandleFunc("/"+temp, value)
	}
	go sh.CheckNumberOfGameList()
	if err := http.ListenAndServe(sh.Address.Get()+":"+sh.Port.Get()[0], nil); err != nil {
		log.Println(err)
	}
}
