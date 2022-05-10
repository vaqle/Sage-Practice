package phandler

import (
	"Df_while_go/sessions"
	"fmt"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type PlayerHandler struct {
	player.NopHandler
	player *player.Player
}

func AddToHandler(player *player.Player) *PlayerHandler {
	handler := &PlayerHandler{
		player: player,
	}
	return handler
}

func (p *PlayerHandler) HandleItemUse(context *event.Context) {
	held, _ := p.player.HeldItems()
	names := []string{"LOL"}
	name := held.CustomName()
	for _, n := range names {
		if n == name {
			switch name {
			case "LOL":
				fmt.Println("LOL")
			}
		}
	}
}
func (p *PlayerHandler) HandleQuit() {
	sessions.DeleteSession(p.player.Name())
}
