# ctrlzr/urldefault-go

Golang wrapper for net/url with the option to specify a default URL and override with supplied parts.

* [Install](#install)
* [Examples](#examples)
* [Testing](#testing)

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u -v github.com/koshatul/urldefault-go
```

## Examples

Works as a dropin replacement for url.Parse() in most situations.

```go
package main 

import (
    "fmt"
    "net/url"
)

func main() {
    u, _ := url.Parse("http://localhost/path")
    fmt.Println(u.String())
}
```

becomes 

```go
package main 

import (
    "fmt"
    "github.com/koshatul/urldefault-go/src/urldefault"
)

func main() {
    u, _ := urldefault.Parse("http://localhost/path")
    fmt.Println(u.String())
}
```

But more functionality is available, by specifying a second parameter, any missing parts of the first parameter are filled in.

```go
package main 

import (
    "fmt"
    "github.com/koshatul/urldefault-go/src/urldefault"
)

func main() {
	u, _ := urldefault.Parse("amqp://differenthost/", "amqp://admin:admin@localhost:5672/vhost")
    fmt.Println(u.String())
}
```
would return `amqp://admin:admin@differenthost:5672/vhost`

## Testing

```sh
go test github.com/koshatul/urldefault-go/...
```
