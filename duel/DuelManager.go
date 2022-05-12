package duel

import (
	"github.com/df-mc/dragonfly/server/player"
	"sync"
)

type Matches struct {
	mList map[string]*Duel
	mutex sync.Mutex
}

var matches = Matches{mList: make(map[string]*Duel)}

func (m *Matches) register(d *Duel) *Duel {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.mList[d.UUID] = d
	return d
}

func (m *Matches) unregister(d *Duel) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.mList, d.UUID)
}

func GetMatch(p *player.Player) *Duel {
	for _, d := range matches.mList {
		for _, p2 := range d.getDuelPlayers() {
			if len(d.getPlayers()) != 0 {
				if p2.Name() == p.Name() {
					return d
				}
			}
		}
	}
	return nil
}

func StartMatch(p1 *player.Player, p2 *player.Player) {
	d := &Duel{
		players: []*player.Player{p1, p2}, UUID: "HEY",
		running: true, duration: 0,
	}
	d.addPlayer(p1)
	d.addPlayer(p2)
	d.tick()
	matches.register(d)

}

func (d *Duel) stop() {
	matches.unregister(d)
}
