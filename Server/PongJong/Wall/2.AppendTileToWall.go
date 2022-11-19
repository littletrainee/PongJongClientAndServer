package Wall

import PJ "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong"

func (w *Wall) AppendTileToWall() {
	var temp []string

	for _, v := range PJ.Tile.Get() {
		for j := 0; j < 9; j++ {
			temp = append(temp, v)
		}
	}
	w.Wall.Set(temp)
}
