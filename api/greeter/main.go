package main

import (
	"GoMicroExample/api"
	greeterApi "GoMicroExample/api/greeter/proto"
	"GoMicroExample/api/user/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"log"
)

type Greeter struct {
	userClient user.UserService
}

func (ga *Greeter) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Print("Received Greeter.Hello API request")

	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("go.micro.api.greeter", "no auth meta-data found in request")
	}
	token := meta["Token"]

	decode, e := api.Decode(token)
	if e != nil {
		return e
	}
	if decode.User.Id == "" {
		return errors.InternalServerError("go.micro.api.user", "invalid user")
	}
	log.Println("Token decoded info:", decode.User)

	rsp.StatusCode = 200
	b, _ := json.Marshal(
		map[string]string{"message": "nice to meet u, " + decode.User.Username + ", your password: " + decode.User.Password + ", your id: " + decode.User.Id})

	rsp.Body = string(b)
	return nil
}

var (
	config map[string]interface{}
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.WrapHandler(api.AuthWrapper),
		micro.Flags(
			cli.StringFlag{
				Name: "profile",
			}, cli.StringFlag{
				Name: "config_server",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			profile := c.String("profile")
			if len(profile) > 0 {
				fmt.Println("profile set to", profile)
			}
			configServer := c.String("config_server")
			if len(configServer) > 0 {
				fmt.Println("config_server set to", configServer)
			}
			//config = conf.GetConfig(configServer, "greeter", profile)
		}))

	greeterApi.RegisterGreeterHandler(service.Server(), &Greeter{
		userClient: user.NewUserService("go.micro.api.user", service.Client())})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
