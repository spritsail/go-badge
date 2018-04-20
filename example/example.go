package main

import (
	"flag"

	"github.com/glaslos/go-badge"
	"github.com/glaslos/go-badge/fonts"
	"github.com/valyala/fasttemplate"
)

var (
	subject = flag.String("subject", "", "Badge subject")
	status  = flag.String("status", "", "Badge status")
	color   = flag.String("color", "blue", "Badge color")
)

func main() {
	flag.Parse()
	tmpl := fasttemplate.New(badge.FlatTemplate, "{{", "}}")
	fd, _ := badge.NewFaceStream(11, 72, fonts.Verdana)
	svg, err := badge.Render(*subject, *status, badge.Color(*color), fd, tmpl)
	if err != nil {
		panic(err)
	}
	print(svg)
}
