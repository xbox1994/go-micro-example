package util

import (
	"common/log"
	"encoding/json"
	"github.com/micro/go-api/proto"
)

type CommonResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func Resp(code int32, err error, rsp *go_api.Response, response interface{}) error {
	if err != nil {
		log.Error(nil, err)
		changeBody(code, err.Error(), rsp, response)
	} else {
		changeBody(code, "ok", rsp, response)
	}
	return nil
}

func changeBody(code int32, msg string, rsp *go_api.Response, response interface{}) {
	var bodyMap map[string]interface{}

	other, _ := json.Marshal(response)
	json.Unmarshal(other, &bodyMap)

	common, _ := json.Marshal(CommonResponse{
		Code:    code,
		Message: msg,
	})
	json.Unmarshal(common, &bodyMap)

	bytes, _ := json.Marshal(bodyMap)
	rsp.Body = string(bytes)
}
