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

var phase = 0

func (d *Duel) tick() {
	c := make(chan int, 1) //buffered channel
	go d.countDown(c)
	finished, _ := <-c
	if finished == 0 && phase == 1 {
		d.start()
	}
}

func (d *Duel) countDown(c chan int) {
	p1 := d.getDuelPlayers()[0]
	p2 := d.getDuelPlayers()[1]
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		var i = 10
		i--
		if i == 0 {
			phase = 1
			c <- i
			ticker.Stop()
			close(c)
		}
		switch i {
		case 1, 2, 3, 4:
			p1.Message(fmt.Sprintf("%d", i))
			p2.Message(fmt.Sprintf("%d", i))
			break
		}
	}
}

func (d *Duel) start() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		var duelTime = &d.duration
		*duelTime++
		if len(d.getDuelPlayers()) == 0 {
			var p *player.Player = nil
			elements := duelPlayers[0:]
			for _, element := range elements {
				p = element
			}
			loser := d.GetOpponet(p)
			d.end(p, loser)
		}
		switch d.duration {
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
