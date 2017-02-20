## Basic Auth

Build a simple JSON API that uses basic auth to access protected `/me` resource


#### Sample Input
```
Header('Authorization', 'Basic admin:admin')
```

#### Sample Output
```
{"user": "admin"}
```

#### Run
```
docker-compose run --rm api go test specs/simple_auth/simple_auth_test.go
```
