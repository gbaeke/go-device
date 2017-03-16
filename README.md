# go-device
Device API with Go Micro

Dockerfile uses the empty scratch image and requires that you build a static exe with the following command:

```bash

```CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
