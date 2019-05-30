PROJECT?=github.com/grubastik/gopherconEU2019
VERSION?=0.0.1
COMMIT=$(shell git log -1 --pretty=format:"%H")
BUILDTIME=$(shell date --utc +%FT%TZ)

test:
	go test --race ./...

build:
	GO111MODULE=on CGO_ENABLED=0 go build \
	-ldflags " -s -w -X ${PROJECT}/internal/diagnostics.Version=${VERSION} -X ${PROJECT}/internal/diagnostics.Commit=${COMMIT} -X ${PROJECT}/internal/diagnostics.BuildTime=${BUILDTIME}" \
	-o bin/workshop \
	${PROJECT}/cmd/workshop

docker-image:
	docker build -t workshop .