package main

import (
	. "github.com/littletrainee/PongJong_Client_And_Server/Server/ServerHolder"
)

func main() {
	var SH ServerHolder
	SH.Start()
	SH.Update()
}
