package colors

import (
	"fmt"
	"image/color"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kenshaw/snaker"
)

// Color is a color.
type Color struct {
	R, G, B, A uint8
	NamedColor NamedColor
}

// New creates a new color, and looks up the named color.
func New(r, g, b, a uint8) Color {
	return Color{r, g, b, a, lookup[mapKey(r, g, b, a)]}
}

// ToColor creates a color.
func ToColor(r, g, b, a uint8, name string) Color {
	return Color{r, g, b, a, NamedColor(name)}
}

// Parse parses a color.
func Parse(s string) (Color, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	for _, f := range []func(string) (Color, bool){
		FromWeb,
		FromName,
		FromRGB,
		FromRGBA,
		FromHex,
	} {
		if c, ok := f(s); ok {
			return c, nil
		}
	}
	return Color{}, ErrInvalidColor
}

// FromColor converts a standard [color.Color] to a color.
func FromColor(clr color.Color) Color {
	if c, ok := clr.(Color); ok {
		return c
	}
	c := color.RGBAModel.Convert(clr).(color.RGBA)
	return New(c.R, c.G, c.B, c.A)
}

// FromWeb converts a web string to a color.
//
// Cribbed from https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func FromWeb(s string) (c Color, ok bool) {
	c.A = 0xff
	n := len(s)
	if n == 0 || s[0] != '#' {
		return c, false
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		ok = false
		return 0
	}
	switch ok = true; n {
	case 9:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		ok = false
	}
	if ok {
		c.NamedColor = lookup[mapKey(c.R, c.G, c.B, c.A)]
	}
	return
}

// FromName converts a name to a color.
func FromName(s string) (Color, bool) {
	n := strings.ToLower(strings.TrimSpace(snaker.ForceCamelIdentifier(s)))
	if c, ok := colors[NamedColor(n)]; ok {
		return ToColor(c.R, c.G, c.B, c.A, n), true
	}
	return Color{}, false
}

// FromRGB converts a rgb string to a color.
func FromRGB(s string) (Color, bool) {
	return fromRE(s, rgbRE, parseDec)
}

// FromRGBA converts a rgba string to a color.
func FromRGBA(s string) (Color, bool) {
	return fromRE(s, rgbaRE, parseDec)
}

// FromHex converts a hex string to a color.
func FromHex(s string) (Color, bool) {
	return fromRE(s, hexRE, parseHex)
}

// UnmarshalText satisfies the [encoding.TextUnmarshaler] interface.
func (clr *Color) UnmarshalText(text []byte) error {
	var err error
	*clr, err = Parse(string(text))
	return err
}

// MarshalText satisfies the [encoding.TextMarshaler] interface.
func (clr *Color) MarshalText() ([]byte, error) {
	return []byte(clr.AsText()), nil
}

// Name returns the color's name.
func (c Color) Name() string {
	return string(c.NamedColor)
}

// RGBA satisfies the [color.Color] interface.
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// Light returns whether the color is a "light" color or not.
func (c Color) Light() bool {
	return math.Sqrt(
		0.299*math.Pow(float64(c.R), 2)+
			0.587*math.Pow(float64(c.G), 2)+
			0.114*math.Pow(float64(c.B), 2),
	) > 130
}

// Dark returs whether or not the color is a "dark" color or not.
func (c Color) Dark() bool {
	return !c.Light()
}

// Is returns true when the colors are equivalent.
func (c Color) Is(clr color.Color) bool {
	b := color.RGBAModel.Convert(clr).(color.RGBA)
	return c.R == b.R && c.G == b.G && c.B == b.B && c.A == b.A
}

// Format satisfies the [fmt.Formatter] interface.
func (c Color) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Flag('#') {
			_, _ = fmt.Fprintf(f, "{R:%d G:%d B:%d A:%d}", c.R, c.G, c.B, c.A)
		} else {
			_, _ = f.Write([]byte(c.AsText()))
		}
	case 's':
		_, _ = f.Write([]byte(c.AsText()))
	case 'd':
		_, _ = f.Write([]byte(c.AsRGB()))
	case 'a':
		_, _ = f.Write([]byte(c.AsRGBA()))
	case 'x':
		_, _ = f.Write([]byte(c.AsHex()))
	case 'n':
		_, _ = f.Write([]byte(c.Name()))
	case 'e':
		_, _ = f.Write([]byte(c.AsWeb()))
	}
}

// AsText returns a string representation of the color.
func (c Color) AsText() string {
	if c.NamedColor != "" {
		return string(c.NamedColor)
	}
	return c.AsWebShort()
}

// AsRGB returns the color formatted as a rgb string, ex: rgb(255,255,255).
func (c Color) AsRGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

// AsRGBA returns the color formatted as a rgba string, ex: rgba(255,255,255,255).
func (c Color) AsRGBA() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%d)", c.R, c.G, c.B, c.A)
}

// AsHex returns the color formatted as a web, ex: hex(aa,bb,cc,dd).
func (c Color) AsHex() string {
	if c.A != 0xff {
		return fmt.Sprintf("hex(%x,%x,%x,%x)", c.R, c.G, c.B, c.A)
	}
	return fmt.Sprintf("hex(%x,%x,%x)", c.R, c.G, c.B)
}

// AsWeb returns the color formatted as a web string, ex: #ffffffff.
func (c Color) AsWeb() string {
	if c.A != 0xff {
		return fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
	}
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// AsWebShort returns the shortest possible formatted web string, ex: #fff.
func (c Color) AsWebShort() string {
	if c.A != 0xff {
		return fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
	}
	if r, g, b := c.R>>4, c.G>>4, c.B>>4; c.R&0xf == r && c.G&0xf == g && c.B&0xf == b {
		return fmt.Sprintf("#%x%x%x", r, g, b)
	}
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// Pflag returns a [Pflag] wrapping the color, that can be used with various
// command-line packages, such as [cobra], and satisfies the [pflag.Value]
// interface.
//
// [cobra]: https://github.com/spf13/cobra
// [pflag.Value]: https://pkg.go.dev/github.com/spf13/pflag#Value
func (c *Color) Pflag() Pflag {
	return Pflag{c}
}

// Pflag wraps a color, for use with command-line packages, such as [cobra].
// Satisfies the [pflag.Value] interface.
//
// [cobra]: https://github.com/spf13/cobra
// [pflag.Value]: https://pkg.go.dev/github.com/spf13/pflag#Value
type Pflag struct {
	c *Color
}

// String satisfies the [pflag.Value] interface.
//
// [pflag.Value]: https://pkg.go.dev/github.com/spf13/pflag#Value
func (f Pflag) String() string {
	return f.c.AsText()
}

// Set satisfies the [pflag.Value] interface.
//
// [pflag.Value]: https://pkg.go.dev/github.com/spf13/pflag#Value
func (f Pflag) Set(s string) error {
	var err error
	*f.c, err = Parse(s)
	return err
}

// Type satisfies the [pflag.Value] interface.
//
// [pflag.Value]: https://pkg.go.dev/github.com/spf13/pflag#Value
func (Pflag) Type() string {
	return "color"
}

// Error is a error.
type Error string

// Error satisfies the [error] interface.
func (err Error) Error() string {
	return string(err)
}

const (
	// ErrInvalidColor is invalid color error.
	ErrInvalidColor Error = "invalid color"
)

// fromRE parses all regexp matches with f.
func fromRE(s string, re *regexp.Regexp, f func(s string) (uint8, bool)) (Color, bool) {
	m := re.FindStringSubmatch(s)
	n := len(m)
	if n != 4 && n != 5 {
		return Color{}, false
	}
	var c Color
	var ok bool
	if c.R, ok = f(m[1]); !ok {
		return Color{}, false
	}
	if c.G, ok = f(m[2]); !ok {
		return Color{}, false
	}
	if c.B, ok = f(m[3]); !ok {
		return Color{}, false
	}
	switch {
	case n == 5 && m[4] != "":
		if c.A, ok = f(m[4]); !ok {
			return Color{}, false
		}
	default:
		c.A = 0xff
	}
	c.NamedColor = lookup[mapKey(c.R, c.G, c.B, c.A)]
	return c, true
}

// parseDec parses a decimal number in s.
func parseDec(s string) (uint8, bool) {
	u, err := strconv.ParseUint(s, 10, 8)
	return uint8(u), err == nil
}

// parseHex parses a hex number in s.
func parseHex(s string) (uint8, bool) {
	u, err := strconv.ParseUint(s, 16, 8)
	return uint8(u), err == nil
}

// regexps.
var (
	rgbRE  = regexp.MustCompile(`(?i)^rgb\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)$`)
	rgbaRE = regexp.MustCompile(`(?i)^rgba\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)$`)
	hexRE  = regexp.MustCompile(`(?i)^hex\(\s*([0-9a-f]{1,2})\s*,\s*([0-9a-f]{1,2})\s*,\s*([0-9a-f]{1,2})\s*(?:,\s*([0-9a-f]{1,2})\s*)?\)$`)
	// webRE  = regexp.MustCompile(`(?i)^#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})?$`)
)
