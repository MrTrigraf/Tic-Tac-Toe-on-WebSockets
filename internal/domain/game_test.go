package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame_CanMove(t *testing.T) {
	tests := []struct {
		name string // description of this test case

		xMask    uint16
		oMask    uint16
		position int
		want     bool
	}{
		{
			name:     "empty board, position 0",
			xMask:    0b000000000,
			oMask:    0b000000000,
			position: 0,
			want:     true,
		},
		{
			name:     "position 0, occupied x",
			xMask:    0b000000001,
			oMask:    0b000000000,
			position: 1,
			want:     false,
		},
		{
			name:     "position 8, occupied o",
			xMask:    0b000000001,
			oMask:    0b100000000,
			position: 8,
			want:     false,
		},
		{
			name:     "position 9, more board",
			xMask:    0b000000001,
			oMask:    0b100000000,
			position: 9,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{
				xMask: tt.xMask,
				oMask: tt.oMask,
			}
			got := game.CanMove(tt.position)
			require.Equal(t, tt.want, got, fmt.Sprintf("CanMove(%d) = %v, want %v", tt.position, got, tt.want))
		})
	}
}

func TestGame_Move(t *testing.T) {
	tests := []struct {
		name string // description of this test case

		xMask    uint16
		oMask    uint16
		symbol   symbol
		position int

		want uint16
	}{
		{
			name:     "position 0, symbol x move",
			xMask:    0b000000000,
			oMask:    0b000000000,
			symbol:   X,
			position: 0,
			want:     0b000000001,
		},
		{
			name:     "position 1, symbol o move",
			xMask:    0b000000000,
			oMask:    0b000000000,
			symbol:   O,
			position: 1,
			want:     0b000000010,
		},
		{
			name:     "position 5, symbol o move with filled board",
			xMask:    0b000000000,
			oMask:    0b000010010,
			symbol:   O,
			position: 7,
			want:     0b010010010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &Player{
				symbol: tt.symbol,
			}
			game := &Game{
				xMask:   tt.xMask,
				oMask:   tt.oMask,
				player1: p1,
			}
			game.Move(p1, tt.position)
			mask := game.xMask
			if tt.symbol == O {
				mask = game.oMask
			}
			require.Equal(t, tt.want, mask, fmt.Sprintf("Move(p1, %d) = mask: %09b, want = %09b", tt.position, mask, tt.want))
		})
	}
}

func TestGame_IsWinMove(t *testing.T) {
	tests := []struct {
		name string // description of this test case

		xMask  uint16
		oMask  uint16
		symbol symbol

		want bool
	}{
		{
			name: "win move x, rows 1",

			xMask:  0b000000111,
			oMask:  0b000000000,
			symbol: X,

			want: true,
		},
		{
			name: "win move o, rows 2",

			xMask:  0b000000000,
			oMask:  0b000111000,
			symbol: O,

			want: true,
		},
		{
			name: "win move o, rows 3",

			xMask:  0b000000000,
			oMask:  0b111000000,
			symbol: O,

			want: true,
		},
		{
			name: "win move x, cols 1",

			xMask:  0b001001001,
			oMask:  0b000000000,
			symbol: X,

			want: true,
		},
		{
			name: "win move x, cols 2",

			xMask:  0b010010010,
			oMask:  0b000000000,
			symbol: X,

			want: true,
		},
		{
			name: "win move o, cols 3",

			xMask:  0b000000000,
			oMask:  0b100100100,
			symbol: O,

			want: true,
		},
		{
			name: "win move o, diag",

			xMask:  0b000000000,
			oMask:  0b100010001,
			symbol: O,

			want: true,
		},
		{
			name: "win move x, anti diag",

			xMask:  0b001010100,
			oMask:  0b000000000,
			symbol: X,

			want: true,
		},
		{
			name: "not win move x",

			xMask:  0b000010100,
			oMask:  0b000000000,
			symbol: X,

			want: false,
		},
		{
			name: "not win move o",

			xMask:  0b000010100,
			oMask:  0b010000000,
			symbol: O,

			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &Player{
				symbol: tt.symbol,
			}
			game := &Game{
				xMask:   tt.xMask,
				oMask:   tt.oMask,
				player1: p1,
			}
			got := game.IsWinMove(p1)
			require.Equal(t, tt.want, got, fmt.Sprintf("IsWinMove(p1) = %v, want = %v", got, tt.want))
		})
	}
}
