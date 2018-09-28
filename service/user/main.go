package main

import (
	"GoMicroExample/service/user/proto"
	tokenService "GoMicroExample/service/user/service"
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
)

type User struct {
	TokenService tokenService.Authable
}

func (u *User) Login(ctx context.Context, req *user.User, rsp *user.Token) error {
	if req.Username != req.Password {
		return errors.New("wrong username or password")
	}

	token, e := u.TokenService.Encode(req)
	if e != nil {
		return e
	}

	rsp.Token = token
	return nil
}

func (u *User) ValidateToken(ctx context.Context, req *user.Token, rsp *user.Token) error {
	if req.Token == "" {
		return errors.New("empty token")
	}

	decode, e := u.TokenService.Decode(req.Token)
	if e != nil {
		return e
	}

	if decode.User.Id == "" {
		rsp.Error = &user.Error{
			Code:    -1,
			Message: "xx",
		}
		rsp.Valid = false
		return errors.New("invalid user")
	}

	rsp.Valid = true
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
	)

	service.Init()

	user.RegisterUserServiceHandler(service.Server(), &User{
		TokenService: &tokenService.TokenService{},
	})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
