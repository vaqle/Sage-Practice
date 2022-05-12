package queue

import "github.com/df-mc/dragonfly/server/player"

var Queued = make(map[string]*Queue)

func RemoveFromQueue(player *player.Player) {
	delete(Queued, player.Name())
}

func AddToQueue(player *player.Player, queue *Queue) {
	Queued[player.Name()] = queue
}

func IsQueued(player *player.Player) bool {
	_, ok := Queued[player.Name()]
	return ok
}
