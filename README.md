# Hello modules
This is a test repo for Golang Modules

## Installation
Execute
```bash
go get -u github.com/MarcoManjarrez/hello_modules
```

## Usage 

package main

```go
import (
	"fmt"

	"github.com/MarcoManjarrez/hello_modules"
)

func main() {
	fmt.Printf(hello_modules.RandomHello(), "Bobo")
}
```
