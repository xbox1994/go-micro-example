package main

import (
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"go-micro-example/api/auth"
)

func init() {
	plugin.Register(&auth.Auth{})
}

func main() {
	cmd.Init()
}
