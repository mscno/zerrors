generate:
    just generate/generate
    go fmt ./...

test:
    @echo "Running tests"
    go test ./...
    cd ./grpczerrors && go test
