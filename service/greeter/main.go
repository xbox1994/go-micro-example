package main

import (
	"GoMicroExample/hystrix"
	"GoMicroExample/service"
	greeterApi "GoMicroExample/service/greeter/proto"
	"GoMicroExample/service/user/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"log"
)

type Greeter struct {
	userClient user.UserService
}

func (ga *Greeter) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Print("Received Greeter.Hello API request")

	meta, _ := metadata.FromContext(ctx)
	log.Println(meta)
	info, e := ga.userClient.GetUserInfo(ctx, &user.Empty{})

	if e != nil {
		return e
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(
		map[string]string{"message": "nice to meet u, I know your username,password,id in token, and email in db.",
			"username": info.Username, "password": info.Password, "id": info.Id, "email": info.Email})

	rsp.Body = string(b)
	return nil
}

var (
//config map[string]interface{}
)

func main() {
	hystrix.Configure([]string{"go.micro.api.user.User.GetUserInfo"})
	greeterService := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.WrapHandler(service.AuthWrapper),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.Flags(
			cli.StringFlag{
				Name: "profile",
			}, cli.StringFlag{
				Name: "config_server",
			},
		),
	)

	greeterService.Init(
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

	greeterApi.RegisterGreeterHandler(greeterService.Server(), &Greeter{
		userClient: user.NewUserService("go.micro.api.user", greeterService.Client())})

	if err := greeterService.Run(); err != nil {
		log.Fatal(err)
	}
}
