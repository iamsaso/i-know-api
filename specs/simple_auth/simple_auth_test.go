package basic_auth

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

// test stores the HTTP testing client preconfigured
var test = baloo.New(fmt.Sprintf("http://%v:3000/me", os.Getenv("DOCKER_HOST_IP")))

const test_schema = `{
  "title": "Me Schema",
  "type": "object",
  "properties": {
    "user": {
      "type": "string"
    }
  },
  "required": ["user"]
}`

func TestProtected(t *testing.T) {
	test.Get("/me").
		SetHeader("Authorization", "Basic YWRtaW46YWRtaW4=").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(test_schema).
		Done()
}
