package forms

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
)

type FFAForm struct {
	Nodebuff, Sumo form.Button
}

func (f FFAForm) Submit(submitter form.Submitter, pressed form.Button) {
	p := submitter.(*player.Player)
	switch pressed.Text {
	case f.Nodebuff.Text:
		p.Message("FFA!")
		break
	case f.Sumo.Text:
		p.Message("Sumo!")
		break
	}
}

func SendFFaForm(player *player.Player) {
	f := form.NewMenu(form.MenuSubmittable(FFAForm{
		Nodebuff: form.Button{Text: "Nodebuff", Image: "textures/items/potion_bottle_splash_heal"},
		Sumo:     form.Button{Text: "Sumo"},
	}), "FFA")
	player.SendForm(f)
}
