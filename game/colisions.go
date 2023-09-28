package driver

import (
	"github.com/Tortugon09/gorace/models"
)

var t *models.Tux
var w *models.Windows

type CollisionDriver struct {
	gameOver bool
}

func NewCollisionDriver(tux *models.Tux, windows *models.Windows) *CollisionDriver {
	t = tux
	w = windows
	return &CollisionDriver{
		gameOver: false,
	}
}

func (c *CollisionDriver) Run() {
	for !c.gameOver {
		if w.GetPosY() >= 400 {
			if w.GetPosX() >= t.GetPosX()-50 && w.GetPosX() <= t.GetPosX()+50 {
				w.SetRunning(false)
				t.SetRunning(false)
				c.gameOver = true
			}
		}
	}
}

func (c *CollisionDriver) GetGameOver() bool {
	return c.gameOver
}
