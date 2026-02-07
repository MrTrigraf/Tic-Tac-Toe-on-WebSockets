package domain

type symbol byte

const (
	X symbol = 'x'
	O symbol = 'o'
)

type Player struct {
	id     int
	name   string
	symbol symbol
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		id:   id,
		name: name,
	}
}

func (p *Player) GetSymbol() symbol {
	return p.symbol
}
