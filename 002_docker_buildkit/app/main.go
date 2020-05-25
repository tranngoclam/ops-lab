package main

import (
	"fmt"
	"github.com/tranngoclam/docker-tips/002_docker_buildkit"
)

// Flavor defines default flavor to brew coffee
var Flavor = "."

func main() {
	_02_docker_buildkit.LoadCoffee()

	fmt.Println(Flavor)
}
