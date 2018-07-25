package bkauth

import (
	"context"
	"fmt"
	"testing"

	"github.com/wangjun861205/nbauth/api"
	"github.com/wangjun861205/nbauth/client"
)

func TestService(t *testing.T) {
	c, err := client.NewClient("127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	resp, err := c.SignUp(context.Background(), &api.SignUpRequest{Username: "admin", Password: "Wt20110523", Phone: "13793148690", Email: "444055828@qq.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
