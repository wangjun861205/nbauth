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
	resp, err := c.SignIn(context.Background(), &api.SignInRequest{
		Ident:    "wangjun",
		Password: "Wt20110523",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
