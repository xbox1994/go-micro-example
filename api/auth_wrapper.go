package api

import (
	"context"
	"errors"
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
		claims, e := Decode(token)
		log.Println("Auth Resp:", claims.User)
		if e != nil {
			return e
		}
		e = fn(ctx, req, resp)
		return e
	}
}
