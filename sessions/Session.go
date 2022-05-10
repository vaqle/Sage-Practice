package sessions

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"time"
)
import "Df_while_go/database"
import "database/sql"

type Session struct {
	player *player.Player
	kills  int
	deaths int
}

func getDatabase() *sql.DB {
	db := database.GetDatabase()
	return db
}

func (s *Session) load() {
	d := getDatabase()
	query, err := d.Query("SELECT * FROM sessionPlayers WHERE name = ?", s.player.Name())
	if err != nil {
		panic(err)
	}
	if !query.Next() {
		s.register()
		return
	}
	//get the kills and deaths
	name := s.player.Name()
	err = query.Scan(&name, &s.kills, &s.deaths)
	if err != nil {
		panic(err)
	}
	//set the player's kills and deaths
	fmt.Printf("Loaded %s's session\n", s.player.Name())
}

func (s Session) register() {
	d := getDatabase()
	d.Query("INSERT INTO sessionPlayers (name, kills, deaths) VALUES (?, ?, ?)", s.player.Name(), 0, 0)
	fmt.Printf("Registered %s", s.player.Name())
}

func (s Session) Save() {

	d := getDatabase()
	d.Query("UPDATE sessionPlayers SET kills = ?, deaths = ? WHERE name = ?", s.kills, s.deaths, s.player.Name())
	fmt.Printf("Saved session for %s", s.player.Name())
}

func (s *Session) AddKills() {
	s.kills++
}

func (s *Session) UpdateScoreboard() {
	sb := scoreboard.New("§r§6§lSage Practice")
	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		if s.player == nil {
			ticker.Stop()
			return
		}
		lines := []string{
			"",
			"§r§6Kills: §r§f" + fmt.Sprintf("%d", s.kills),
			"§r§6Deaths: §r§f" + fmt.Sprintf("%d", s.deaths),
			"",
		}
		i := 0
		for _, line := range lines {
			sb.Set(i, line)
			i++
		}
		s.player.SendScoreboard(sb)
	}
}

func (s *Session) SendLobbyItems() {
	p := s.player
	Rduels := item.NewStack(item.Sword{
		Tier: item.ToolTierIron}, 1).WithCustomName("§r§6§lRanked Duels")
	_ = p.Inventory().SetItem(0, Rduels)
}

func (s Session) GetKills() string {
	kills := fmt.Sprintf("%d", s.kills)
	return kills
}
