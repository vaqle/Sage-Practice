package duel

import (
	"github.com/df-mc/dragonfly/server/player"
	"time"
)

type DuelMatch struct {
	players  []*player.Player
	running  bool
	duration int
}

var alivePlayers = make(map[string]*player.Player)

var phase int = 0

var dur = 10

func (m DuelMatch) tick() {
	go func() {
		for {
			if isAlive(m.players[0]) && isAlive(m.players[1]) {
				if phase == 0 && dur == 10 {
					countDown()
				}
				if phase == 1 {
					//send kit
				}
			} else {
				var winner *player.Player = nil
				var loser *player.Player = nil

				m.end(winner, loser)
				return
			}

		}
	}()
}

func countDown() {
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		i := &dur
		*i--
		if *i == 0 {
			phase = 1
			ticker.Stop()
		}

	}
}

func (m DuelMatch) end(winner *player.Player, loser *player.Player) {

}

func (m DuelMatch) getDuration() int {
	return m.duration
}

func isAlive(player *player.Player) bool {
	return alivePlayers[player.Name()] != nil
}

func addAlivePlayer(player *player.Player) {
	alivePlayers[player.Name()] = player
}

func (m DuelMatch) getPlayers() []*player.Player {
	return m.players
}

func (m DuelMatch) getPlayer() *player.Player {
	return m.players[0]
}

func (m DuelMatch) getOpponet() *player.Player {
	return m.players[1]
}
