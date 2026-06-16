package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed KosugiMaru-Regular.ttf
var kosugiFontBytes []byte

//go:embed background.png
var BackgroundPNG []byte

var background *ebiten.Image
var textSource *text.GoTextFaceSource
var whitePixel = ebiten.NewImage(1, 1)

func init() {
	whitePixel.Fill(color.White)

	img, err := png.Decode(bytes.NewReader(BackgroundPNG))
	if err != nil {
		log.Fatal("image load error: ", err)
	}
	background = ebiten.NewImageFromImage(img)

	fontReader := bytes.NewReader(kosugiFontBytes)
	textSource, err = text.NewGoTextFaceSource(fontReader)
	if err != nil {
		log.Fatal("font load error: ", err)
	}
}

func main() {
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
	kana["sa"] = []string{
		"sa", "さ", "サ",
		"shi", "し", "シ",
		"su", "す", "ス",
		"se", "せ", "セ",
		"so", "そ", "ソ",
	}
	kana["ta"] = []string{
		"ta", "た", "タ",
		"chi", "ち", "チ",
		"tsu", "つ", "ツ",
		"te", "て", "テ",
		"to", "と", "ト",
	}
	kana["na"] = []string{
		"na", "な", "ナ",
		"ni", "に", "ニ",
		"nu", "ぬ", "ヌ",
		"ne", "ね", "ネ",
		"no", "の", "ノ",
	}
	kana["ha"] = []string{
		"ha", "は", "ハ",
		"hi", "ひ", "ヒ",
		"fu", "ふ", "フ",
		"he", "へ", "ヘ",
		"ho", "ほ", "ホ",
	}
	kana["ma"] = []string{
		"ma", "ま", "マ",
		"mi", "み", "ミ",
		"mu", "む", "ム",
		"me", "め", "メ",
		"mo", "も", "モ",
	}
	kana["y+wa"] = []string{
		"ya", "や", "ヤ",
		"yu", "ゆ", "ユ",
		"yo", "よ", "ヨ",
		"wa", "わ", "ワ",
		"wo", "を", "ヲ",
	}
	kana["ra"] = []string{
		"ra", "ら", "ラ",
		"ri", "り", "リ",
		"ru", "る", "ル",
		"re", "れ", "レ",
		"ro", "ろ", "ロ",
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
		MainFontfaceLow3: &text.GoTextFace{
			Source: textSource,
			Size:   60,
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
		Status:     "normal",
	})

	game.MBtn = append(game.MBtn, MainBtn{
		MainID:     0,
		X:          256,
		Y:          256,
		W:          128,
		H:          128,
		Clr:        []uint8{255, 183, 197, 128},
		Visibility: true,
		Status:     "normal",
	})

	game.MBtn = append(game.MBtn, MainBtn{
		MainID:     1,
		X:          384,
		Y:          256,
		W:          128,
		H:          128,
		Clr:        []uint8{255, 183, 197, 128},
		Visibility: true,
		Status:     "normal",
	})

	for i := 0; i < 9; i++ {
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
		Y:          320,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     0,
		Id:         1,
		Y:          320,
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
		Y:          320,
		Clr:        []uint8{255, 183, 197, 0},
		Visibility: true,
	})

	game.ResBtn = append(game.ResBtn, ResultBtn{
		MainID:     1,
		Id:         0,
		Y:          320,
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
