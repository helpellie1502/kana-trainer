package main

import "math/rand"

func (g *Game) GetNewActiveRow() []string {
	allRows := g.KanaConfTaken
	randNumRow := rand.Intn(len(allRows)) // 8
	randNumLetter := rand.Intn(4)
	//allRows := []string{"a", "ka", "sa", "ta", "na", "ha", "ma", "ra", "y+wa"}
	row := allRows[randNumRow] //5-6 a a a y y y i i i e e e o o o --- 0 3 6 9 12
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
