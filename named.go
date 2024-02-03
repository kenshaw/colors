package colors

import (
	"fmt"
	"image/color"
)

// Register registers an additional color.
func Register(n NamedColor, clr color.Color) {
	c := color.RGBAModel.Convert(clr).(color.RGBA)
	colors[n] = c
	lookup[mapKey(c.R, c.G, c.B, c.A)] = n
}

// RegisterName registers an additional color.
func RegisterName(s string, clr color.Color) {
	Register(NamedColor(s), clr)
}

// Map returns a map of all named colors.
func Map() map[NamedColor]Color {
	m := make(map[NamedColor]Color, len(colors))
	for k := range colors {
		m[k] = k.Color()
	}
	return m
}

// MapString returns a map of all named colors.
func MapString() map[string]Color {
	m := make(map[string]Color, len(colors))
	for k := range colors {
		m[string(k)] = k.Color()
	}
	return m
}

// NamedColor is a named color.
type NamedColor string

// Named colors.
const (
	Aliceblue            NamedColor = "aliceblue"
	Antiquewhite         NamedColor = "antiquewhite"
	Aqua                 NamedColor = "aqua"
	Aquamarine           NamedColor = "aquamarine"
	Azure                NamedColor = "azure"
	Beige                NamedColor = "beige"
	Bisque               NamedColor = "bisque"
	Black                NamedColor = "black"
	Blanchedalmond       NamedColor = "blanchedalmond"
	Blue                 NamedColor = "blue"
	Blueviolet           NamedColor = "blueviolet"
	Brown                NamedColor = "brown"
	Burlywood            NamedColor = "burlywood"
	Cadetblue            NamedColor = "cadetblue"
	Chartreuse           NamedColor = "chartreuse"
	Chocolate            NamedColor = "chocolate"
	Coral                NamedColor = "coral"
	Cornflowerblue       NamedColor = "cornflowerblue"
	Cornsilk             NamedColor = "cornsilk"
	Crimson              NamedColor = "crimson"
	Cyan                 NamedColor = "cyan"
	Darkblue             NamedColor = "darkblue"
	Darkcyan             NamedColor = "darkcyan"
	Darkgoldenrod        NamedColor = "darkgoldenrod"
	Darkgray             NamedColor = "darkgray"
	Darkgreen            NamedColor = "darkgreen"
	Darkgrey             NamedColor = "darkgrey"
	Darkkhaki            NamedColor = "darkkhaki"
	Darkmagenta          NamedColor = "darkmagenta"
	Darkolivegreen       NamedColor = "darkolivegreen"
	Darkorange           NamedColor = "darkorange"
	Darkorchid           NamedColor = "darkorchid"
	Darkred              NamedColor = "darkred"
	Darksalmon           NamedColor = "darksalmon"
	Darkseagreen         NamedColor = "darkseagreen"
	Darkslateblue        NamedColor = "darkslateblue"
	Darkslategray        NamedColor = "darkslategray"
	Darkslategrey        NamedColor = "darkslategrey"
	Darkturquoise        NamedColor = "darkturquoise"
	Darkviolet           NamedColor = "darkviolet"
	Deeppink             NamedColor = "deeppink"
	Deepskyblue          NamedColor = "deepskyblue"
	Dimgray              NamedColor = "dimgray"
	Dimgrey              NamedColor = "dimgrey"
	Dodgerblue           NamedColor = "dodgerblue"
	Firebrick            NamedColor = "firebrick"
	Floralwhite          NamedColor = "floralwhite"
	Forestgreen          NamedColor = "forestgreen"
	Fuchsia              NamedColor = "fuchsia"
	Gainsboro            NamedColor = "gainsboro"
	Ghostwhite           NamedColor = "ghostwhite"
	Gold                 NamedColor = "gold"
	Goldenrod            NamedColor = "goldenrod"
	Gray                 NamedColor = "gray"
	Green                NamedColor = "green"
	Greenyellow          NamedColor = "greenyellow"
	Grey                 NamedColor = "grey"
	Honeydew             NamedColor = "honeydew"
	Hotpink              NamedColor = "hotpink"
	Indianred            NamedColor = "indianred"
	Indigo               NamedColor = "indigo"
	Ivory                NamedColor = "ivory"
	Khaki                NamedColor = "khaki"
	Lavender             NamedColor = "lavender"
	Lavenderblush        NamedColor = "lavenderblush"
	Lawngreen            NamedColor = "lawngreen"
	Lemonchiffon         NamedColor = "lemonchiffon"
	Lightblue            NamedColor = "lightblue"
	Lightcoral           NamedColor = "lightcoral"
	Lightcyan            NamedColor = "lightcyan"
	Lightgoldenrodyellow NamedColor = "lightgoldenrodyellow"
	Lightgray            NamedColor = "lightgray"
	Lightgreen           NamedColor = "lightgreen"
	Lightgrey            NamedColor = "lightgrey"
	Lightpink            NamedColor = "lightpink"
	Lightsalmon          NamedColor = "lightsalmon"
	Lightseagreen        NamedColor = "lightseagreen"
	Lightskyblue         NamedColor = "lightskyblue"
	Lightslategray       NamedColor = "lightslategray"
	Lightslategrey       NamedColor = "lightslategrey"
	Lightsteelblue       NamedColor = "lightsteelblue"
	Lightyellow          NamedColor = "lightyellow"
	Lime                 NamedColor = "lime"
	Limegreen            NamedColor = "limegreen"
	Linen                NamedColor = "linen"
	Magenta              NamedColor = "magenta"
	Maroon               NamedColor = "maroon"
	Mediumaquamarine     NamedColor = "mediumaquamarine"
	Mediumblue           NamedColor = "mediumblue"
	Mediumorchid         NamedColor = "mediumorchid"
	Mediumpurple         NamedColor = "mediumpurple"
	Mediumseagreen       NamedColor = "mediumseagreen"
	Mediumslateblue      NamedColor = "mediumslateblue"
	Mediumspringgreen    NamedColor = "mediumspringgreen"
	Mediumturquoise      NamedColor = "mediumturquoise"
	Mediumvioletred      NamedColor = "mediumvioletred"
	Midnightblue         NamedColor = "midnightblue"
	Mintcream            NamedColor = "mintcream"
	Mistyrose            NamedColor = "mistyrose"
	Moccasin             NamedColor = "moccasin"
	Navajowhite          NamedColor = "navajowhite"
	Navy                 NamedColor = "navy"
	Oldlace              NamedColor = "oldlace"
	Olive                NamedColor = "olive"
	Olivedrab            NamedColor = "olivedrab"
	Orange               NamedColor = "orange"
	Orangered            NamedColor = "orangered"
	Orchid               NamedColor = "orchid"
	Palegoldenrod        NamedColor = "palegoldenrod"
	Palegreen            NamedColor = "palegreen"
	Paleturquoise        NamedColor = "paleturquoise"
	Palevioletred        NamedColor = "palevioletred"
	Papayawhip           NamedColor = "papayawhip"
	Peachpuff            NamedColor = "peachpuff"
	Peru                 NamedColor = "peru"
	Pink                 NamedColor = "pink"
	Plum                 NamedColor = "plum"
	Powderblue           NamedColor = "powderblue"
	Purple               NamedColor = "purple"
	Red                  NamedColor = "red"
	Rosybrown            NamedColor = "rosybrown"
	Royalblue            NamedColor = "royalblue"
	Saddlebrown          NamedColor = "saddlebrown"
	Salmon               NamedColor = "salmon"
	Sandybrown           NamedColor = "sandybrown"
	Seagreen             NamedColor = "seagreen"
	Seashell             NamedColor = "seashell"
	Sienna               NamedColor = "sienna"
	Silver               NamedColor = "silver"
	Skyblue              NamedColor = "skyblue"
	Slateblue            NamedColor = "slateblue"
	Slategray            NamedColor = "slategray"
	Slategrey            NamedColor = "slategrey"
	Snow                 NamedColor = "snow"
	Springgreen          NamedColor = "springgreen"
	Steelblue            NamedColor = "steelblue"
	Tan                  NamedColor = "tan"
	Teal                 NamedColor = "teal"
	Thistle              NamedColor = "thistle"
	Tomato               NamedColor = "tomato"
	Transparent          NamedColor = "transparent"
	Turquoise            NamedColor = "turquoise"
	Violet               NamedColor = "violet"
	Wheat                NamedColor = "wheat"
	White                NamedColor = "white"
	Whitesmoke           NamedColor = "whitesmoke"
	Yellow               NamedColor = "yellow"
	Yellowgreen          NamedColor = "yellowgreen"
)

// RGBA satisfies the [color.Color] interface.
func (c NamedColor) RGBA() (r, g, b, a uint32) {
	if c, ok := colors[c]; ok {
		return c.RGBA()
	}
	return
}

// Color returns the [Color].
func (c NamedColor) Color() Color {
	if v, ok := colors[c]; ok {
		return Color{v.R, v.G, v.B, v.A, c}
	}
	return Color{}
}

// Format satisfies the [fmt.Formatter] interface.
func (c NamedColor) Format(f fmt.State, verb rune) {
	c.Color().Format(f, verb)
}

// colors contains the named colors defined in the SVG 1.1 spec.
//
// Taken from golang.org/x/image/colornames/table.go
var colors = map[NamedColor]color.RGBA{
	Aliceblue:            {0xf0, 0xf8, 0xff, 0xff}, // rgb(240, 248, 255)
	Antiquewhite:         {0xfa, 0xeb, 0xd7, 0xff}, // rgb(250, 235, 215)
	Aqua:                 {0x00, 0xff, 0xff, 0xff}, // rgb(0, 255, 255)
	Aquamarine:           {0x7f, 0xff, 0xd4, 0xff}, // rgb(127, 255, 212)
	Azure:                {0xf0, 0xff, 0xff, 0xff}, // rgb(240, 255, 255)
	Beige:                {0xf5, 0xf5, 0xdc, 0xff}, // rgb(245, 245, 220)
	Bisque:               {0xff, 0xe4, 0xc4, 0xff}, // rgb(255, 228, 196)
	Black:                {0x00, 0x00, 0x00, 0xff}, // rgb(0, 0, 0)
	Blanchedalmond:       {0xff, 0xeb, 0xcd, 0xff}, // rgb(255, 235, 205)
	Blue:                 {0x00, 0x00, 0xff, 0xff}, // rgb(0, 0, 255)
	Blueviolet:           {0x8a, 0x2b, 0xe2, 0xff}, // rgb(138, 43, 226)
	Brown:                {0xa5, 0x2a, 0x2a, 0xff}, // rgb(165, 42, 42)
	Burlywood:            {0xde, 0xb8, 0x87, 0xff}, // rgb(222, 184, 135)
	Cadetblue:            {0x5f, 0x9e, 0xa0, 0xff}, // rgb(95, 158, 160)
	Chartreuse:           {0x7f, 0xff, 0x00, 0xff}, // rgb(127, 255, 0)
	Chocolate:            {0xd2, 0x69, 0x1e, 0xff}, // rgb(210, 105, 30)
	Coral:                {0xff, 0x7f, 0x50, 0xff}, // rgb(255, 127, 80)
	Cornflowerblue:       {0x64, 0x95, 0xed, 0xff}, // rgb(100, 149, 237)
	Cornsilk:             {0xff, 0xf8, 0xdc, 0xff}, // rgb(255, 248, 220)
	Crimson:              {0xdc, 0x14, 0x3c, 0xff}, // rgb(220, 20, 60)
	Cyan:                 {0x00, 0xff, 0xff, 0xff}, // rgb(0, 255, 255)
	Darkblue:             {0x00, 0x00, 0x8b, 0xff}, // rgb(0, 0, 139)
	Darkcyan:             {0x00, 0x8b, 0x8b, 0xff}, // rgb(0, 139, 139)
	Darkgoldenrod:        {0xb8, 0x86, 0x0b, 0xff}, // rgb(184, 134, 11)
	Darkgray:             {0xa9, 0xa9, 0xa9, 0xff}, // rgb(169, 169, 169)
	Darkgreen:            {0x00, 0x64, 0x00, 0xff}, // rgb(0, 100, 0)
	Darkgrey:             {0xa9, 0xa9, 0xa9, 0xff}, // rgb(169, 169, 169)
	Darkkhaki:            {0xbd, 0xb7, 0x6b, 0xff}, // rgb(189, 183, 107)
	Darkmagenta:          {0x8b, 0x00, 0x8b, 0xff}, // rgb(139, 0, 139)
	Darkolivegreen:       {0x55, 0x6b, 0x2f, 0xff}, // rgb(85, 107, 47)
	Darkorange:           {0xff, 0x8c, 0x00, 0xff}, // rgb(255, 140, 0)
	Darkorchid:           {0x99, 0x32, 0xcc, 0xff}, // rgb(153, 50, 204)
	Darkred:              {0x8b, 0x00, 0x00, 0xff}, // rgb(139, 0, 0)
	Darksalmon:           {0xe9, 0x96, 0x7a, 0xff}, // rgb(233, 150, 122)
	Darkseagreen:         {0x8f, 0xbc, 0x8f, 0xff}, // rgb(143, 188, 143)
	Darkslateblue:        {0x48, 0x3d, 0x8b, 0xff}, // rgb(72, 61, 139)
	Darkslategray:        {0x2f, 0x4f, 0x4f, 0xff}, // rgb(47, 79, 79)
	Darkslategrey:        {0x2f, 0x4f, 0x4f, 0xff}, // rgb(47, 79, 79)
	Darkturquoise:        {0x00, 0xce, 0xd1, 0xff}, // rgb(0, 206, 209)
	Darkviolet:           {0x94, 0x00, 0xd3, 0xff}, // rgb(148, 0, 211)
	Deeppink:             {0xff, 0x14, 0x93, 0xff}, // rgb(255, 20, 147)
	Deepskyblue:          {0x00, 0xbf, 0xff, 0xff}, // rgb(0, 191, 255)
	Dimgray:              {0x69, 0x69, 0x69, 0xff}, // rgb(105, 105, 105)
	Dimgrey:              {0x69, 0x69, 0x69, 0xff}, // rgb(105, 105, 105)
	Dodgerblue:           {0x1e, 0x90, 0xff, 0xff}, // rgb(30, 144, 255)
	Firebrick:            {0xb2, 0x22, 0x22, 0xff}, // rgb(178, 34, 34)
	Floralwhite:          {0xff, 0xfa, 0xf0, 0xff}, // rgb(255, 250, 240)
	Forestgreen:          {0x22, 0x8b, 0x22, 0xff}, // rgb(34, 139, 34)
	Fuchsia:              {0xff, 0x00, 0xff, 0xff}, // rgb(255, 0, 255)
	Gainsboro:            {0xdc, 0xdc, 0xdc, 0xff}, // rgb(220, 220, 220)
	Ghostwhite:           {0xf8, 0xf8, 0xff, 0xff}, // rgb(248, 248, 255)
	Gold:                 {0xff, 0xd7, 0x00, 0xff}, // rgb(255, 215, 0)
	Goldenrod:            {0xda, 0xa5, 0x20, 0xff}, // rgb(218, 165, 32)
	Gray:                 {0x80, 0x80, 0x80, 0xff}, // rgb(128, 128, 128)
	Green:                {0x00, 0x80, 0x00, 0xff}, // rgb(0, 128, 0)
	Greenyellow:          {0xad, 0xff, 0x2f, 0xff}, // rgb(173, 255, 47)
	Grey:                 {0x80, 0x80, 0x80, 0xff}, // rgb(128, 128, 128)
	Honeydew:             {0xf0, 0xff, 0xf0, 0xff}, // rgb(240, 255, 240)
	Hotpink:              {0xff, 0x69, 0xb4, 0xff}, // rgb(255, 105, 180)
	Indianred:            {0xcd, 0x5c, 0x5c, 0xff}, // rgb(205, 92, 92)
	Indigo:               {0x4b, 0x00, 0x82, 0xff}, // rgb(75, 0, 130)
	Ivory:                {0xff, 0xff, 0xf0, 0xff}, // rgb(255, 255, 240)
	Khaki:                {0xf0, 0xe6, 0x8c, 0xff}, // rgb(240, 230, 140)
	Lavender:             {0xe6, 0xe6, 0xfa, 0xff}, // rgb(230, 230, 250)
	Lavenderblush:        {0xff, 0xf0, 0xf5, 0xff}, // rgb(255, 240, 245)
	Lawngreen:            {0x7c, 0xfc, 0x00, 0xff}, // rgb(124, 252, 0)
	Lemonchiffon:         {0xff, 0xfa, 0xcd, 0xff}, // rgb(255, 250, 205)
	Lightblue:            {0xad, 0xd8, 0xe6, 0xff}, // rgb(173, 216, 230)
	Lightcoral:           {0xf0, 0x80, 0x80, 0xff}, // rgb(240, 128, 128)
	Lightcyan:            {0xe0, 0xff, 0xff, 0xff}, // rgb(224, 255, 255)
	Lightgoldenrodyellow: {0xfa, 0xfa, 0xd2, 0xff}, // rgb(250, 250, 210)
	Lightgray:            {0xd3, 0xd3, 0xd3, 0xff}, // rgb(211, 211, 211)
	Lightgreen:           {0x90, 0xee, 0x90, 0xff}, // rgb(144, 238, 144)
	Lightgrey:            {0xd3, 0xd3, 0xd3, 0xff}, // rgb(211, 211, 211)
	Lightpink:            {0xff, 0xb6, 0xc1, 0xff}, // rgb(255, 182, 193)
	Lightsalmon:          {0xff, 0xa0, 0x7a, 0xff}, // rgb(255, 160, 122)
	Lightseagreen:        {0x20, 0xb2, 0xaa, 0xff}, // rgb(32, 178, 170)
	Lightskyblue:         {0x87, 0xce, 0xfa, 0xff}, // rgb(135, 206, 250)
	Lightslategray:       {0x77, 0x88, 0x99, 0xff}, // rgb(119, 136, 153)
	Lightslategrey:       {0x77, 0x88, 0x99, 0xff}, // rgb(119, 136, 153)
	Lightsteelblue:       {0xb0, 0xc4, 0xde, 0xff}, // rgb(176, 196, 222)
	Lightyellow:          {0xff, 0xff, 0xe0, 0xff}, // rgb(255, 255, 224)
	Lime:                 {0x00, 0xff, 0x00, 0xff}, // rgb(0, 255, 0)
	Limegreen:            {0x32, 0xcd, 0x32, 0xff}, // rgb(50, 205, 50)
	Linen:                {0xfa, 0xf0, 0xe6, 0xff}, // rgb(250, 240, 230)
	Magenta:              {0xff, 0x00, 0xff, 0xff}, // rgb(255, 0, 255)
	Maroon:               {0x80, 0x00, 0x00, 0xff}, // rgb(128, 0, 0)
	Mediumaquamarine:     {0x66, 0xcd, 0xaa, 0xff}, // rgb(102, 205, 170)
	Mediumblue:           {0x00, 0x00, 0xcd, 0xff}, // rgb(0, 0, 205)
	Mediumorchid:         {0xba, 0x55, 0xd3, 0xff}, // rgb(186, 85, 211)
	Mediumpurple:         {0x93, 0x70, 0xdb, 0xff}, // rgb(147, 112, 219)
	Mediumseagreen:       {0x3c, 0xb3, 0x71, 0xff}, // rgb(60, 179, 113)
	Mediumslateblue:      {0x7b, 0x68, 0xee, 0xff}, // rgb(123, 104, 238)
	Mediumspringgreen:    {0x00, 0xfa, 0x9a, 0xff}, // rgb(0, 250, 154)
	Mediumturquoise:      {0x48, 0xd1, 0xcc, 0xff}, // rgb(72, 209, 204)
	Mediumvioletred:      {0xc7, 0x15, 0x85, 0xff}, // rgb(199, 21, 133)
	Midnightblue:         {0x19, 0x19, 0x70, 0xff}, // rgb(25, 25, 112)
	Mintcream:            {0xf5, 0xff, 0xfa, 0xff}, // rgb(245, 255, 250)
	Mistyrose:            {0xff, 0xe4, 0xe1, 0xff}, // rgb(255, 228, 225)
	Moccasin:             {0xff, 0xe4, 0xb5, 0xff}, // rgb(255, 228, 181)
	Navajowhite:          {0xff, 0xde, 0xad, 0xff}, // rgb(255, 222, 173)
	Navy:                 {0x00, 0x00, 0x80, 0xff}, // rgb(0, 0, 128)
	Oldlace:              {0xfd, 0xf5, 0xe6, 0xff}, // rgb(253, 245, 230)
	Olive:                {0x80, 0x80, 0x00, 0xff}, // rgb(128, 128, 0)
	Olivedrab:            {0x6b, 0x8e, 0x23, 0xff}, // rgb(107, 142, 35)
	Orange:               {0xff, 0xa5, 0x00, 0xff}, // rgb(255, 165, 0)
	Orangered:            {0xff, 0x45, 0x00, 0xff}, // rgb(255, 69, 0)
	Orchid:               {0xda, 0x70, 0xd6, 0xff}, // rgb(218, 112, 214)
	Palegoldenrod:        {0xee, 0xe8, 0xaa, 0xff}, // rgb(238, 232, 170)
	Palegreen:            {0x98, 0xfb, 0x98, 0xff}, // rgb(152, 251, 152)
	Paleturquoise:        {0xaf, 0xee, 0xee, 0xff}, // rgb(175, 238, 238)
	Palevioletred:        {0xdb, 0x70, 0x93, 0xff}, // rgb(219, 112, 147)
	Papayawhip:           {0xff, 0xef, 0xd5, 0xff}, // rgb(255, 239, 213)
	Peachpuff:            {0xff, 0xda, 0xb9, 0xff}, // rgb(255, 218, 185)
	Peru:                 {0xcd, 0x85, 0x3f, 0xff}, // rgb(205, 133, 63)
	Pink:                 {0xff, 0xc0, 0xcb, 0xff}, // rgb(255, 192, 203)
	Plum:                 {0xdd, 0xa0, 0xdd, 0xff}, // rgb(221, 160, 221)
	Powderblue:           {0xb0, 0xe0, 0xe6, 0xff}, // rgb(176, 224, 230)
	Purple:               {0x80, 0x00, 0x80, 0xff}, // rgb(128, 0, 128)
	Red:                  {0xff, 0x00, 0x00, 0xff}, // rgb(255, 0, 0)
	Rosybrown:            {0xbc, 0x8f, 0x8f, 0xff}, // rgb(188, 143, 143)
	Royalblue:            {0x41, 0x69, 0xe1, 0xff}, // rgb(65, 105, 225)
	Saddlebrown:          {0x8b, 0x45, 0x13, 0xff}, // rgb(139, 69, 19)
	Salmon:               {0xfa, 0x80, 0x72, 0xff}, // rgb(250, 128, 114)
	Sandybrown:           {0xf4, 0xa4, 0x60, 0xff}, // rgb(244, 164, 96)
	Seagreen:             {0x2e, 0x8b, 0x57, 0xff}, // rgb(46, 139, 87)
	Seashell:             {0xff, 0xf5, 0xee, 0xff}, // rgb(255, 245, 238)
	Sienna:               {0xa0, 0x52, 0x2d, 0xff}, // rgb(160, 82, 45)
	Silver:               {0xc0, 0xc0, 0xc0, 0xff}, // rgb(192, 192, 192)
	Skyblue:              {0x87, 0xce, 0xeb, 0xff}, // rgb(135, 206, 235)
	Slateblue:            {0x6a, 0x5a, 0xcd, 0xff}, // rgb(106, 90, 205)
	Slategray:            {0x70, 0x80, 0x90, 0xff}, // rgb(112, 128, 144)
	Slategrey:            {0x70, 0x80, 0x90, 0xff}, // rgb(112, 128, 144)
	Snow:                 {0xff, 0xfa, 0xfa, 0xff}, // rgb(255, 250, 250)
	Springgreen:          {0x00, 0xff, 0x7f, 0xff}, // rgb(0, 255, 127)
	Steelblue:            {0x46, 0x82, 0xb4, 0xff}, // rgb(70, 130, 180)
	Tan:                  {0xd2, 0xb4, 0x8c, 0xff}, // rgb(210, 180, 140)
	Teal:                 {0x00, 0x80, 0x80, 0xff}, // rgb(0, 128, 128)
	Thistle:              {0xd8, 0xbf, 0xd8, 0xff}, // rgb(216, 191, 216)
	Transparent:          {0x00, 0x00, 0x00, 0x00}, // rgba(0,0,0,0)
	Tomato:               {0xff, 0x63, 0x47, 0xff}, // rgb(255, 99, 71)
	Turquoise:            {0x40, 0xe0, 0xd0, 0xff}, // rgb(64, 224, 208)
	Violet:               {0xee, 0x82, 0xee, 0xff}, // rgb(238, 130, 238)
	Wheat:                {0xf5, 0xde, 0xb3, 0xff}, // rgb(245, 222, 179)
	White:                {0xff, 0xff, 0xff, 0xff}, // rgb(255, 255, 255)
	Whitesmoke:           {0xf5, 0xf5, 0xf5, 0xff}, // rgb(245, 245, 245)
	Yellow:               {0xff, 0xff, 0x00, 0xff}, // rgb(255, 255, 0)
	Yellowgreen:          {0x9a, 0xcd, 0x32, 0xff}, // rgb(154, 205, 50)
}

// lookup is the lookup map.
var lookup map[uint32]NamedColor

func init() {
	lookup = make(map[uint32]NamedColor, len(colors))
	for k, v := range colors {
		lookup[mapKey(v.R, v.G, v.B, v.A)] = k
	}
}

// mapKey returns a map lookup key for r, g, b, a.
func mapKey(r, g, b, a uint8) uint32 {
	return uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
}
