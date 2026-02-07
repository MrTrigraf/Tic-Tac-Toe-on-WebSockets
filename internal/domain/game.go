package domain

type state string

const (
	StateWaiting   = "waiting"
	StatePlaying   = "playing"
	StateFinishing = "finishing"
)

// все выигрышные комбинации
var winMasks []uint16 = []uint16{
	0b000000111, //rows
	0b000111000, //rows
	0b111000000, //rows

	0b001001001, //cols
	0b010010010, //cols
	0b100100100, //cols

	0b100010001, //diag
	0b001010100, //anti diag
}

// Game хранит в себе данные о игре.
//
// xMask и oMask хранят ходы определенных игроков в битовом виде.
type Game struct {
	xMask         uint16
	oMask         uint16
	player1       *Player
	player2       *Player
	currentMotion uint8
	stateGame     state
}

func NewGame(player *Player) *Game {
	return &Game{
		xMask:         0b000000000,
		oMask:         0b000000000,
		player1:       player,
		currentMotion: 0,
		stateGame:     StateWaiting,
	}
}

// CanMove проверяет может ли игрок сходить на определенную клетку на доске.
func (g *Game) CanMove(n int) bool {
	if n > 9 {
		return false
	}
	board := g.oMask | g.xMask
	if (board & 1 << n) == 0 {
		return true
	}
	return false
}

// Move сохранение хода игрока в его маску.
func (g *Game) Move(player *Player, n int) {
	if player.GetSymbol() == X {
		g.xMask |= 1 << n
		return
	}

	g.oMask |= 1 << n
}

// IsWinMove проверяет был ли ход победным.
func (g *Game) IsWinMove(player *Player) bool {
	pMask := g.oMask
	if player.GetSymbol() == X {
		pMask = g.xMask

	}
	for _, m := range winMasks {
		if (m & pMask) == m {
			return true
		}
	}
	return false
}
