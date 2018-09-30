package main

import (
	"GoMicroExample/api/auth"
	"GoMicroExample/service/user/proto"
	userApi "GoMicroExample/service/user/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"log"
)

type UserService struct {
}

func (us *UserService) GetUserInfo(ctx context.Context, req *userApi.Empty, rsp *userApi.UserInfo) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("go.micro.api.user", "no auth meta-data found in request")
	}
	rsp.Id = meta["X-Example-Id"]
	rsp.Username = meta["X-Example-Username"]
	rsp.Password = "password from db"
	return nil
}

func (us *UserService) Login(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.user", "require post")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.user", "expect application/json")
	}

	var userInfo user.UserInfo
	json.Unmarshal([]byte(req.Body), &userInfo)

	token, e := auth.Encode(&userInfo)
	if e != nil {
		return e
	}
	b, _ := json.Marshal(map[string]string{
		"token": token,
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	userService := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	userService.Init()

	userApi.RegisterUserHandler(userService.Server(), &UserService{})

	if err := userService.Run(); err != nil {
		log.Fatal(err)
	}
}
