package fonts

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	_ "github.com/kyoh86/gigamoji/assets/assets"
	"github.com/rakyll/statik/fs"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	r := regexp.MustCompile(`-(\d+)\.ttf$`)
	s, err := fs.New()
	if err != nil {
		panic(err)
	}
	if err := fs.Walk(s, "/fonts", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		m := r.FindStringSubmatch(info.Name())
		if len(m) <= 1 {
			return nil
		}
		size := m[1]
		f, err := strconv.ParseFloat(size, 64)
		if err != nil {
			return err
		}
		fontMap[size] = fontInfo{
			fontface: info.Name(),
			fontsize: f,
		}
		sizes = append(sizes, size)
		return nil
	}); err != nil {
		panic(err)
	}
}

type fontInfo struct {
	fontface string
	fontsize float64
}

var fontMap = map[string]fontInfo{}
var sizes = []string{}

func SupportedSizes() []string {
	return sizes
}

func OpenFace(size string) (font.Face, error) {
	s, err := fs.New()
	if err != nil {
		return nil, err
	}
	info, ok := fontMap[size]
	if !ok {
		return nil, fmt.Errorf("unsupported size %s", size)
	}
	f, err := s.Open("/fonts/" + info.fontface)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(f)

	ft, err := truetype.Parse(buf)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(ft, &truetype.Options{Size: info.fontsize})
	return face, nil
}

func LineHeight(face font.Face, spacing fixed.Int26_6) fixed.Int26_6 {
	return fixed.Int26_6(int32(face.Metrics().Ascent) +
		int32(face.Metrics().Descent) -
		int32(fixed.I(1))).
		Mul(spacing)
}

type ParagraphScale struct {
	Height fixed.Int26_6
	Width  fixed.Int26_6
	Lines  []LinePosition
}

type LinePosition struct {
	Text string
	fixed.Point26_6
	width fixed.Int26_6
}

func ScaleParagraph(drawer font.Drawer, align string, text string) (s ParagraphScale) {
	lh := LineHeight(drawer.Face, fixed.Int26_6(80)) // 1.25 * 64 = 80 (int26_6)
	lines := strings.Split(text, "\n")
	s.Height = lh.Mul(fixed.I(len(lines)))
	ly := drawer.Face.Metrics().Ascent
	for _, l := range lines {
		w := drawer.MeasureString(l)
		if s.Width < w {
			s.Width = w
		}
		s.Lines = append(s.Lines, LinePosition{
			Text:      l,
			Point26_6: fixed.Point26_6{Y: ly},
			width:     w,
		})
		ly += lh
	}
	var xa fixed.Int26_6
	switch align {
	case "right":
		xa = 1
	case "center":
		xa = 2
	default:
		return
	}
	for i, l := range s.Lines {
		s.Lines[i].X = (s.Width - l.width) / xa
	}
	return
}
