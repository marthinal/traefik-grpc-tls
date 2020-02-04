__TRAEFIK - GRPC with TLS EXAMPLE__

Traefik as reverse proxy for gRPC application with self-signed certificates.

Steps:

```
./scripts/traefik.sh 

cd client

go run main.go
```

This is the result:

````
marthinal$ go run main.go 
2020/02/04 17:43:30 Greeting: Hello world
````