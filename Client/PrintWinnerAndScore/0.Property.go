package PrintWinnerAndScore

type PrintWinnerAndScore struct {
	Name     string     `json:"name"`
	Hand     []string   `json:"hand"`
	Meld     [][]string `json:"meld"`
	Lastone  string     `json:"lastone"`
	Istsumo  string     `json:"istsumo"`
	YakuType string     `json:"yakutype"`
	Score    string     `json:"score"`
	Color    uint8      `json:"color"`
	Kind     uint8      `json:"kind"`
}
