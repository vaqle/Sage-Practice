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
	rankID int
}

func getDatabase() *sql.DB {
	db := database.GetDatabase()
	return db
}

func (s *Session) Load() {
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
	s.init()
}

func (s Session) register() {
	d := getDatabase()
	_, _ = d.Query("INSERT INTO sessionPlayers (name, kills, deaths) VALUES (?, ?, ?)", s.player.Name(), 0, 0)
	fmt.Printf("Registered %s", s.player.Name())
	s.init()
}

func (s *Session) init() {
	s.SendLobbyItems()
	s.UpdateScoreboard()
}

func (s Session) Save() {

	d := getDatabase()
	_, _ = d.Query("UPDATE sessionPlayers SET kills = ?, deaths = ? WHERE name = ?", s.kills, s.deaths, s.player.Name())
	fmt.Printf("Saved session for %s", s.player.Name())
}

func (s *Session) AddKills() {
	s.kills++
}

func (s *Session) UpdateScoreboard() {
	sb := scoreboard.New("§6§lSage")
	sb.RemovePadding()
	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		if s.player == nil {
			ticker.Stop()
			return
		}

		sb.Set(0, "\uE000")
		sb.Set(1, " §r§fOnline: §6"+database.Players.String())
		sb.Set(2, " §r§fMatches: §6"+"0")
		sb.Set(3, "\n")
		sb.Set(4, " §r§6sagehcf.club")
		sb.Set(5, "§e\uE000")
		s.player.SendScoreboard(sb)
	}
}

func (s *Session) SendLobbyItems() {
	p := s.player
	Rduels := item.NewStack(item.Sword{
		Tier: item.ToolTierIron}, 1).WithCustomName("§r§6§lRanked Duels")
	_ = p.Inventory().SetItem(0, Rduels)
}

func (s Session) getRankId() int {
	return s.rankID
}

func (s Session) GetKills() string {
	kills := fmt.Sprintf("%d", s.kills)
	return kills
}
