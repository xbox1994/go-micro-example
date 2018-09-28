package main

import (
	"GoMicroExample/api"
	userApi "GoMicroExample/api/user/proto"
	userService "GoMicroExample/service/user/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"log"
)

type User struct {
	userServiceClient userService.UserService
}

func (ua *User) Login(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	authResp, err := ua.userServiceClient.Login(context.Background(), &userService.User{
		Id:       "id",
		Username: "username",
		Password: "username",
	})

	if err != nil {
		return err
	}

	log.Println("Login Resp:", authResp)
	b, _ := json.Marshal(map[string]string{
		"token": authResp.Token,
	})
	rsp.Body = string(b)
	return err
}

func AuthWithoutLoginWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if req.Service() == "go.micro.api.user" && req.Method() == "User.Login" {
			err := fn(ctx, req, resp)
			return err
		}
		return api.AuthWrapper(fn)(ctx, req, resp)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		micro.WrapHandler(AuthWithoutLoginWrapper),
	)

	service.Init()

	userApi.RegisterUserHandler(service.Server(), &User{
		userServiceClient: userService.NewUserService("go.micro.srv.user", client.DefaultClient),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
