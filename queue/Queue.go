package queue

import (
	"github.com/df-mc/dragonfly/server/player"
)

type Queue struct {
	player    *player.Player
	queuetype string
	ranked    bool
}

func (q *Queue) GetPlayer() *player.Player {
	return q.player
}

func (q *Queue) GetQueueType() string {
	return q.queuetype
}

func (q *Queue) GetRanked() bool {
	return q.ranked
}

func (q *Queue) Equals(other *Queue) bool {
	return q.player.Name() == other.player.Name()
}
