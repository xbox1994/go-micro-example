package api

import (
	"GoMicroExample/api/user/proto"
	"context"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if req.Service() == "go.micro.api.user" && req.Method() == "User.Login" {
			err := fn(ctx, req, resp)
			return err
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]
		authClient := user.NewUserService("go.micro.api.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &user.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
