package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed KosugiMaru-Regular.ttf
var kosugiFontBytes []byte

var background *ebiten.Image
var whitePixel = ebiten.NewImage(1, 1)

func init() {
	whitePixel.Fill(color.White)

	file, err := os.Open("./assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	background = ebiten.NewImageFromImage(img)
}

func main() {
	//"★"
	kana := make(map[string][]string)
	kana["a"] = []string{
		"a", "あ", "ア",
		"i", "い", "イ",
		"u", "う", "ウ",
		"e", "え", "エ",
		"o", "お", "オ",
	}
	kana["ka"] = []string{
		"ka", "か", "カ",
		"ki", "き", "キ",
		"ku", "く", "ク",
		"ke", "け", "ケ",
		"ko", "こ", "コ",
	}
	//ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	fontReader := bytes.NewReader(kosugiFontBytes)
	textSource, err := text.NewGoTextFaceSource(fontReader)
	if err != nil {
		log.Fatal("Ошибка загрузки шрифта:", err)
	}

	game := &Game{
		Width:  640,
		Height: 640,

		SContNum:          0,
		MainBtnFontAlpha:  128,
		MainFontAlphaFlag: true,

		Kana:          kana,
		KanaConfTaken: []string{"a", "ka"},

		UserRow: make([]string, 4),

		BackgroundConf: BackgroundConfig{
			Alpha: 0.6,
		},
		ObjectsConf: ObjectsConfig{
			//12
			Pad: 16,
			Clr: color.NRGBA{255, 183, 197, 32},
		},

		Source: textSource,
		TopFontface: &text.GoTextFace{
			Source: textSource,
			Size:   24,
		},
		MainFontface: &text.GoTextFace{
			Source: textSource,
			Size:   24,
		},

		MainFontfaceLow: &text.GoTextFace{
			Source: textSource,
			Size:   96,
		},
		ResultsFontface: &text.GoTextFace{
			Source: textSource,
			Size:   32,
		},
		CheckFontface: &text.GoTextFace{
			Source: textSource,
			Size:   28,
		},
		MBtn:      make([]MainBtn, 0),
		RBtn:      make([]RowBtn, 0),
		ResBtn:    make([]ResultBtn, 0),
		CheckBtns: make([]CheckButton, 0),
	}

	game.CheckBtns = append(game.CheckBtns, CheckButton{
		Value:      "view",
		X:          128,
		Y:          384,
		H:          54,
		W:          128,
		Clr:        []uint8{202, 111, 247, 128},
		Visibility: true,
	})

	game.MBtn = append(game.MBtn, MainBtn{
		MainID:     -1,
		X:          128,
		Y:          256,
		W:          128,
		H:          128,
		Clr:        []uint8{255, 183, 197, 128},
		Visibility: true,
	})

	game.MBtn = append(game.MBtn, MainBtn{
		MainID:     0,
		X:          256,
		Y:          256,
		W:          128,
		H:          128,
		Clr:        []uint8{255, 183, 197, 128},
		Visibility: true,
	})

	game.MBtn = append(game.MBtn, MainBtn{
		MainID:     1,
		X:          384,
		Y:          256,
		W:          128,
		H:          128,
		Clr:        []uint8{255, 183, 197, 128},
		Visibility: true,
	})

	for i := 0; i < 9; i++ {
		//topLavel := 500
		//if i > 4 {fl = truehh = 500}
		allRows := []string{
			"a", "ka", "sa", "ta", "na", "ha", "ma", "ra", "y+wa",
		}
		game.RBtn = append(game.RBtn, RowBtn{
			Id:         i + 1,
			Value:      allRows[i],
			W:          54,
			H:          40,
			Taken:      false,
			Clr:        []uint8{255, 183, 197, 240},
			Visibility: true,
		})
	}

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     0,
		Id:         0,
		Y:          256,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     0,
		Id:         1,
		Y:          256,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     0,
		Id:         0,
		Y:          256 + 64,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     0,
		Id:         1,
		Y:          256 + 64,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     1,
		Id:         1,
		Y:          256,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     1,
		Id:         0,
		Y:          256,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     1,
		Id:         1,
		Y:          256 + 64,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     1,
		Id:         0,
		Y:          256 + 64,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.RBtn[0].Taken = true
	game.RBtn[1].Taken = true

	game.ActiveRow = game.GetNewActiveRow()
	game.ResBtnValues = game.GetNewResBtnValues()
	game.UpdateResBtnValues()

	ebiten.SetWindowSize(game.Width, game.Height)
	ebiten.SetWindowTitle("Kana trainer")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
