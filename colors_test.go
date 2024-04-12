package colors

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/kenshaw/snaker"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		exp  color.Color
		v    []string
	}{
		{"black", color.Black, []string{
			"black",
			"rgb(0,0,0)",
			"rgba(0,0,0,255)",
			"hex(00,00,00)",
			"hex(00,00,00,ff)",
			"#000000",
			"#000000ff",
		}},
		{"white", color.White, []string{
			"white",
			"rgb(255,255,255)",
			"rgba(255,255,255,255)",
			"hex(ff,ff,ff)",
			"hex( ff , ff, ff, ff)",
			"#ffffff",
			"#ffffffff",
		}},
		{"red", Red, []string{
			"red",
			"rgb(255,0,0)",
			"rgba(255,0,0,255)",
			"hex(ff,0,0)",
			"hex(ff,0,0,ff)",
			"#ff0000",
			"#ff0000ff",
		}},
		{"lime", Lime, []string{
			"lime",
			"rgb(0,255,0)",
			"rgba( 0, 255, 0, 255)",
			"hex(0,ff,0)",
			"hex(0,ff,0,ff)",
			"#00ff00",
			"#00ff00ff",
		}},
		{"blue", Blue, []string{
			"BLUE",
			"RGB(0,0,255)",
			" RGBA( 0, 0,255,255)",
			"HEX(0,0,FF)",
			"HEX(0,0,FF,FF)",
			"#0000FF",
			"#0000FFFF",
		}},
		{"mistyrose", Mistyrose, []string{ // {0xff, 0xe4, 0xe1, 0xff} rgb(255, 228, 225)
			"  Misty_Rose  ",
			"rgb(255, 228  ,   225)",
			"   rgba(255, 228, 225,   255  )",
			"hex(ff,e4,e1  )",
			"  hex(ff,  e4,e1,ff)",
			"#ffe4e1   ",
			"  #ffe4e1ff  ",
		}},
		{"indianred", Indianred, []string{ // {0xcd, 0x5c, 0x5c, 0xff} rgb(205, 92, 92)
			"indIAN_red",
			"Rgb(205,92,92)",
			"Rgba(205,92,92,255)",
			"Hex(cd,5c,5c)",
			"Hex(cd,5c,5c,fF)",
			"#CD5c5c",
			"#CD5c5cfF",
		}},
	}
	m := allColors(false)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		c := color.NRGBAModel.Convert(m[key]).(color.NRGBA)
		v := []string{
			key[strings.LastIndexByte(key, '/')+1:],
			fmt.Sprintf("rgba( %d, %d, %d, %d )", c.R, c.G, c.B, c.A),
			fmt.Sprintf("hex( %x, %x, %x, %x )", c.R, c.G, c.B, c.A),
		}
		tests = append(tests, struct {
			name string
			exp  color.Color
			v    []string
		}{
			key, c, v,
		})
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, ss := range test.v {
				s := ss
				t.Run(s, func(t *testing.T) {
					testParse(t, test.name, s, test.exp)
				})
			}
		})
	}
}

func testParse(t *testing.T, name, s string, exp color.Color) {
	t.Helper()
	if i := strings.LastIndexByte(name, '/'); i != 0 {
		name = name[i+1:]
	}
	t.Logf("%s: %v", name, colors[NamedColor(name)])
	c, err := Parse(s)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case !c.Is(exp):
		t.Errorf("expected %v, got %v", exp, c)
	}
	t.Logf("color: %#v light: %t dark: %t", c, c.Light(), c.Dark())
	var f func() string
	var n string
	switch z := strings.ToLower(s); {
	case strings.Contains(z, "rgb("):
		n, f = "AsRGB", c.AsRGB
	case strings.Contains(z, "rgba("):
		n, f = "AsRGBA", c.AsRGBA
	case strings.Contains(z, "hex("):
		n, f = "AsHex", c.AsHex
	case strings.Contains(z, "#"):
		n, f = "AsWeb", c.AsWeb
	case isNamedColor(s):
		n, f = "NamedColor", func() string {
			return string(c.NamedColor)
		}
	default:
		t.Fatalf("invalid test %q", s)
	}
	t.Logf("%s: %s", n, f())
}

func TestBad(t *testing.T) {
	tests := []string{
		"",
		" ",
		"Unknown",
		"rgb()",
		"rgb(0)",
		"rgb(0,0)",
		"rgb(0,0,a)",
		"rgb(256,256,256)",
		"rgb(100,100,100,50)",
		"rgb(,,)",
		"rgba()",
		"rgba(0)",
		"rgba(0,0)",
		"rgba(0,0,a)",
		"rgba(0,0,0,a)",
		"rgba(,,,)",
		"rgba(256,256,256,256)",
		"rgba(100,100,100,)",
		"hex()",
		"hex(0)",
		"hex(0,0)",
		"hex(0,0,z)",
		"hex(0,0,0,z)",
		"#",
		"#a",
		"#ab",
		"#coo",
		"#bada",
		"#coolao",
		"__",
	}
	for i, s := range tests {
		if _, err := Parse(s); err == nil {
			t.Errorf("test %d %q expected error", i, s)
		}
	}
}

func TestInverse(t *testing.T) {
	m := allColors(true)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			testInverse(t, FromColor(m[key]))
		})
	}
}

func testInverse(t *testing.T, exp Color) {
	t.Helper()
	tests := []func() string{
		func() string {
			return fmt.Sprintf("%#v", exp)
		},
		func() string {
			c := color.NRGBAModel.Convert(exp).(color.NRGBA)
			return fmt.Sprintf("nrgba{R:%d G:%d B:%d A:%d}", c.R, c.G, c.B, c.A)
		},
		exp.AsRGB,
		exp.AsRGBA,
		exp.AsHex,
		exp.Name,
		exp.AsWeb,
	}
	for i, f := range tests {
		s := f()
		t.Logf("%d: %q", i, s)
		switch {
		case i < 2:
			continue
		case s == "" || (i == 2 && exp.A != 0xff):
			t.Log("  skipping")
			continue
		}
		check(t, s, exp)
	}
	t.Log("--")
	for _, verb := range []rune{'s', 'v', 'd', 'a', 'x', 'n', 'e'} {
		s := fmt.Sprintf("%"+string(verb), exp)
		switch {
		case verb == 'd' && exp.A != 0xff:
			t.Logf("%c: skipping %q", verb, s)
			continue
		case s == "":
			t.Logf("%c: is empty!", verb)
			continue
		}
		t.Logf("%c: %q", verb, s)
		check(t, s, exp)
	}
}

func TestIs(t *testing.T) {
	tests := []struct {
		a, b color.Color
		exp  bool
	}{
		{Transparent, color.Transparent, true},
		{Black, color.Black, true},
		{White, color.Transparent, false},
		{color.RGBA{0, 0, 0, 0}, color.Transparent, true},
		{color.Opaque, color.Transparent, false},
		{color.Opaque, color.White, true},
		{color.Opaque, color.Black, false},
		{Blue, color.RGBA{0, 0, 255, 255}, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v := Is(test.a, test.b)
			t.Logf("%t", v)
			if v != test.exp {
				t.Errorf("expected %t, got: %t", test.exp, v)
			}
		})
	}
}

func allColors(opaque bool) map[string]color.Color {
	z := MapString()
	m := make(map[string]color.Color)
	for k, v := range z {
		m["named/"+k] = v
	}
	// additional tests
	m["std/black"] = color.Black
	m["std/white"] = color.White
	m["std/transparent"] = color.Transparent
	if opaque {
		m["std/opaque"] = color.Opaque
	}
	return m
}

func check(t *testing.T, s string, exp Color) {
	c, err := Parse(s)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case !c.Is(exp):
		t.Errorf("expected %v, got %v", exp, c)
	}
}

func isNamedColor(s string) bool {
	name := strings.ToLower(snaker.ForceCamelIdentifier(s))
	_, ok := colors[NamedColor(name)]
	return ok
}
