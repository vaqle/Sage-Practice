package duel

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"time"
)

var duelPlayers = make([]*player.Player, 1)

type Duel struct {
	players  []*player.Player
	UUID     string
	running  bool
	duration int
}

func (d *Duel) isAlive(player *player.Player) bool {
	for _, p := range d.getDuelPlayers() {
		if p.Name() == player.Name() {
			return true
		}
	}
	return false
}

func (d *Duel) getDuelPlayers() []*player.Player {
	return d.players
}

func (d *Duel) addPlayer(player *player.Player) {
	d.players = append(d.players, player)
}

func (d *Duel) getPlayers() []*player.Player {
	return duelPlayers[0:1]
}

func (d *Duel) isRunning() bool {
	return d.running
}

func (d *Duel) setRunning(r bool) {
	d.running = r
}

func (d *Duel) GetDuration() int {
	return d.duration
}

func (d *Duel) setDuration(dur int) {
	d.duration = dur
}

func (d *Duel) getUUID() string {
	return d.UUID
}

var cd = 10
var phase = 0

func (d *Duel) tick() {
	go func() {
		for {
			if phase == 0 && cd == 10 {
				d.countDown()
				if cd == 0 {
					d.start()
					return
				}
			}
		}
	}()
}

func (d *Duel) countDown() {
	p1 := d.getDuelPlayers()[0]
	p2 := d.getDuelPlayers()[1]
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		i := &cd
		*i--
		if *i == 0 {
			phase = 1
			d.start()
			ticker.Stop()
		}
		switch i {
		case i:
			p1.Message(fmt.Sprintf("%d", *i))
			p2.Message(fmt.Sprintf("%d", *i))
			break
		}
	}
}

func (d *Duel) start() {
	fmt.Println("DUEL STARTED")
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		duration := &d.duration
		*duration++
		amount := len(d.getDuelPlayers())
		if amount == 1 {
			elements := duelPlayers[1:]
			winner := elements[0]
			loser := d.GetOpponet(winner)
			winner.Message("You won the duel!")
			d.end(winner, loser)
		}
		switch *duration {
		case 60:
			fmt.Println("DUEL ENDED")
			d.stop()
			ticker.Stop()
			return
		}
	}
}

func (d *Duel) GetOpponet(winner *player.Player) *player.Player {
	for _, p := range d.getDuelPlayers() {
		if p.Name() != winner.Name() {
			return p
		}
	}
	return nil
}

func (d *Duel) end(winner *player.Player, loser *player.Player) {
	winner.Message("You won the duel!")
	loser.Message("You lost the duel!")
	d.stop()
}
