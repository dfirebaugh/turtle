PROJECTNAME=turtle

all: native windows-client-cross-compile wasm examples web

run:
	go run ./cmd/emulator

run-gl:
	go run -tags gl ./cmd/emulator

test:
	go test ./...

native:
	go build -o .dist/$(PROJECTNAME) ./cmd/emulator

wasm:
	GOOS=js GOARCH=wasm go build -o .dist/$(PROJECTNAME).wasm ./cmd/webemulator

.PHONY: web
web:
	mkdir -p .dist \
	&& cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .dist/ \
	&& cp -R web/* .dist/

web-server: ## for testing - note: requires node
	npx es-dev-server --root-dir .dist

clean:
	rm -rf .dist/
