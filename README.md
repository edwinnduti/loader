# Loader

- These are loaders which can be adjusted speed as in accordance the user to show a user that a specific process in loading.

### Installation

```bash
go get github.com/edwinnduti/loader
```

- The types of loaders/spinners include:

| label           | name | 		figures        	|
| --------------- | ---- | ---------------------------- |
| rotating circes | dc   | "⠈⠁", "⠈⠑", "⠈⠱", "⠈⡱", "⢀⡱", "⢄⡱", "⢄⡱", "⢆⡱", "⢎⡱", "⢎⡰", "⢎⡠", "⢎⡀", "⢎⠁", "⠎⠁", "⠊⠁" |
| rotating blocks | rb   | "⣾","⣽","⣻","⢿","⡿", "⣟","⣯","⣷" |
| rotating sticks | rc   | 	"\\", "|", "/", "-" 	|

### Usage

```go
package main

import (
    "fmt"
    "os"
    "time"
    "github.com/edwinnduti/loader"
)

func main(){
	loader := New(os.Stdout, "dc", 100, "Waiting", "")

	loader.Initialize()
	time.Sleep(time.Second * 5)
	loader.End()
}

```

Happy coding and anime watching!
