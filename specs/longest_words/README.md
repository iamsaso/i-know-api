## Hello World

Build a JSON API with `/words` endpoint that accepts words and returns a list with the longest words of same length.


#### Sample Input
```
{"words": "The manifestation of the existential paradigm is infinitesimally larger than the exponentially evolved humanistic peon; indeed this precept is fundamentally beyond the cognisance of any finite mind"}
```

#### Sample Output
```
["infinitesimally"]
```

#### Run
```
docker-compose run --rm api go test specs/longest_words/longest_words_test.go
```
