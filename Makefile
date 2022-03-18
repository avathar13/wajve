GOCMD = go
GOLINT = golangci-lint

run:
	@GO111MODULE=on $(GOCMD) run cmd/wajve/main.go

lint:
	@GO111MODULE=on $(GOLINT) run ./... -v