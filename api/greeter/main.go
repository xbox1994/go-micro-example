package main

import (
	"GoMicroExample/api"
	greeterApi "GoMicroExample/api/greeter/proto"
	"GoMicroExample/config"
	greeterService "GoMicroExample/service/greeter/proto"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"log"
	"strings"
)

type Greeter struct {
	greeterServiceClient greeterService.GreeterService
}

func (ga *Greeter) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.New("no name")
	}

	ga.greeterServiceClient.Hello(context.TODO(), &greeterService.HelloRequest{
		Name: name.Values[0],
	})

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": config["string"].(string) + "and we got your name " + strings.Join(name.Values, " "),
	})
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
			config = conf.GetConfig(configServer, "greeter", profile)
		}))

	greeterApi.RegisterGreeterHandler(service.Server(), &Greeter{
		greeterServiceClient: greeterService.NewGreeterService("go.micro.srv.greeter", client.DefaultClient),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
