package main

import (
	"github.com/go-martini/martini"
)

const (
	SERVER_ADDR string = ":6603"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "server is running ..."
	})

	m.RunOnAddr(SERVER_ADDR)
}
