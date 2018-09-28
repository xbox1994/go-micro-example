package main

import (
	"GoMicroExample/api"
	greeterApi "GoMicroExample/api/greeter/proto"
	greeterService "GoMicroExample/service/greeter/proto"
	"context"
	"encoding/json"
	"errors"
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
		"message": "got your name " + strings.Join(name.Values, " "),
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.WrapHandler(api.AuthWrapper),
	)

	service.Init()

	greeterApi.RegisterGreeterHandler(service.Server(), &Greeter{
		greeterServiceClient: greeterService.NewGreeterService("go.micro.srv.greeter", client.DefaultClient),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
