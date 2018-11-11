package tests

import (
	"testing"

	"github.com/leia/TetriFast_server/network"
)

func TestConnection(t *testing.T) {
	network.StartServer("localhost:8081")
}
