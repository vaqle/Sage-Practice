package phandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
)

type WorldHandler struct {
	world.NopHandler
	world *world.World
}

func AddToWorldHandler(w *world.World) *WorldHandler {
	handler := &WorldHandler{
		world: w,
	}
	return handler
}

func (w *WorldHandler) HandleLiquidFlow(ctx *event.Context, pos cube.Pos, pos2 cube.Pos, bl1 world.Block, bl2 world.Block) {
	ctx.Cancel()
}
