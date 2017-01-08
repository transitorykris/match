package main

import (
	"fmt"

	"github.com/transitorykris/v6match"
)

func main() {
	fmt.Println(v6match.MatchIP("This string has 1.2 4:5 no IPs in it")) // []
	fmt.Println(v6match.MatchIP("This 1:2:3::4 is an IPv6 address"))     // [1:2:3::4]
	fmt.Println(v6match.MatchIP("1:2:3::4, 5:6:7::/32, ::/0"))           // [1:2:3::4 5:6:7::/32 ::/0]
	fmt.Println(v6match.MatchIP("192.168.1.1 is not a v6 address"))      // [192.168.1.1]
	fmt.Println(v6match.MatchIP("192.168.1.0/24 is a fine CIDR"))        // [192.168.1.0/24]
}
