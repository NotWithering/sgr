# SGR [![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/NotWithering/sgr/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/NotWithering/sgr)](https://goreportcard.com/report/github.com/NotWithering/sgr)

**SGR** (**S**elect **G**raphic **R**endition) is a simple Go package constaining a large block of constants that attempts to implement every SGR Ansi escape code listed [here](https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters).

I was annoyed how I couldn't use a traditional ansi package in constants due to their complicated approach, so I just decided to make one big constant.

## Example
```go
package main

import (
	"github.com/NotWithering/sgr"
)

const woah string = sgr.Italic + "Hello, world!" + sgr.Reset

func main() {
	fmt.Println(woah)
}
```
> _Hello, world!_