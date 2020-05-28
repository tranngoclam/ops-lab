package main

import (
	"fmt"
	"github.com/tranngoclam/ops-lab/buildkit/secret"
)

var Secret = "."

func main() {
	fmt.Println(Secret)
	fmt.Println(secret.Secret)
}
