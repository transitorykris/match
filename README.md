# v6match

Finds all the IPv6 addresses in a string.

## Usage

```go
package main

import (
	"fmt"

	"github.com/transitorykris/v6match"
)

func main() {
	fmt.Println(v6match.MatchIPv6("This 1:2:3::4 is an IPv6 address")) // [1:2:3::4]
	fmt.Println(v6match.MatchIPv6("1:2:3::4, 5:6:7::/32, ::/0"))       // [1:2:3::4 5:6:7::/32 ::/0]
	fmt.Println(v6match.MatchIPv6("192.168.1.1 is not a v6 address"))  // []
}
```
