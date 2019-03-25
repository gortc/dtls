lint:
	@golangci-lint run ./...
	@echo "ok"
install-fuzz:
	go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
	go get github.com/dvyukov/go-fuzz/go-fuzz
install:
	go get gortc.io/api
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
check-api:
	@cd api && bash ./check.sh
test:
	@./go.test.sh
