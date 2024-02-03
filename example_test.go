package colors_test

import (
	"fmt"
	"log"

	"github.com/kenshaw/colors"
)

func Example() {
	for _, s := range []string{
		"red",
		"rgba(192, 192, 192, 255)",
		"navajo white",
		"#ffeeaa",
		"rgb(106,90,205)",
		"rgb(0,255,0)",
		"hex(f,e,a)",
		"#fff",
		"rgba(26,33,80,22)",
		"#cdcdcd80",
		"#cdcdcdff",
	} {
		c, err := colors.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
	}
	// Output:
	// red
	// silver
	// navajowhite
	// #fea
	// slateblue
	// lime
	// #0f0e0a
	// white
	// #1a215016
	// #cdcdcd80
	// #cdcdcd
}
