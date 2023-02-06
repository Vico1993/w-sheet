default: build lint

.PHONY: build test lint

build:
	@ echo "🛠  Start building 🛠"
	@ go build -a \
			 -o "./w-sheet"
	@ echo "🛠  Build done 🛠"

test:
	go test -v ./...

lint:
	@ echo "🪛  Start linting 🪛"
	@ golangci-lint run ./... -v
	@ echo "🪛  Lint done 🪛"