package main

import (
	"GoMicroExample/api/auth"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
)

func init() {
	plugin.Register(&auth.Auth{})
}

func main() {
	cmd.Init()
}
