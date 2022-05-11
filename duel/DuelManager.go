package duel
//TODO USE MUTEXES FOR DUELS add() destroy() get()
var duels = make(map[string]*DuelMatch)

func getDuels() map[string]*DuelMatch {
	return duels
}
