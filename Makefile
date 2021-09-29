version = 1.0.0

all: build-server build-web build-server-image save-server-image

.PHONY: build-server
build-server:
	GOOS=linux GOARCH=amd64 go build -o ./cmd/server/server ./cmd/server

.PHONY: build-web
build-web:
	statik -src=./web

.PHONY: build-server-image
build-server-image:
	docker build \
		-f ./build/package/server/Dockerfile \
		. \
		-t kevisong/hello-go-web:$(version)

.PHONY: save-server-image
save-server-image:
	docker save \
	-o ./build/package/server/$(version).image \
	kevisong/hello-go-web:$(version)

.PHONY: push
push:
	docker push \
	kevisong/hello-go-web:$(version)
