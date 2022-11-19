package PongJong

import "github.com/littletrainee/GetSet"

var (
	Tile_Color []rune = []rune{'r', 'g', 'b'}
	Tile_Kind  []rune = []rune{'b', 'c', 'p'}
	Tile       GetSet.Type[[]string]
)
