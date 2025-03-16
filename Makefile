# Detect the operating system and set TAILWINDCSS_OS_ARCH accordingly
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	TAILWINDCSS_OS_ARCH := linux-x64
else ifeq ($(UNAME_S),Darwin)
	TAILWINDCSS_OS_ARCH := macos-arm64
else
	$(error This Makefile only supports Linux and macOS. Windows is not supported.)
endif

.PHONY: benchmark
benchmark:
	go test -bench=.

.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o public/styles/app.css --minify

.PHONY: build-docker
build-docker: build-css
	docker build --platform linux/amd64,linux/arm64 .

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start: build-css
	go run ./cmd/app

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(TAILWINDCSS_OS_ARCH)
	mv tailwindcss-$(TAILWINDCSS_OS_ARCH) tailwindcss
	chmod a+x tailwindcss
	mkdir -p node_modules/tailwindcss/lib && ln -sf tailwindcss node_modules/tailwindcss/lib/cli.js
	echo '{"devDependencies": {"tailwindcss": "latest"}}' >package.json

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

.PHONY: watch-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o public/styles/app.css --watch
