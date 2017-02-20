package hello_world

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

// test stores the HTTP testing client preconfigured
var test = baloo.New(fmt.Sprintf("http://%v:3000", os.Getenv("DOCKER_HOST_IP")))

const schema = `{
  "title": "Longest Words",
	"type": "array",
	"items": {
		"type": "string",
		"pattern": "^infinitesimally$"
	}
}`

func TestLongestWord01(t *testing.T) {
	test.Post("/longest_words").
		JSON(map[string]string{"words": "The manifestation of the existential paradigm is infinitesimally larger than the exponentially evolved humanistic peon; indeed this precept is fundamentally beyond the cognisance of any finite mind"}).
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()
}
