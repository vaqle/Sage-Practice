package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Gamemode struct {
	Number int
}

func (c Gamemode) Run(src cmd.Source, output *cmd.Output) {
	sender := src.(*player.Player)
	switch c.Number {
	case 0:
		sender.SetGameMode(world.GameModeSurvival)
		output.Printf("Gamemode set to survival.")
	case 1:
		sender.SetGameMode(world.GameModeCreative)
		output.Printf("Gamemode set to creative.")

	}
}
