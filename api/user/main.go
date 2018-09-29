package main

import (
	"GoMicroExample/api/user/proto"
	userApi "GoMicroExample/api/user/proto"
	tokenService "GoMicroExample/api/user/service"
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"log"
)

type UserService struct {
	TokenService tokenService.Authable
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

	var user user.UserInfo
	json.Unmarshal([]byte(req.Body), &user)

	token, e := us.TokenService.Encode(&user)
	if e != nil {
		return e
	}
	b, _ := json.Marshal(map[string]string{
		"token": token,
	})
	rsp.Body = string(b)
	return nil
}

func (us *UserService) ValidateToken(ctx context.Context, req *userApi.Token, rsp *userApi.UserInfo) error {
	if req.Token == "" {
		return errors.InternalServerError("go.micro.api.user", "empty token")
	}

	decode, e := us.TokenService.Decode(req.Token)
	if e != nil {
		return e
	}

	if decode.User.Id == "" {
		return errors.InternalServerError("go.micro.api.user", "invalid user")
	}

	rsp.Username = decode.User.Username
	rsp.Password = decode.User.Password
	rsp.Id = decode.User.Id
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		//micro.WrapHandler(api.AuthWrapper),
	)

	service.Init()

	userApi.RegisterUserHandler(service.Server(), &UserService{
		TokenService: &tokenService.TokenService{},
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
