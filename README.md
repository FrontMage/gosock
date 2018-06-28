# gosock
Golang unix sock lib

## Usage
`$ go get github.com/FrontMage/gosock`

```go
import "github.com/FrontMage/gosock"
import "net"

gosock.Listen("/tmp/go.sock", func(conn net.Conn){/* Handle incoming msg */})
```