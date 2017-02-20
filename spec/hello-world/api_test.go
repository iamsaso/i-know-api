package hello_world

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

// test stores the HTTP testing client preconfigured
var test = baloo.New(fmt.Sprintf("http://%v:3000", os.Getenv("DOCKER_HOST_IP")))

func TestPing(t *testing.T) {
	test.Get("/ping").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(`{"status": "pong"}`).
		Done()
}
