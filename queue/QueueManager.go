package queue

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
)

var queued = make(map[string]*Queue)

func getQueuedPlayers() map[string]*Queue {
	return queued
}

func getQueue(player *player.Player) *Queue {
	return queued[player.Name()]
}

func removeFromQueue(player *player.Player) {
	delete(queued, player.Name())
}

func addToQueue(player *player.Player, queue *Queue) {
	queued[player.Name()] = queue
}

func matchMaking() {
	go func() {
		for _, queue := range queued {
			s := server.Server{}
			players := s.Players()
			for _, onlineplayers := range players {
				if isQueued(onlineplayers) {
					q := getQueue(onlineplayers)
					if q.Equals(queue) == false {
						//TODO
					}
				}
			}
		}
	}()
}

func isQueued(player *player.Player) bool {
	_, ok := queued[player.Name()]
	return ok
}
