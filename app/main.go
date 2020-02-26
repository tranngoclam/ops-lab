package main

import (
	"flag"
	"fmt"
	"github.com/tranngoclam/go-coffee/app/build"
)

// Flavor defines default flavor to brew coffee
var Flavor = "."

var KeyOne = "ko"

func main() {
	secret := flag.String("secret", "", "this is secret")
	flag.Parse()

	if *secret == build.KeyOne {
		fmt.Println("congratz kk")
	}
}
