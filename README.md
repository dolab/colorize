[![CircleCI](https://dl.circleci.com/status-badge/img/gh/dolab/colorize/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/dolab/colorize/tree/master) [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/dolab/colorize) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dolab/colorize/master/LICENSE)

# colorize

Simple golang command line colorize for outputs highlight.

# Install

```go
go get -u github.com/dolab/colorize
```

# Usage

- Standard

```go
import "github.com/dolab/colorize"

func main() {
    brush := colorize.New("yellow+bBuih:black+h")
    brush.Paint("Hello, colorize!")
}
```

- Customize

```go
import "github.com/dolab/colorize"

func main() {
    brush := colorize.New("yellow+bBuih:black+h")

    // gain colors
    colorDraw, colorClean := brush.Colour()

    // custom output
    buf := bytes.NewBufferString(colorDraw)
    buf.WriteString("Hello, colorize!")
    buf.WriteString(colorClean)

    buf.WriteTo(os.Stdout)
}
```


# Format

> `foreground_color`+`color_attributes`:`background_color`+`color_attributes`

## buildin colors

- ColorRed
- ColorGreen
- ColorYellow
- ColorBlue
- ColorMagenta
- ColorCyan
- ColorWhite
- ColorGray
- ColorBlack

## available attributes
- b = bold foreground
- B = blink foreground
- u = underline foreground
- i = inverse
- h = high intensity (bright) foreground, background

# Windows support?

see [go-colorable](https://github.com/mattn/go-colorable)

# License

MIT

# Author

Spring MC
