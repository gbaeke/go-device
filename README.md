# go-device
Device API with Go Micro

To run as server: go run main.go
To run as client: go run main.go --run_client

Without any parametes, server and client expect Consul to registery or query the service. If you want to run the service and client without any external registries like Consul, use multicast DNS instead like so:

Server: go run main.go --registry mdns
Client: go run main.go --registry mdns

The Dockerfile uses the empty scratch image and requires that you build a static exe with the following command:

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`

To run in a container, you can use environment variables to specify options like the registry to use like so:

`docker run --env MICRO_REGISTRY=mdns image_tag`
