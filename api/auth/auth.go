package auth

import (
	"GoMicroExample/service/constant/code"
	"GoMicroExample/service/util"
	"encoding/json"
	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
	"log"
	"net/http"
)

type Auth struct {
}

func (*Auth) Flags() []cli.Flag {
	return nil
}

func (*Auth) Commands() []cli.Command {
	return nil
}

func (*Auth) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth plugin received: " + r.URL.Path)
			if r.URL.Path == "/user/login" {
				h.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			userFromToken, e := Decode(token)

			if e != nil {
				response, _ := json.Marshal(util.CommonResponse{
					Code:    code.AuthorizationError,
					Message: "please login",
				})
				w.Write(response)
				return
			}

			r.Header.Set("X-Example-Id", userFromToken.Id)
			r.Header.Set("X-Example-Username", userFromToken.Username)
			h.ServeHTTP(w, r)
			return
		})
	}
}

func (*Auth) Init(*cli.Context) error {
	return nil
}

func (*Auth) String() string {
	return ""
}
