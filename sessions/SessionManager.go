package sessions

import (
	"github.com/df-mc/dragonfly/server/player"
)

var sessions = make(map[string]*Session)

// CreateSession /**
func CreateSession(player *player.Player) {
	sessions[player.Name()] = &Session{player: player}
}

// GetSession /**
func GetSession(name string) *Session {
	return sessions[name]
}

/**
Removes the sessions for the given player.
*/

func DeleteSession(player string) {
	go sessions[player].Save()
	delete(sessions, player)
}
