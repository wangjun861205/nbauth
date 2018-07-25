package bkauth

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"./api"
	"./client"
	"./server"
)

func TestService(t *testing.T) {
	go func() {
		log.Fatal(server.Serve("127.0.0.1:9999", "./", "wangjun:Wt20110523@tcp(localhost:12345)/bk_dalian"))
	}()
	time.Sleep(3 * time.Second)
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
