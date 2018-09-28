package conf

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	log.Printf("%v", GetConfig("http://127.0.0.1:8081", "greeter", "test"))
}
