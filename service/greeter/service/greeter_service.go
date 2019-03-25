package service

import (
	"GoMicroExample/service/config"
	"GoMicroExample/service/constant/code"
	"GoMicroExample/service/greeter/dto"
	"GoMicroExample/service/user/proto"
	"context"
	"errors"
)

type GreeterService struct {
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (this *GreeterService) Greeter(ctx context.Context, userClient user.UserService, req *dto.HelloRequest) (*dto.HelloResponse, int32, error) {
	if req == nil || req.Name == "" {
		return nil, code.InvalidParam, errors.New("param invalid")
	}
	info, e := userClient.GetUserInfo(ctx, &user.Empty{})
	if e != nil {
		return nil, code.InternalServerCallError, e
	}
	return &dto.HelloResponse{
		SettingMessage: config.LocalConfig.Greetings.String,
		Id:             info.Id,
		Username:       info.Username,
		Password:       info.Password,
	}, code.OK, nil
}
