package duel

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"golang.org/x/exp/slices"
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
	return slices.Contains(d.players, player)
}

func (d *Duel) GetDuelPlayers() []*player.Player {
	return d.players
}

func (d *Duel) addPlayer(player *player.Player) {
	d.players = append(d.players, player)
}

func (d *Duel) GetPlayers() []*player.Player {
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

var (
	channel2 = make(chan int)
	cd       = 10
	phase    = 0
)

func (d *Duel) tick() {
	go func() {
		for {
			if phase == 0 && cd == 10 {
				fmt.Println("Countdown started")
				d.countDown()
			}
			select {
			case <-channel2:
				d.stop()
				break
			}
		}
	}()
}

func (d *Duel) countDown() {
	p1 := d.GetDuelPlayers()[0]
	p2 := d.GetDuelPlayers()[1]
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		var i = &cd
		*i--
		fmt.Println(*i)
		if *i == 0 {
			d.start()
			phase = 1
			fmt.Println("DUEL STARTED???")
			ticker.Stop()
		}
		switch *i {
		case 1, 2, 3, 4, 5:
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
		var duelTime = &d.duration
		*duelTime++
		if *duelTime == 60 {
			fmt.Println("DUEL TIME IS UP")
			d.stop()
			ticker.Stop()
		}
	}
}

func (d *Duel) GetOpponet(winner *player.Player) *player.Player {
	for _, p := range d.GetDuelPlayers() {
		if p.Name() != winner.Name() {
			return p
		}
	}
	return nil
}

func (d *Duel) End(winner *player.Player, loser *player.Player) {
	winner.Message("You won the duel!")
	loser.Message("You lost the duel!")
	d.stop()
}
