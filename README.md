# singidate

## install

> goto `go.mod` add

```go
require github.com/singi2016cn/singidate v0.2.0
```

## use

```go
package main

import (
	"fmt"

	"github.com/singi2016cn/singidate"
)

func main() {
    fmt.Println(singidate.Timestamp2DateTime(1604560223)) //2020-11-05 15:10:23
    fmt.Println(singidate.DateTime2TimestampWithShanghai("2020-11-05 15:10:23")) //1604560223
}
```