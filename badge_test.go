package badge

import (
	"testing"

	"github.com/valyala/fasttemplate"
)

func TestRender(t *testing.T) {
	fd, err := NewFace(11, 72, "fonts/verdana.ttf")
	if err != nil {
		t.Error(err)
	}
	tmpl := fasttemplate.New(FlatTemplate, "{{", "}}")
	Render("test", "status", ColorGray, fd, tmpl)
}

func BenchmarkRender(b *testing.B) {
	fd, err := NewFace(11, 72, "fonts/verdana.ttf")
	if err != nil {
		b.Error(err)
	}
	tmpl := fasttemplate.New(FlatTemplate, "{{", "}}")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Render("test", "status", ColorGray, fd, tmpl)
	}
}
