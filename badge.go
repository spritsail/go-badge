package badge

import (
	"io/ioutil"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/golang/freetype/truetype"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/svg"
	"github.com/valyala/fasttemplate"
	"golang.org/x/image/font"
)

// SVG minifier instance
var svgMin *minify.M

const (
	dpi      = 72
	fontsize = 11
	extraDx  = 10
)

func measureString(s string, face font.Face) float64 {
	sm := font.MeasureString(face, s)
	// this 64 is weird but it's the way I've found how to convert fixed.Int26_6 to float64
	return float64(sm)/64 + extraDx
}

// Render renders a badge of the given color, with given subject and status to w.
func Render(subject, status string, color Color, fd font.Face, tmpl *fasttemplate.Template) (svg string, err error) {
	subjectDx := measureString(subject, fd)
	statusDx := measureString(status, fd) - 2

	data := map[string]interface{}{
		"subject":   subject,
		"status":    status,
		"color":     color.String(),
		"dx":        humanize.Ftoa(subjectDx + statusDx),
		"subjectDx": humanize.Ftoa(subjectDx),
		"subjectX":  humanize.Ftoa(subjectDx / 2.0),
		"statusDx":  humanize.Ftoa(statusDx),
		"statusX":   humanize.Ftoa(subjectDx + statusDx/2.0 - 1),
	}

	svg, err = svgMin.String("image/svg+xml", tmpl.ExecuteString(data))
	return
}

// NewFace creates a new face based on font, size and dpi
func NewFace(size, dpi float64, fontPath string) (face font.Face, err error) {
	f, err := os.Open(fontPath)
	if err != nil {
		return
	}
	defer f.Close()
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	return NewFaceStream(size, dpi, raw)
}

// NewFaceStream creates a new face based on font bytes, size and dpi
func NewFaceStream(size, dpi float64, raw []byte) (face font.Face, err error) {
	ttf, err := truetype.Parse(raw)
	if err != nil {
		return
	}
	return truetype.NewFace(ttf, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}), nil
}

func init() {
	svgMin = minify.New()
	svgMin.AddFunc("image/svg+xml", svg.Minify)
}
