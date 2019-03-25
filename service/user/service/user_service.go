package service

import (
	"GoMicroExample/api/auth"
	"GoMicroExample/service/constant/code"
	"GoMicroExample/service/user/dto"
	"GoMicroExample/service/user/proto"
	"errors"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) Login(req *user.UserInfo) (*dto.LoginResponse, int32, error) {
	if req == nil || req.Id == "" || req.Username == "" || req.Password == "" {
		return nil, code.InvalidParam, errors.New("param invalid")
	}

	token, e := auth.Encode(req)
	if e != nil {
		return nil, code.JwtEncodeError, e
	}

	return &dto.LoginResponse{
		Token: token,
	}, code.OK, nil
}
