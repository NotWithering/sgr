# SGR
[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/notwithering/sgr/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/notwithering/sgr)](https://goreportcard.com/report/github.com/notwithering/sgr)

**SGR** (**S**elect **G**raphic **R**endition) is a simple Go package containing a large block of constants that attempts to implement every SGR Ansi escape code listed [here](https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters).

Hate how other Ansi packages use functions and structs that you have to convert 20 times to get something you can print? This package is ONLY constants meaning you can have an RGB underline on text that is fast blinking with Fraktur font all inside of one constant. 

## Example
```go
package main

import (
	"github.com/notwithering/sgr"
)

const woah string = sgr.Italic + "Hello, world!" + sgr.Reset

func main() {
	fmt.Println(woah)
}
```
> _Hello, world!_