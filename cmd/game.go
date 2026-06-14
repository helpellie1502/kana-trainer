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

	MRow string
	MBtn []MainBtn
	RBtn []RowBtn

	Kana          map[string][]string
	KanaConfTaken []string

	BackgroundConf BackgroundConfig
	ObjectsConf    ObjectsConfig

	Source *text.GoTextFaceSource

	TopFontface     *text.GoTextFace
	MainFontface    *text.GoTextFace
	MainFontfaceLow *text.GoTextFace
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
func (g *Game) IsHoveredRBtns(cur_x, cur_y float32) {
	for i, _ := range g.RBtn {
		g.RBtn[i].Hovered = g.RBtn[i].Contains(cur_x, cur_y)
	}
}

func (g *Game) IsHoveredMBtns(cur_x, cur_y float32) {
	for i, _ := range g.MBtn {
		g.MBtn[i].Hovered = g.MBtn[i].Contains(cur_x, cur_y, g.ObjectsConf.Pad)
	}
}

func (g *Game) GenerateMainRowBtn(btns *[]MainBtn, row string) {
	for i := 0; i < len(*btns); i++ {
		(*btns)[i].Text = g.Kana[row][i]
	}
}

func (g *Game) DrawBackground(screen *ebiten.Image, clr color.Color, al float32) {
	screen.Fill(clr)
	op := &ebiten.DrawImageOptions{}

	w := background.Bounds().Dx()
	h := background.Bounds().Dy()
	op.GeoM.Scale(
		float64(g.Width)/float64(w),
		float64(g.Height)/float64(h),
	)

	op.ColorScale.ScaleAlpha(al)
	screen.DrawImage(background, op)
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
func (g *Game) DrawTopPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		0+g.ObjectsConf.Pad, 0+g.ObjectsConf.Pad,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 112,
		12, g.ObjectsConf.Clr,
	)

	op := &text.DrawOptions{}
	//op.GeoM.Translate(float64(g.ObjectsConf.Pad)*3, 54) // Координаты X, Y
	op.GeoM.Translate(float64(36+g.ObjectsConf.Pad), float64(0+g.ObjectsConf.Pad)+42) // Координаты X, Y
	text.Draw(screen, fmt.Sprintf("SCORE: %d", g.SContNum), g.TopFontface, op)
}

func (g *Game) DrawMidPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		0+g.ObjectsConf.Pad, 128+g.ObjectsConf.Pad,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 336,
		12, g.ObjectsConf.Clr,
	)
}

func (g *Game) DrawBotPanel(screen *ebiten.Image) {
	fillRoundedRect(
		screen,
		0+g.ObjectsConf.Pad, 464+g.ObjectsConf.Pad*2,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 48,
		12, g.ObjectsConf.Clr,
	)
	fillRoundedRect(
		screen,
		0+g.ObjectsConf.Pad, 464+48+g.ObjectsConf.Pad*2+6,
		float32(g.Width)-(g.ObjectsConf.Pad*2), 48,
		12, g.ObjectsConf.Clr,
	)
}

func (g *Game) DrawMainBtn(screen *ebiten.Image, btn *MainBtn) {

	fillRoundedRect(screen, btn.X+(g.ObjectsConf.Pad*float32(btn.Id))*2, btn.Y, btn.W, btn.H, 12,
		color.NRGBA{btn.Clr[0], btn.Clr[1], btn.Clr[2], btn.Clr[3]})
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(btn.X)+float64(g.ObjectsConf.Pad)*float64(btn.Id)*2+16, 256+12)
	text.Draw(screen, btn.Text, g.MainFontfaceLow, op)
}

func (g *Game) DrawRowBtn(screen *ebiten.Image, btn *RowBtn) {
	var clr color.Color
	botLavel := 500 + 40 + g.ObjectsConf.Pad - 2
	topLavel := 500
	if btn.Taken {
		clr = color.NRGBA{74, 247, 143, btn.Clr[3]}
		btn.Y = float32(topLavel)
		if btn.Hovered {
			//Println("Hovered", btn.Value)
			btn.Y += 2
		}
	} else {
		clr = color.NRGBA{255, 183, 197, btn.Clr[3]} //240
		btn.Y = float32(botLavel)
		if btn.Hovered {
			//Println("Hovered", btn.Value)
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Update() error {
	cur_x, cur_y := ebiten.CursorPosition()
	cur_hovered := false
	g.IsHoveredRBtns(float32(cur_x), float32(cur_y))
	g.IsHoveredMBtns(float32(cur_x), float32(cur_y))

	g.MainFontAn()
	g.GenerateMainRowBtn(&g.MBtn, "a")

	// ------------------------------------- hover
	for i, _ := range g.MBtn {
		if g.MBtn[i].Hovered {
			g.MBtn[i].Clr[3] = 128
			cur_hovered = true
		} else {
			g.MBtn[i].Clr[3] = g.MainBtnFontAlpha
		}
	}

	for i, _ := range g.RBtn {
		if g.RBtn[i].Hovered {
			g.RBtn[i].Clr[3] = 255
			cur_hovered = true
		} else {
			g.RBtn[i].Clr[3] = 240
		}
	}

	// ------------------------------------click

	for i, _ := range g.MBtn {
		if g.MBtn[i].Clicked() {
			fmt.Println("Clicked: ", g.MBtn[i].Text)
		}
	}

	for i, _ := range g.RBtn {
		if g.RBtn[i].Clicked() {
			fmt.Println("Clicked: ", g.RBtn[i].Value, "taken:", g.RBtn[i].Taken)
			if g.RBtn[i].Taken {
				g.RBtn[i].Taken = false
				g.KanaConfDelete(g.RBtn[i].Value)
			} else {
				g.RBtn[i].Taken = true
				g.KanaConfAdd(g.RBtn[i].Value)
			}
			fmt.Println(g.KanaConfTaken)
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

	g.DrawBackground(screen, color.NRGBA{255, 255, 255, 255}, g.BackgroundConf.Alpha)

	g.DrawTopPanel(screen)
	g.DrawMidPanel(screen)
	g.DrawBotPanel(screen)

	for i, _ := range g.MBtn {
		g.DrawMainBtn(screen, &g.MBtn[i])
	}

	for i, _ := range g.RBtn {
		g.DrawRowBtn(screen, &g.RBtn[i])
	}
}
