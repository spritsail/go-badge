package badge

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/valyala/fasttemplate"
	"golang.org/x/image/font"
)

const (
	dpi      = 72
	fontsize = 11
	extraDx  = 13
)

func measureString(s string, face font.Face) float64 {
	sm := font.MeasureString(face, s)
	// this 64 is weird but it's the way I've found how to convert fixed.Int26_6 to float64
	return float64(sm)/64 + extraDx
}

func floatStr(f float64) string {
	return strconv.Itoa(int(f))
}

// Render renders a badge of the given color, with given subject and status to w.
func Render(subject, status string, color Color, fd font.Face, tmpl *fasttemplate.Template) string {
	subjectDx := measureString(subject, fd)
	statusDx := measureString(status, fd)

	data := map[string]interface{}{
		"subject":   subject,
		"status":    status,
		"color":     color.String(),
		"dx":        floatStr(subjectDx + statusDx),
		"subjectDx": floatStr(subjectDx),
		"subjectX":  floatStr(subjectDx/2.0 + 1),
		"statusDx":  floatStr(statusDx),
		"statusX":   floatStr(subjectDx + statusDx/2.0 - 1),
	}
	return tmpl.ExecuteString(data)
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
