package ui

import (
	"TurkishDraughts/Board"

	"github.com/faiface/pixel/imdraw"
)

type DrawTheme interface {
	DrawBoard(*imdraw.IMDraw)
	DrawPieces(*imdraw.IMDraw, *board.BoardState)
	DrawSelected(*imdraw.IMDraw, int)
	DrawMoves(*imdraw.IMDraw, int, map[int][]int)
	DrawChecks(*imdraw.IMDraw, map[int][]int)
}