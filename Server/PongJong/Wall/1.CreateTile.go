package Wall

import PJ "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong"

func (w *Wall) CreateTile() {
	var temp []string
	for _, v := range PJ.Tile_Color {
		for _, w := range PJ.Tile_Kind {
			temp = append(temp, string(v)+string(w))
		}
	}
	PJ.Tile.Set(temp)
}
