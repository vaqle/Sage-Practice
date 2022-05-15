package phandler

import (
	"Df_while_go/duel"
	"Df_while_go/forms"
	"Df_while_go/queue"
	"Df_while_go/sessions"
	"github.com/df-mc/dragonfly/server/entity/damage"
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

func (p *PlayerHandler) HandleDeath(source damage.Source) {
	if duel.GetMatch(p.player) == nil {
		return
	}
	match := duel.GetMatch(p.player)
	if s, yes := source.(damage.SourceEntityAttack); yes {
		for _, dplayers := range match.GetDuelPlayers() {
			if s.Attacker.Name() == match.GetOpponet(dplayers).Name() {
				match.End(match.GetOpponet(p.player), p.player)
			}
		}
	}
}

func (p *PlayerHandler) HandleItemUse(*event.Context) {
	held, _ := p.player.HeldItems()
	names := []string{"§r§6§lRanked Duels"}
	name := held.CustomName()
	for _, n := range names {
		if n == name {
			event.C().Cancel()
			switch name {
			case "§r§6§lRanked Duels":
				forms.SendFFaForm(p.player)
			}
		}
	}
}
func (p *PlayerHandler) HandleQuit() {
	queue.RemoveFromQueue(p.player)
	sessions.DeleteSession(p.player.Name())
}
