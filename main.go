package main

import (
	"github.com/thorstenkloehn/thorsten/server"
	"net/http"
)

func main() {

	var server server.Start

	http.ListenAndServe(":80", server.Start())
}
