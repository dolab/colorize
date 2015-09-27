# colorize
Simple golang command line color formatting

# Install
```go
go get -u github.com/dolab/colorize
```

# Usage
```go
import "github.com/dolab/colorize"

colorBrush := colorize.New("yellow+bBuih:black+h")
colorBrush.Paint("Hello, colorize!")
```

# Format
> foreground_color+color_attributes:background_color+color_attributes

## colors
- ColorRed
- ColorGreen
- ColorYellow
- ColorBlue
- ColorMagenta
- ColorCyan
- ColorWhite
- ColorGray
- ColorBlack

## attributes
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
