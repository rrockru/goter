package world

import (
	"../types"
)

func NewWorld() *types.World {
	w := new(types.World)
	w.DayTime = true
	w.Time = 13500.
	
	w.RightWorld = 134400.
	w.BottomWorld = 38400.
	
	w.MaxTilesX = int(w.RightWorld / 16 + 1)
	w.MaxTilesY = int(w.BottomWorld / 16 + 1)
	
	w.CloudLimit = 200
	w.NumClouds = w.CloudLimit
	
	return w
}