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
		"Aa", "あ", "ア",
		"i", "い", "イ",
		"u", "う", "ウ",
		"e", "え", "エ",
		"o", "お", "オ",
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
		MRow: "a",
		MBtn: make([]MainBtn, 0),
		RBtn: make([]RowBtn, 0),
	}

	game.MBtn = append(game.MBtn, MainBtn{
		Id:   -1,
		X:    128,
		Y:    256,
		W:    128,
		H:    128,
		Text: "",
		Clr:  []uint8{255, 183, 197, 128},
	})

	game.MBtn = append(game.MBtn, MainBtn{
		Id:   0,
		X:    256,
		Y:    256,
		W:    128,
		H:    128,
		Text: "",
		Clr:  []uint8{255, 183, 197, 128},
	})

	game.MBtn = append(game.MBtn, MainBtn{
		Id:  1,
		X:   384,
		Y:   256,
		W:   128,
		H:   128,
		Clr: []uint8{255, 183, 197, 128},
	})

	for i := 0; i < 9; i++ {
		//topLavel := 500
		//if i > 4 {fl = truehh = 500}
		allRows := []string{
			"a", "ka", "sa", "ta", "na", "ha", "ma", "ra", "y+wa",
		}
		game.RBtn = append(game.RBtn, RowBtn{
			Id:    i + 1,
			Value: allRows[i],
			W:     54,
			H:     40,
			Taken: false,
			Clr:   []uint8{255, 183, 197, 240},
		})
	}

	game.RBtn[0].Taken = true
	game.RBtn[1].Taken = true

	ebiten.SetWindowSize(game.Width, game.Height)
	ebiten.SetWindowTitle("Kana trainer")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
