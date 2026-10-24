package main

// Canary only: reaches x/net/http2 at a vulnerable version so govulncheck fails.
import (
	"net"
	"os"

	"golang.org/x/net/http2"
)

func init() {
	if os.Getenv("GOWSAY_CANARY") == "serve" {
		c, _ := net.Pipe()
		(&http2.Server{}).ServeConn(c, nil)
	}
}
