package colors

import (
	"fmt"
	"image/color"
	"sort"
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
			"hex(ff,ff,ff,ff)",
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
			"rgba(0,255,0,255)",
			"hex(0,ff,0)",
			"hex(0,ff,0,ff)",
			"#00ff00",
			"#00ff00ff",
		}},
		{"blue", Blue, []string{
			"BLUE",
			"RGB(0,0,255)",
			"RGBA(0,0,255,255)",
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
	for _, tt := range tests {
		test := tt
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
	t.Logf("%s: %v", name, colors[NamedColor(name)])
	c, err := Parse(s)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case !c.Is(exp):
		t.Errorf("expected %v, got %v", exp, c)
	}
	var f func() string
	switch {
	case strings.Contains(s, "rgb("):
		f = c.AsRGB
	case strings.Contains(s, "rgba("):
		f = c.AsRGBA
	case strings.Contains(s, "hex("):
		f = c.AsHex
	case strings.Contains(s, "#"):
		f = c.AsWeb
	case isNamedColor(s):
		f = func() string {
			return string(c.NamedColor)
		}
	default:
		return
	}
	t.Logf("got: %s", f())
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
		"rgb(,,)",
		"rgba()",
		"rgba(0)",
		"rgba(0,0)",
		"rgba(0,0,a)",
		"rgba(0,0,0,a)",
		"rgba(,,,)",
		"rgba(256,256,256,256)",
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
	z := MapString()
	m := make(map[string]color.Color)
	for k, v := range z {
		m["named/"+k] = v
	}

	// additional tests
	m["std/black"] = color.Black
	m["std/white"] = color.White
	m["std/transparent"] = color.Transparent
	m["std/opaque"] = color.Opaque

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, ss := range keys {
		s := ss
		t.Run(s, func(t *testing.T) {
			testInverse(t, s, FromColor(m[s]))
		})
	}
}

func testInverse(t *testing.T, name string, exp Color) {
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
