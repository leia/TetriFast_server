package main

import (
	"github.com/leia/TetriFast_server/blocks"
	"github.com/leia/TetriFast_server/network"
)

func main() {
	arr := []string{"A2", "H5", "D8", "C22", "A22", "C4", "E12"}
	blocks.GenerateSpecials(arr)
	network.StartServer("localhost:8080")
}
