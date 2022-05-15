# hello-world-service

very basic golang api service

## pushing to gcr

`docker buildx build --platform linux/amd64,linux/arm64 --tag us.gcr.io/telvanni-platform/hello-world-service:$TAG 
--push .`

## API

GET `/metrics`

prometheus end point

GET `/hello`

prints `{"message":"hello world"}`

POST `/hello`

with request body

```json
{"message": "foo"}
```

the endpoint will print the value of the `message` field