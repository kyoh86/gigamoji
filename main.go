package main

import (
	"fmt"
	"image"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/kyoh86/gigamoji/internal/fonts"
	"golang.org/x/image/font"
)

// nolint
var (
	version = "snapshot"
	commit  = "snapshot"
	date    = "snapshot"
)

func main() {
	app := kingpin.New("gigamoji", "generate emoji (like :+1: in slack) banner which support some bitmap font faces").Version(version).Author("kyoh86")

	var p struct {
		Size       string
		Text       string
		Foreground string
		Background string
		Align      string
	}

	app.Arg("text", "A text to bannerize").Required().StringVar(&p.Text)
	app.Flag("size", "Size of the charactors (by points)").Short('s').Envar("GIGAMOJI_SIZE").Default("8").EnumVar(&p.Size, fonts.SupportedSizes()...)
	app.Flag("foreground", "Font face string").Short('f').Envar("GIGAMOJI_FG").Default("\u25A0").StringVar(&p.Foreground)
	app.Flag("background", "Background string").Short('b').Envar("GIGAMOJI_BG").Default("\u25A1").StringVar(&p.Background)
	app.Flag("align", "Horizontal alignment of text").Short('a').Envar("GIGAMOJI_ALIGN").Default("left").EnumVar(&p.Align, "left", "center", "right")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	face, err := fonts.OpenFace(p.Size)
	if err != nil {
		panic(err)
	}
	defer face.Close()

	dr := font.Drawer{
		Src:  image.White,
		Face: face,
	}

	scale := fonts.ScaleParagraph(dr, p.Align, p.Text)

	img := image.NewRGBA(image.Rect(0, 0, scale.Width.Ceil(), scale.Height.Ceil()))
	dr.Dst = img

	for _, l := range scale.Lines {
		dr.Dot = l.Point26_6
		dr.DrawString(l.Text)
	}

	for y := 0; y < scale.Height.Ceil(); y++ {
		for x := 0; x < scale.Width.Ceil(); x++ {
			if img.RGBAAt(x, y).A > 124 {
				fmt.Print(p.Foreground)
			} else {
				fmt.Print(p.Background)
			}
		}
		fmt.Println()
	}
}
