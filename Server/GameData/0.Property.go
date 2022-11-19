package GameData

import (
	Enum "github.com/littletrainee/PongJong_Client_And_Server/Server/EnumHolder"
	PS "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/PrintWinnerAndScore"
)

type GameData struct {
	GameOn               bool                   `json:"gameon"`
	ThisPlayerHand       []string               `json:"thisplayerhand"`
	ThisPlayerMeld       [][]string             `json:"thisplayermeld"`
	ThisPlayerRiver      []string               `json:"thisplayerriver"`
	ThisPlayerCodeNumber uint8                  `json:"thisplayercodenumber"`
	ThisPlayerDecision   string                 `json:"thisplayerdecision"`
	ThisPlayerIsRiiChi   bool                   `json:"thisplayerisriichi"`
	OppositePlayerName   string                 `json:"oppositeplayername"`
	OppositePlayerMeld   [][]string             `json:"oppositeplayermeld"`
	OppositePlayerRiver  []string               `json:"oppositeplayerriver"`
	WhoToDiscard         uint8                  `json:"whotodiscard"`
	Action               Enum.Action            `json:"action"`
	PrintWin             PS.PrintWinnerAndScore `json:"printwinnerandscore"`
}
