package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type BackgroundConfig struct {
	Alpha float32
}
type ObjectsConfig struct {
	Pad float32
	Clr color.Color
}
type Game struct {
	Width  int
	Height int

	SContNum          int
	MainBtnFontAlpha  uint8
	MainFontAlphaFlag bool

	MBtn      []MainBtn
	RBtn      []RowBtn
	ResBtn    []ResultBtn
	CheckBtns []CheckButton

	Kana          map[string][]string // вся кана , ключ -- название ряда
	KanaConfTaken []string            // выбранные пользователем ряды

	ActiveRow    []string
	UserRow      []string
	ResBtnValues []string

	BackgroundConf BackgroundConfig
	ObjectsConf    ObjectsConfig

	Source *text.GoTextFaceSource

	TopFontface      *text.GoTextFace
	MainFontface     *text.GoTextFace
	MainFontfaceLow  *text.GoTextFace
	MainFontfaceLow3 *text.GoTextFace
	ResultsFontface  *text.GoTextFace
	CheckFontface    *text.GoTextFace
}

func (g *Game) View() {
	if g.UserRow[1] != g.ActiveRow[1] {
		g.MBtn[2].Status = "error"
	}
	if g.UserRow[2] != g.ActiveRow[2] {
		g.MBtn[1].Status = "error"
	}
	if g.UserRow[1] == g.ActiveRow[1] {
		g.MBtn[2].Status = "good"
	}
	if g.UserRow[2] == g.ActiveRow[2] {
		g.MBtn[1].Status = "good"
	}
	g.CheckBtns[0].Value = "next"
}

func (g *Game) Next() {
	num := 0
	if g.UserRow[1] != g.ActiveRow[1] {
		num -= len(g.KanaConfTaken)
	}
	if g.UserRow[2] != g.ActiveRow[2] {
		num -= len(g.KanaConfTaken)
	}
	if g.UserRow[1] == g.ActiveRow[1] {
		num += len(g.KanaConfTaken)
	}
	if g.UserRow[2] == g.ActiveRow[2] {
		num += len(g.KanaConfTaken)
	}

	g.SContNum += num

	g.MBtn[1].Status = "normal"
	g.MBtn[2].Status = "normal"
	g.ActiveRow = g.GetNewActiveRow()
	g.ResBtnValues = g.GetNewResBtnValues()
	g.UpdateResultValues()
	g.CheckBtns[0].Value = "view"
}

func (g *Game) KanaConfAdd(row string) {
	g.KanaConfTaken = append(g.KanaConfTaken, row)
}

func (g *Game) KanaConfDelete(row string) {
	for i := 0; i < len(g.KanaConfTaken); i++ {
		if row == g.KanaConfTaken[i] {
			left := g.KanaConfTaken[:i]
			right := g.KanaConfTaken[i+1:]
			g.KanaConfTaken = append(left, right...)
		}
	}
}

func (g *Game) UpdateResBtnValues() {
	for i, _ := range g.ResBtn {
		g.ResBtn[i].Value = g.ResBtnValues[i]
	}
}

func (g *Game) IsHoveredRBtns(cur_x, cur_y float32) {
	for i, _ := range g.RBtn {
		g.RBtn[i].Hovered = g.RBtn[i].Contains(cur_x, cur_y)
	}
}

func (g *Game) IsHoveredCheckBtns(cur_x, cur_y float32) {
	for i, _ := range g.CheckBtns {
		g.CheckBtns[i].Hovered = g.CheckBtns[i].Contains(cur_x, cur_y, g.ObjectsConf.Pad)
	}
}

func (g *Game) IsHoveredResBtns(cur_x, cur_y float32) {
	for i, _ := range g.ResBtn {
		g.ResBtn[i].Hovered = g.ResBtn[i].Contains(cur_x, cur_y)
	}
}

func (g *Game) IsHoveredMBtns(cur_x, cur_y float32) {
	for i, _ := range g.MBtn {
		g.MBtn[i].Hovered = g.MBtn[i].Contains(cur_x, cur_y, g.ObjectsConf.Pad)
	}
}

func (g *Game) DrawBackground(screen *ebiten.Image, clr color.Color) {
	screen.Fill(clr)
	op := &ebiten.DrawImageOptions{}

	w := background.Bounds().Dx()
	h := background.Bounds().Dy()
	op.GeoM.Scale(
		float64(g.Width)/float64(w),
		float64(g.Height)/float64(h),
	)
	op.ColorScale.ScaleAlpha(g.BackgroundConf.Alpha)
	screen.DrawImage(background, op)
}

func (g *Game) DrawTopPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		0+g.ObjectsConf.Pad, 0+g.ObjectsConf.Pad,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 112,
		12, g.ObjectsConf.Clr,
	)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(36+g.ObjectsConf.Pad), float64(0+g.ObjectsConf.Pad)+42) // Координаты X, Y
	text.Draw(screen, fmt.Sprintf("SCORE: %d", g.SContNum), g.TopFontface, op)
}

func (g *Game) DrawMidPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		g.ObjectsConf.Pad, 128+g.ObjectsConf.Pad,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 336,
		12, g.ObjectsConf.Clr,
	)
}

func (g *Game) DrawBotPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		g.ObjectsConf.Pad, 464+g.ObjectsConf.Pad*2,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 48,
		12, g.ObjectsConf.Clr,
	)
	fillRoundedRect(
		screen,
		g.ObjectsConf.Pad, 464+48+g.ObjectsConf.Pad*2+6,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 48,
		12, g.ObjectsConf.Clr,
	)
}

func (g *Game) DrawMainBtn(screen *ebiten.Image, btn *MainBtn) {

	fillRoundedRect(screen, btn.X+(g.ObjectsConf.Pad*float32(btn.MainID))*2, btn.Y, btn.W, btn.H, 12,
		color.NRGBA{btn.Clr[0], btn.Clr[1], btn.Clr[2], btn.Clr[3]})
	op := &text.DrawOptions{}
	if g.ActiveRow[3] == "a" && btn.MainID == -1 {
		op.GeoM.Translate(float64(btn.X)+float64(g.ObjectsConf.Pad)*float64(btn.MainID)*2+40, 256+12)
	} else {
		if len(g.UserRow[0]) == 3 {
			if btn.MainID == -1 {
				op.GeoM.Translate(float64(btn.X)+float64(g.ObjectsConf.Pad)*float64(btn.MainID)*2+20, 256+32)
			} else {
				op.GeoM.Translate(float64(btn.X)+float64(g.ObjectsConf.Pad)*float64(btn.MainID)*2+16, 256+12)
			}
		} else {
			op.GeoM.Translate(float64(btn.X)+float64(g.ObjectsConf.Pad)*float64(btn.MainID)*2+16, 256+12)
		}
	}
	if btn.MainID == -1 {
		if len(g.UserRow[0]) == 3 {
			text.Draw(screen, g.UserRow[0], g.MainFontfaceLow3, op)
		} else {
			text.Draw(screen, g.UserRow[0], g.MainFontfaceLow, op)
		}
	} else if btn.MainID == 1 {
		text.Draw(screen, g.UserRow[1], g.MainFontfaceLow, op)
	} else {
		text.Draw(screen, g.UserRow[2], g.MainFontfaceLow, op)
	}
}

func (g *Game) DrawResultBtn(screen *ebiten.Image, btn *ResultBtn) {

	if btn.MainID == 0 {
		btn.X = g.MBtn[0].X - g.ObjectsConf.Pad*2 + 64*float32(btn.Id) + 128 + 32
	} else if btn.MainID == 1 {
		btn.X = g.MBtn[2].X + g.ObjectsConf.Pad*2 + 64*float32(btn.Id)
	}

	btn.W = g.MBtn[0].W / 2
	btn.H = g.MBtn[0].H / 2
	fillRoundedRect(screen, btn.X, btn.Y, btn.W, btn.H, 12, color.NRGBA{btn.Clr[0], btn.Clr[1], btn.Clr[2], btn.Clr[3]})

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(btn.X)+16, float64(btn.Y)+12)
	text.Draw(screen, btn.Value, g.ResultsFontface, op)
}

func (g *Game) DrawRowBtn(screen *ebiten.Image, btn *RowBtn) {
	var clr color.Color
	botLavel := 500 + 40 + g.ObjectsConf.Pad - 2
	topLavel := 500
	if btn.Taken {
		clr = color.NRGBA{74, 247, 143, btn.Clr[3]}
		btn.Y = float32(topLavel)
		if btn.Hovered {
			btn.Y += 2
		}
	} else {
		clr = color.NRGBA{255, 183, 197, btn.Clr[3]} //240
		btn.Y = float32(botLavel)
		if btn.Hovered {
			btn.Y -= 2
		}
	}

	btn.X = float32(g.Width) - g.ObjectsConf.Pad*2 - (btn.W * float32(btn.Id)) - (float32(btn.Id) * g.ObjectsConf.Pad / 2)
	fillRoundedRect(screen, btn.X, btn.Y, btn.W, btn.H, 12, clr)
	op := &text.DrawOptions{}
	if btn.Id == 1 {
		op.GeoM.Translate(float64(btn.X+20), float64(btn.Y+6))
	} else if btn.Id == 9 {
		op.GeoM.Translate(float64(btn.X+2), float64(btn.Y+6))
	} else {
		op.GeoM.Translate(float64(btn.X+14), float64(btn.Y+6))
	}
	text.Draw(screen, btn.Value, g.MainFontface, op)
}

func (g *Game) DrawCheckButton(screen *ebiten.Image, btn *CheckButton) {
	x := btn.X + (g.ObjectsConf.Pad * (-2))
	y := btn.Y + g.ObjectsConf.Pad/2
	if btn.Hovered {
		x = btn.X + (g.ObjectsConf.Pad * (-2)) + 2
		y = btn.Y + g.ObjectsConf.Pad/2 + 2
	}

	fillRoundedRect(screen, x, y, btn.W, btn.H,
		12, color.NRGBA{btn.Clr[0], btn.Clr[1], btn.Clr[2], btn.Clr[3]})
	padding := g.ObjectsConf.Pad * 2 * float32(-1)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x+padding)+70, float64(y)+12) // Координаты X, Y
	text.Draw(screen, btn.Value, g.CheckFontface, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Update() error {
	cur_x, cur_y := ebiten.CursorPosition()
	cur_hovered := false
	g.IsHoveredRBtns(float32(cur_x), float32(cur_y))
	g.IsHoveredMBtns(float32(cur_x), float32(cur_y))
	g.IsHoveredResBtns(float32(cur_x), float32(cur_y))
	g.IsHoveredCheckBtns(float32(cur_x), float32(cur_y))

	g.MainFontAn()

	for i, _ := range g.CheckBtns { // HOVER CHECK BUTTONS --- --- --- ---
		if g.CheckBtns[i].Visibility {
			if g.CheckBtns[i].Hovered {
				g.CheckBtns[i].Clr[3] = 220
				cur_hovered = true
			} else {
				g.CheckBtns[i].Clr[3] = g.MainBtnFontAlpha
			}
		}
	}

	for i, _ := range g.MBtn { // HOVER MAIN BUTTONS --- --- --- ---
		if g.MBtn[i].Visibility {
			if g.MBtn[i].Hovered {
				g.MBtn[i].Clr[3] = 128
				cur_hovered = true
			} else {
				if g.MBtn[i].Status == "error" {
					g.MBtn[i].Clr = []uint8{255, 0, 0, g.MainBtnFontAlpha}
				} else if g.MBtn[i].Status == "normal" {
					g.MBtn[i].Clr = []uint8{255, 183, 197, g.MainBtnFontAlpha}
				} else if g.MBtn[i].Status == "good" {
					g.MBtn[i].Clr = []uint8{0, 255, 0, g.MainBtnFontAlpha}
				}
			}
		}
	}

	for i, _ := range g.RBtn { // HOVER ROW BUTTONS --- --- --- ---
		if g.RBtn[i].Visibility {
			if g.RBtn[i].Hovered {
				g.RBtn[i].Clr[3] = 255
				cur_hovered = true
			} else {

				g.RBtn[i].Clr[3] = 240
			}
		}
	}

	for i, _ := range g.ResBtn { // HOVER RESULT BUTTONS --- --- --- ---
		if g.ResBtn[i].Visibility {
			if g.ResBtn[i].Hovered {
				g.ResBtn[i].Clr[3] = 128
			} else {
				g.ResBtn[i].Clr[3] = 0
			}
		}
	}

	for i, _ := range g.CheckBtns { // CLICK CHECK BUTTONS ~~~ ~~~ ~~~ ~~~
		if g.CheckBtns[i].Visibility {
			if g.CheckBtns[i].Clicked() {
				if g.CheckBtns[i].Value == "view" {
					g.View()
				} else if g.CheckBtns[i].Value == "next" {
					g.Next()
				}
			}
		}
	}

	for i, _ := range g.RBtn { // CLICK ROW BUTTONS ~~~ ~~~ ~~~ ~~~
		if g.RBtn[i].Visibility {
			if g.RBtn[i].Clicked() {
				fmt.Println("Clicked row: ", g.RBtn[i].Value, "taken:", g.RBtn[i].Taken)
				if g.RBtn[i].Taken {
					g.RBtn[i].Taken = false
					g.KanaConfDelete(g.RBtn[i].Value)
				} else {
					g.RBtn[i].Taken = true
					g.KanaConfAdd(g.RBtn[i].Value)
				}
			}
		}
	}

	for i, _ := range g.MBtn { // CLICK MAIN BUTTONS ~~~ ~~~ ~~~ ~~~
		if g.MBtn[i].Visibility {
			if g.MBtn[i].Clicked() {
				if g.MBtn[i].MainID == -1 {
					//fmt.Println(g.ActiveRow)
					//fmt.Println(g.UserRow)
				}
			}
		}
	}

	for i, _ := range g.ResBtn { // CLICK RESULT BUTTONS ~~~ ~~~ ~~~ ~~~
		if g.ResBtn[i].Visibility {
			if g.ResBtn[i].Clicked() {
				if g.ResBtn[i].MainID == 0 {
					//fmt.Println("HIRAGANA : ", g.ResBtn[i].Value)
					g.UserRow[2] = g.ResBtn[i].Value
					g.HideHiraganaBTNS()
				}
				if g.ResBtn[i].MainID == 1 {
					//	fmt.Println("KATAKANA : ", g.ResBtn[i].Value)
					g.UserRow[1] = g.ResBtn[i].Value
					g.HideKatakanaBTNS()
				}
			}
		}
	}

	if cur_hovered {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.DrawBackground(screen, color.NRGBA{255, 255, 255, 255})

	g.DrawTopPanel(screen)
	g.DrawMidPanel(screen)
	g.DrawBotPanel(screen)

	for i, _ := range g.MBtn {
		if g.MBtn[i].Visibility {
			g.DrawMainBtn(screen, &g.MBtn[i])
		}
	}

	for i, _ := range g.ResBtn {
		if g.ResBtn[i].Visibility {
			g.DrawResultBtn(screen, &g.ResBtn[i])
		}
	}

	for i, _ := range g.RBtn {
		if g.RBtn[i].Visibility {
			g.DrawRowBtn(screen, &g.RBtn[i])
		}
	}

	for i, _ := range g.CheckBtns {
		if g.CheckBtns[i].Visibility {
			g.DrawCheckButton(screen, &g.CheckBtns[i])
		}
	}
}
