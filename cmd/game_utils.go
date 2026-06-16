package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) GetNewActiveRow() []string {
	allRows := g.KanaConfTaken
	randNumRow := rand.Intn(len(allRows)) // 8
	randNumLetter := rand.Intn(4)
	row := allRows[randNumRow]
	letter := 3 * randNumLetter
	g.UserRow[0] = g.Kana[row][letter]
	g.UserRow[1] = ""
	g.UserRow[2] = ""
	g.UserRow[3] = row
	return []string{g.Kana[row][letter], g.Kana[row][letter+2], g.Kana[row][letter+1], row}
}

func (g *Game) GetNewResBtnValues() []string {
	numLetter := 0
	res := []string{}
	var resf []int
	row := g.ActiveRow[3]
	for i := range g.Kana[row] {
		if g.Kana[row][i] == g.ActiveRow[0] {
			break
		}
		numLetter += 1
	}
	if numLetter < 6 {
		resf = []int{numLetter, numLetter + 3, numLetter + 6, numLetter + 9} // 3 6 9 12
	} else if numLetter == 6 {
		resf = []int{numLetter, numLetter + 3, numLetter - 3, numLetter + 6}
	} else {
		resf = []int{numLetter, numLetter - 3, numLetter - 6, numLetter - 9}
	}
	//
	res = append(res, g.Kana[row][resf[0]+1], g.Kana[row][resf[1]+1], g.Kana[row][resf[2]+1], g.Kana[row][resf[3]+1],
		g.Kana[row][resf[0]+2], g.Kana[row][resf[1]+2], g.Kana[row][resf[2]+2], g.Kana[row][resf[3]+2])
	return res
}

func (g *Game) MainFontAn() {
	if g.MainFontAlphaFlag {
		if g.MainBtnFontAlpha == 40 {
			g.MainFontAlphaFlag = false
		}
		g.MainBtnFontAlpha -= 1
	} else {
		if g.MainBtnFontAlpha == 128 {
			g.MainFontAlphaFlag = true
		}
		g.MainBtnFontAlpha += 1
	}
}

func (g *Game) UpdateResultValues() {
	perm := rand.Perm(4) // случайный порядок [0..3]

	for i := 0; i < 4; i++ {
		idx := perm[i]

		g.ResBtn[i].Value = g.ResBtnValues[idx]
		g.ResBtn[i+4].Value = g.ResBtnValues[idx+4]
	}
	g.ShowHiraganaBTNS()
	g.ShowKatakanaBTNS()
}

func fillRoundedRect(screen *ebiten.Image, x, y, w, h, r float32, clr color.Color) {
	var path vector.Path

	// Начинаем сверху слева
	path.MoveTo(x+r, y)
	path.LineTo(x+w-r, y)
	path.Arc(x+w-r, y+r, r, -math.Pi/2, 0, vector.Clockwise)

	// Правая сторона
	path.LineTo(x+w, y+h-r)
	path.Arc(x+w-r, y+h-r, r, 0, math.Pi/2, vector.Clockwise)

	// Нижняя сторона
	path.LineTo(x+r, y+h)
	path.Arc(x+r, y+h-r, r, math.Pi/2, math.Pi, vector.Clockwise)

	// Левая сторона
	path.LineTo(x, y+r)
	path.Arc(x+r, y+r, r, math.Pi, 3*math.Pi/2, vector.Clockwise)

	path.Close()

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	rf, gf, bf, af := clr.RGBA()

	for i := range vs {
		vs[i].ColorR = float32(rf) / 0xffff
		vs[i].ColorG = float32(gf) / 0xffff
		vs[i].ColorB = float32(bf) / 0xffff
		vs[i].ColorA = float32(af) / 0xffff
	}
	screen.DrawTriangles(vs, is, whitePixel, nil)
}
