package main

import (
	"flag"

	"github.com/spritsail/go-badge"
)

var (
	subject = flag.String("subject", "", "Badge subject")
	status  = flag.String("status", "", "Badge status")
	color   = flag.String("color", "blue", "Badge color")
)

func main() {
	flag.Parse()
	svg, err := badge.RenderDef(*subject, *status, badge.Color(*color))
	if err != nil {
		panic(err)
	}
	print(svg)
}
