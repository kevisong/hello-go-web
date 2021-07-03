version = 1.0.0

all: build-server build-server-image save-server-image

build-server:
	GOOS=linux GOARCH=amd64 go build -o ./cmd/server/server ./cmd/server

build-server-image:
	docker build \
		-f ./build/package/server/Dockerfile \
		. \
		-t kevisong/hello-go-web:$(version)

save-server-image:
	docker save \
	-o ./build/package/server/$(version).image \
	kevisong/hello-go-web:$(version)