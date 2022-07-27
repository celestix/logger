# Logger
Logger is a package designed for logging in modular programs specifically, It contains many cool features like colorised logging, different logging levels, etc.

You can use this package to implement logging in your Go program, for any futher help you can check out the [documentations](https://pkg.go.dev/github.com/anonyindian/logger).

[![Go Reference](https://pkg.go.dev/badge/github.com/anonyindian/logger.svg)](https://pkg.go.dev/github.com/anonyindian/logger) [![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

**Note**: This library is in the alpha stage yet and doesn't cover every aspect of logging.

## Installation
You can download the library with the help of standard `go get` command.

```bash
go get github.com/anonyindian/logger
```

## Usage
```go
package main
import (
	logger "github.com/anonyindian/logger"
    "os"
)
func main() {
    log := logger.New(os.Stderr, nil)
    log.Println("Hello World")
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html)
<br>Licensed Under <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">GNU General Public License v3</a>
