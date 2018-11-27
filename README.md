# chinese-telegraph

## usage

```go
package main

import (
	"fmt"
	"github.com/lightsing/chinese-telegraph"
)

func main() {
	converter := chinese_telegraph.NewConverter()

	test := make([]uint, 0, 4)
	test = append(test, uint(22))
	test = append(test, uint(2429))
	test = append(test, uint(7193))
	test = append(test, uint(1032))


	fmt.Println(converter.GetString(test))
}
```

Result:
```$xslt
中文电报
```