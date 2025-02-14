.PHONY: build install clean test integration dep x-platform release docker
VERSION=`egrep -o '[0-9]+\.[0-9a-z.\-]+' version.go`
GIT_SHA=`git rev-parse --short HEAD || echo`

build:
	@echo "Building go-jams..."
	@mkdir -p bin
	@go build -ldflags "-X main.GitSHA=${GIT_SHA}" -o bin/server .

install:
	@echo "Installing go-jams..."
	@install -c bin/server /usr/local/bin/server

clean:
	@rm -f bin/*

dep:
	@dep ensure

x-platform:
	@docker build -q -t jams_builder -f docker/Dockerfile.build.alpine .
	@for platform in darwin linux windows; do \
		if [ $$platform == windows ]; then extension=.exe; fi; \
		docker run -it --rm -v ${PWD}:/app -e "GOOS=$$platform" -e "GOARCH=amd64" -e "CGO_ENABLED=0" jams_builder go build -ldflags="-s -w -X main.GitSHA=${GIT_SHA}" -o bin/server-${VERSION}-$$platform-amd64$$extension; \
	done
	@docker run -it --rm -v ${PWD}:/app -e "GOOS=linux" -e "GOARCH=arm64" -e "CGO_ENABLED=0" jams_builder go build -ldflags="-s -w -X main.GitSHA=${GIT_SHA}" -o bin/server-${VERSION}-linux-arm64;
	@upx bin/server-${VERSION}-*

release:
	@docker build -q -t fips_builder -f docker/Dockerfile.build.mariner .
	@docker run -it --rm -v ${PWD}:/app -e "GOOS=linux" -e "GOARCH=amd64" -e "GOEXPERIMENT=opensslcrypto" -e "CGO_ENABLED=1" fips_builder go build -o bin/server-fips-${VERSION}-linux-amd64;

docker:
	@docker build --build-arg VERSION=${VERSION} -f docker/Dockerfile -t go-jams:${VERSION} .
	@docker build --build-arg VERSION=${VERSION} -f docker/Dockerfile.fips -t go-jams:fips-${VERSION} .
	@docker build --build-arg VERSION=${VERSION} -f docker/Dockerfile.datadog -t go-jams:proxy-${VERSION} .