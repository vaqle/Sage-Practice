package sessions

import (
	"github.com/df-mc/dragonfly/server/player"
)

var sessions = make(map[string]*Session)

// CreateSession /**
func CreateSession(player *player.Player) *Session {
	sessions[player.Name()] = &Session{player: player}
	go sessions[player.Name()].load()
	return sessions[player.Name()]
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
