# 🪁 Go Client for [Instatus](https://instatus.com)
> *Lightweight and speedy Go client for Instatus*

## Why did you build this?
This is mainly for the [CLI](https://github.com/auguwu/instatus-cli) I am making and for my Kubernetes watcher, [Kanata](https://github.com/auguwu/Kanata)
but anyone can use this!

## Usage
To use the library, you can simply just `go get` it!

```shell
$ go get github.com/gxben/instatus-go
```

and now, you can start using it!

```go
package main

import (
	"fmt"
	"github.com/gxben/instatus-go"
)

func main() {
	client := instatus.NewClient(
		instatus.WithToken("<token>"),
		instatus.WithUserAgent("some/user-agent v0.0.0"),
	)
	
	res, err := client.Statuspages.List()
	if err != nil {
		panic(err)
    }
	
	fmt.Println(res)
}
```

## License
**instatus-go** is released under the [MIT License](/LICENSE) by Noel. Read the **LICENSE** file in the repository
for more information.
