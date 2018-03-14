# go-badge [![GoDoc](https://godoc.org/github.com/glaslos/go-badge?status.svg)](https://godoc.org/github.com/glaslos/go-badge)

go-badge is a library to render shield badges to SVG.

## Installation

Using `go get`

```
go get github.com/glaslos/go-badge
```

## Usage

```go
import (
	"os"

	"github.com/glaslos/go-badge"
	"github.com/valyala/fasttemplate"
)

func main() {
	tmpl := fasttemplate.New(badge.FlatTemplate, "{{", "}}")
	fd, _ := badge.NewFace(11, 72, "fonts/vera.ttf")
	println(badge.Render("godoc", "reference", "#5272B4", fd, tmpl))
}
```

## License

MIT
