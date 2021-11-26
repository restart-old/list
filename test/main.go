package main

import (
	"fmt"

	"github.com/RestartFU/whitelist"
	"github.com/df-mc/dragonfly/server"
)

func main() {
	wl, _ := whitelist.New("./whitelist.json")
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	s.CloseOnProgramEnd()
	for {
		if p, err := s.Accept(); err != nil {
			return
		} else {
			if !wl.Whitelisted(p.Name()) {
				fmt.Println("not whitelisted")
				err := wl.Add(p.Name())
				fmt.Println(err)
			}
		}
	}
}
