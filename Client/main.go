package main

import (
	"github.com/littletrainee/PongJong_Client_And_Server/Client/ClientHolder"
)

func main() {
	var CH ClientHolder.ClientHolder
	CH.Start("Logan")
	CH.Update()
}
