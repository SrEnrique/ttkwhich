# ttkwhich

## Usage 

### Install package 

``` bash
go get github.com/srenrique/ttkwhich
```


### Importa package
``` go
import (
	"fmt"
	"github.com/srenrique/ttkwhich"
)
```

### Linux usage

#### Methods

1. Platform

Get current operative system platform


``` go
	platform := ttkwhich.Platform()
	fmt.Println(mesplatformsage)
```

2. Which

ttkwhich.Which has which method that searches set of directories for an executable file based on the PATH environment variable.

``` go
	fmt.Println(ttkwhich.Which("cat"))
```

3. Which_in_path

ttkwhich.Which_in_path You can also specify directly the paths.

``` go
	fmt.Println(ttkwhich.Which_in_path("ruby", []string{"/home/luis/.rbenv/shims"}))
```


## Example

``` go
package main

import (
	"fmt"

	"github.com/srenrique/ttkwhich"
)

func main() {
	// Get a greeting message and print it.
	platform := ttkwhich.Platform()
	fmt.Println("Platform:", platform)

	fmt.Println("cat path:", ttkwhich.Which("cat"))

	fmt.Println("Search ruby in scpecify path:", ttkwhich.Which_in_path("ruby", []string{"/home/luis/.rbenv/shims"}))
}
```