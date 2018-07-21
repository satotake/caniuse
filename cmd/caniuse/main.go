package main

import (
	"fmt"
	"github.com/satotake/go-caniuse/pkg/caniuse"
)

func main() {
	data := caniuse.GetAll()
	fmt.Println(data)
}
