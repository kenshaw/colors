# colors

`colors` is a Go package to parse color strings. Provides a type to use easily
[with `cobra` / `pflag`][cobra] or other command line packages.

[Using][] | [Example][] | [About][]

[Using]: #using "Using"
[Example]: #example "Example"
[About]: #about "About"

[![Unit Tests][colors-ci-status]][colors-ci]
[![Go Reference][goref-colors-status]][goref-colors]
[![Discord Discussion][discord-status]][discord]

[colors-ci]: https://github.com/kenshaw/colors/actions/workflows/test.yml
[colors-ci-status]: https://github.com/kenshaw/colors/actions/workflows/test.yml/badge.svg
[goref-colors]: https://pkg.go.dev/github.com/kenshaw/colors
[goref-colors-status]: https://pkg.go.dev/badge/github.com/kenshaw/colors.svg
[discord]: https://discord.gg/yJKEzc7prt "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"

## Using

Install in the usual Go fashion:

```sh
$ go get -u github.com/kenshaw/colors@latest
```

## Example

```go
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
		"hex(0,0,0,0)",
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
	// transparent
}
```

## About

Built to support the [`fv`][fv], [`iv`][iv], and [`usql`][usql] applications.

[cobra]: https://github.com/spf13/cobra
[iv]: https://github.com/kenshaw/iv
[fv]: https://github.com/kenshaw/fv
[usql]: https://github.com/xo/usql
