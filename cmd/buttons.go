package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MainBtn struct {
	MainID     int
	X          float32
	Y          float32
	W          float32
	H          float32
	Value      string
	Hovered    bool
	Status     string
	Clr        []uint8
	Visibility bool
}

type ResultBtn struct {
	MainID     int
	Id         int
	Value      string
	X          float32
	Y          float32
	W          float32
	H          float32
	Hovered    bool
	Clr        []uint8
	Visibility bool
}

type RowBtn struct {
	Id         int
	Value      string
	X          float32
	Y          float32
	W          float32
	H          float32
	Taken      bool
	Hovered    bool
	Clr        []uint8
	Visibility bool
}

type CheckButton struct {
	Value      string
	X          float32
	Y          float32
	W          float32
	H          float32
	Hovered    bool
	Clr        []uint8
	Visibility bool
}

func (b *RowBtn) Contains(x, y float32) bool {
	return (x >= b.X) && (x <= b.X+b.W) && (y >= b.Y) && (y <= b.Y+b.H)
}

func (b *MainBtn) Contains(x, y, pad float32) bool {
	padding := pad * 2 * float32(b.MainID)
	res := (x >= b.X+padding) && (x <= b.X+b.W+padding) && (y >= b.Y) && (y <= b.Y+b.H)

	return res
}

func (b *ResultBtn) Contains(x, y float32) bool {
	return (x >= b.X) && (x <= b.X+b.W) && (y >= b.Y) && (y <= b.Y+b.H)
}

func (b *MainBtn) Clicked() bool {
	return b.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (b *RowBtn) Clicked() bool {
	return b.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (b *CheckButton) Clicked() bool {
	return b.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (b *ResultBtn) Clicked() bool {
	return b.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (b *MainBtn) Show() {
	b.Visibility = true
}

func (b *MainBtn) Hide() {
	b.Visibility = false
}

func (b *RowBtn) Show() {
	b.Visibility = true
}

func (b *RowBtn) Hide() {
	b.Visibility = false
}

func (b *ResultBtn) Show() {
	b.Visibility = true
}

func (b *ResultBtn) Hide() {
	b.Visibility = false
}

func (g *Game) ShowHiraganaBTNS() {
	for i := 0; i < 4; i++ {
		g.ResBtn[i].Visibility = true
	}
}

func (g *Game) HideHiraganaBTNS() {
	for i := 0; i < 4; i++ {
		g.ResBtn[i].Visibility = false
	}
}

func (g *Game) HideKatakanaBTNS() {
	for i := 4; i < 8; i++ {
		g.ResBtn[i].Visibility = false
	}
}

func (g *Game) ShowKatakanaBTNS() {
	for i := 4; i < 8; i++ {
		g.ResBtn[i].Visibility = true
	}
}

func (b *CheckButton) Contains(x, y, pad float32) bool {
	padding := pad * 2 * float32(-1)
	res := (x >= b.X+padding) && (x <= b.X+b.W+padding) && (y >= b.Y) && (y <= b.Y+b.H)
	return res
}
